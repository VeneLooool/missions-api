package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/VeneLooool/missions-api/internal/app/api/v1/missions"
	"github.com/VeneLooool/missions-api/internal/app/api/v1/planner"
	"github.com/VeneLooool/missions-api/internal/config"
	mission_pb "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	planner_pb "github.com/VeneLooool/missions-api/internal/pb/api/v1/planner"
	"github.com/VeneLooool/missions-api/internal/pkg/db"
	mission_repo "github.com/VeneLooool/missions-api/internal/repository/missions"
	planner_repo "github.com/VeneLooool/missions-api/internal/repository/planner"
	mission_uc "github.com/VeneLooool/missions-api/internal/usecase/missions"
	planner_uc "github.com/VeneLooool/missions-api/internal/usecase/planner"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatalf("failed to create new config: %s", err.Error())
	}

	go func() {
		if err := runGRPC(ctx, cfg); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runHTTPGateway(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(ctx context.Context, cfg *config.Config) error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	dbAdapter, err := db.New(ctx)
	if err != nil {
		return err
	}
	defer dbAdapter.Close(ctx)

	missionsServer, plannerService, err := newServices(ctx, dbAdapter)
	if err != nil {
		return err
	}
	mission_pb.RegisterMissionsServer(grpcServer, missionsServer)
	planner_pb.RegisterPlannerServer(grpcServer, plannerService)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	log.Printf("gRPC server listening on :%s\n", cfg.GrpcPort)
	if err = grpcServer.Serve(grpcListener); err != nil {
		return err
	}
	return nil
}

func runHTTPGateway(ctx context.Context, cfg *config.Config) error {
	mux := runtime.NewServeMux()
	err := mission_pb.RegisterMissionsHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", cfg.GrpcPort), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err.Error())
	}
	err = planner_pb.RegisterPlannerHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", cfg.GrpcPort), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err.Error())
	}

	// Serve Swagger JSON and Swagger UI
	fs := http.FileServer(http.Dir("./swagger-ui")) // директория со статикой UI
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	// Serve Swagger JSON файл
	http.HandleFunc("/swagger/missions.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/pb/api/v1/missions/missions.swagger.json")
	})
	// Serve Swagger JSON файл
	http.HandleFunc("/swagger/planner.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/pb/api/v1/planner/planner.swagger.json")
	})

	withCORS := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Для preflight-запросов
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	// gRPC → REST mux
	http.Handle("/", withCORS(mux))

	log.Printf("HTTP gateway listening on :%s\n", cfg.HttpPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), nil); err != nil {
		return err
	}

	return nil
}

func newServices(ctx context.Context, dbAdapter db.DataBase) (*missions.Implementation, *planner.Implementation, error) {
	plannerRepo := planner_repo.New(dbAdapter)
	plannerUC := planner_uc.New(plannerRepo)

	missionRepo := mission_repo.New(dbAdapter)
	missionUC := mission_uc.New(missionRepo, plannerUC)

	return missions.NewService(missionUC), planner.NewService(plannerUC), nil
}
