package missions

import (
	common "github.com/VeneLooool/missions-api/internal/app/api/v1"
	"github.com/VeneLooool/missions-api/internal/model"
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
	proto_model "github.com/VeneLooool/missions-api/internal/pb/api/v1/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func transformMissionStatusesToModel(protoStatuses []proto_model.MissionStatus) []model.MissionStatus {
	statuses := make([]model.MissionStatus, 0, len(protoStatuses))
	for _, protoStatus := range protoStatuses {
		statuses = append(statuses, common.MissionStatusToModel[protoStatus])
	}
	return statuses
}

func transformMissionsToProto(missions []model.Mission) []*proto_model.Mission {
	protoMissions := make([]*proto_model.Mission, 0, len(missions))
	for _, mission := range missions {
		protoMissions = append(protoMissions, transformMissionToProto(mission))
	}
	return protoMissions
}

func transformMissionToProto(mission model.Mission) *proto_model.Mission {
	return &proto_model.Mission{
		Id:   mission.ID,
		Name: mission.Name,

		Type:   common.MissionTypeToProto[mission.Type],
		Status: common.MissionStatusToProto[mission.Status],

		CreatedBy: mission.CreatedBy,
		UpdatedBy: mission.UpdatedBy,

		CreatedAt: timestamppb.New(mission.CreatedAt),
		UpdatedAt: timestamppb.New(mission.UpdatedAt),
		StartedAt: timestamppb.New(mission.StartedAt),

		FieldId: mission.FieldID,
		DroneId: mission.DroneID,
	}
}

func transformCreateMissionReqToModel(req *desc.CreateMission_Request) model.Mission {
	if req == nil {
		return model.Mission{}
	}

	return model.Mission{
		Name:   req.GetName(),
		Type:   common.MissionTypeToModel[req.GetType()],
		Status: common.MissionStatusToModel[req.GetStatus()],

		StartedAt: req.GetStartedAt().AsTime(),

		CreatedBy: req.GetCreatedBy(),
		UpdatedBy: req.GetCreatedBy(),

		FieldID: req.GetFieldId(),
		DroneID: req.GetDroneId(),
	}
}

func transformUpdateMissionReqToModel(req *desc.UpdateMission_Request) model.Mission {
	if req == nil {
		return model.Mission{}
	}

	return model.Mission{
		ID:        req.GetId(),
		Name:      req.GetName(),
		Status:    common.MissionStatusToModel[req.GetStatus()],
		StartedAt: req.GetStartedAt().AsTime(),
		UpdatedBy: req.UpdatedBy,
	}
}
