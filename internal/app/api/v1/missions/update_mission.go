package missions

import (
	"context"

	common "github.com/VeneLooool/missions-api/internal/app/api/v1"
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateMission(ctx context.Context, req *desc.UpdateMission_Request) (*desc.UpdateMission_Response, error) {
	mission, err := i.missionUC.Update(ctx, transformUpdateMissionReqToModel(req), common.TransformMissionPlanToModel(req.GetPlan()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.UpdateMission_Response{
		Mission: transformMissionToProto(mission),
	}, nil
}
