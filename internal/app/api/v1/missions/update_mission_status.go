package missions

import (
	"context"
	common "github.com/VeneLooool/missions-api/internal/app/api/v1"
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateMissionStatus(ctx context.Context, req *desc.UpdateMissionStatus_Request) (*desc.UpdateMissionStatus_Response, error) {
	mission, err := i.missionUC.UpdateStatus(ctx, req.GetId(), common.MissionStatusToModel[req.GetStatus()])
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.UpdateMissionStatus_Response{
		Mission: transformMissionToProto(mission),
	}, nil
}
