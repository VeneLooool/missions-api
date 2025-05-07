package missions

import (
	"context"
	
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeleteMission(ctx context.Context, req *desc.DeleteMission_Request) (*empty.Empty, error) {
	if err := i.missionUC.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}
