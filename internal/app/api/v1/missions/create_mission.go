package missions

import (
	"context"

	common "github.com/VeneLooool/missions-api/internal/app/api/v1"
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateMission(ctx context.Context, req *desc.CreateMission_Request) (*desc.CreateMission_Response, error) {
	mission, err := i.missionUC.Create(ctx, transformCreateMissionReqToModel(req), common.TransformMissionPlanToModel(req.GetPlan()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.CreateMission_Response{
		Mission: transformMissionToProto(mission),
	}, nil
}
