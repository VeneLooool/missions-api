package missions

import (
	"context"

	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetMissionsInStatuses(ctx context.Context, req *desc.GetMissionsInStatuses_Request) (*desc.GetMissionsInStatuses_Response, error) {
	missions, err := i.missionUC.GetMissionsInStatuses(ctx, transformMissionStatusesToModel(req.GetStatuses()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.GetMissionsInStatuses_Response{
		Missions: transformMissionsToProto(missions),
	}, nil
}
