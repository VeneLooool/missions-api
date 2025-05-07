package missions

import (
	"context"
	
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetMissionsByAuthor(ctx context.Context, req *desc.GetMissionsByAuthor_Request) (*desc.GetMissionsByAuthor_Response, error) {
	missions, err := i.missionUC.GetByAuthor(ctx, req.GetLogin())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.GetMissionsByAuthor_Response{
		Missions: transformMissionsToProto(missions),
	}, nil
}
