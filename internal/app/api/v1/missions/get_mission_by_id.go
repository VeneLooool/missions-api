package missions

import (
	"context"
	"errors"
	
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	"github.com/VeneLooool/missions-api/internal/pkg/error_hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetMissionByID(ctx context.Context, req *desc.GetMissionByID_Request) (*desc.GetMissionByID_Response, error) {
	mission, err := i.missionUC.Get(ctx, req.GetId())
	if err != nil {
		code := codes.Internal
		if errors.Is(err, error_hub.ErrMissionNotFound) {
			code = codes.NotFound
		}
		return nil, status.Error(code, err.Error())
	}
	return &desc.GetMissionByID_Response{
		Mission: transformMissionToProto(mission),
	}, nil
}
