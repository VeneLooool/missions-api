package planner

import (
	"context"
	"errors"
	
	common "github.com/VeneLooool/missions-api/internal/app/api/v1"
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/planner"
	"github.com/VeneLooool/missions-api/internal/pkg/error_hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetMissionPlanByMissionID(ctx context.Context, req *desc.GetMissionPlanByMissionID_Request) (*desc.GetMissionPlanByMissionID_Response, error) {
	missionPlan, err := i.plannerUC.GetMissionPlanByMissionID(ctx, req.GetMissionId())
	if err != nil {
		code := codes.Internal
		if errors.Is(err, error_hub.ErrMissionPlanNotFound) {
			code = codes.NotFound
		}
		return nil, status.Error(code, err.Error())
	}
	return &desc.GetMissionPlanByMissionID_Response{
		Plan: common.TransformMissionPlanToProto(missionPlan),
	}, nil
}
