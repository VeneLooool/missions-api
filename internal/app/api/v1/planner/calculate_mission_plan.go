package planner

import (
	"context"

	common "github.com/VeneLooool/missions-api/internal/app/api/v1"
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/planner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CalculateMissionPlan(ctx context.Context, req *desc.CalculateMissionPlan_Request) (*desc.CalculateMissionPlan_Response, error) {
	missionType := common.MissionTypeToModel[req.GetType()]
	borderCoordinates := common.TransformCoordinatesToModel(req.GetMissionBorders())

	plan, err := i.plannerUC.CalculateMissionPlan(ctx, borderCoordinates, missionType)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CalculateMissionPlan_Response{
		Plan: common.TransformMissionPlanToProto(plan),
	}, nil
}
