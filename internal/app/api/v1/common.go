package v1

import (
	"github.com/VeneLooool/missions-api/internal/model"
	proto_model "github.com/VeneLooool/missions-api/internal/pb/api/v1/model"
)

var (
	MissionTypeToProto = map[model.MissionType]proto_model.MissionType{
		model.MissionTypePatrol:   proto_model.MissionType_MISSION_TYPE_PATROL,
		model.MissionTypeResearch: proto_model.MissionType_MISSION_TYPE_RESEARCH,
	}
	MissionStatusToProto = map[model.MissionStatus]proto_model.MissionStatus{
		model.MissionStatusCreated:   proto_model.MissionStatus_MISSION_STATUS_CREATED,
		model.MissionStatusScheduled: proto_model.MissionStatus_MISSION_STATUS_SCHEDULED,
		model.MissionStatusPending:   proto_model.MissionStatus_MISSION_STATUS_PENDING,
		model.MissionStatusRunning:   proto_model.MissionStatus_MISSION_STATUS_RUNNING,
		model.MissionStatusAnalyse:   proto_model.MissionStatus_MISSION_STATUS_ANALYSE,
		model.MissionStatusCanceled:  proto_model.MissionStatus_MISSION_STATUS_CANCELED,
		model.MissionStatusWarning:   proto_model.MissionStatus_MISSION_STATUS_WARNING,
		model.MissionStatusFailed:    proto_model.MissionStatus_MISSION_STATUS_FAILED,
		model.MissionStatusSuccess:   proto_model.MissionStatus_MISSION_STATUS_SUCCESS,
	}

	MissionTypeToModel = map[proto_model.MissionType]model.MissionType{
		proto_model.MissionType_MISSION_TYPE_PATROL:   model.MissionTypePatrol,
		proto_model.MissionType_MISSION_TYPE_RESEARCH: model.MissionTypeResearch,
	}
	MissionStatusToModel = map[proto_model.MissionStatus]model.MissionStatus{
		proto_model.MissionStatus_MISSION_STATUS_CREATED:   model.MissionStatusCreated,
		proto_model.MissionStatus_MISSION_STATUS_SCHEDULED: model.MissionStatusScheduled,
		proto_model.MissionStatus_MISSION_STATUS_PENDING:   model.MissionStatusPending,
		proto_model.MissionStatus_MISSION_STATUS_RUNNING:   model.MissionStatusRunning,
		proto_model.MissionStatus_MISSION_STATUS_ANALYSE:   model.MissionStatusAnalyse,
		proto_model.MissionStatus_MISSION_STATUS_CANCELED:  model.MissionStatusCanceled,
		proto_model.MissionStatus_MISSION_STATUS_WARNING:   model.MissionStatusWarning,
		proto_model.MissionStatus_MISSION_STATUS_FAILED:    model.MissionStatusFailed,
		proto_model.MissionStatus_MISSION_STATUS_SUCCESS:   model.MissionStatusSuccess,
	}
)

func TransformCoordinatesToModel(protoCoordinates []*proto_model.Coordinate) model.Coordinates {
	coordinates := make([]model.Coordinate, 0, len(protoCoordinates))
	for _, coordinate := range protoCoordinates {
		coordinates = append(coordinates, model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return coordinates
}

func TransformCoordinatesToProto(coordinates model.Coordinates) []*proto_model.Coordinate {
	protoCoordinates := make([]*proto_model.Coordinate, 0, len(coordinates))
	for _, coordinate := range coordinates {
		protoCoordinates = append(protoCoordinates, &proto_model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return protoCoordinates
}

func TransformMissionPlanToProto(plan model.Plan) *proto_model.Plan {
	return &proto_model.Plan{
		Coordinates: TransformCoordinatesToProto(plan.Coordinates),
	}
}

func TransformMissionPlanToModel(protoPlan *proto_model.Plan) model.Plan {
	return model.Plan{
		Coordinates: TransformCoordinatesToModel(protoPlan.GetCoordinates()),
	}
}
