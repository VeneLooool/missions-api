package planner

import (
	"context"

	"github.com/VeneLooool/missions-api/internal/model"
)

type PlannerUC interface {
	CalculateMissionPlan(ctx context.Context, borderCoordinates model.Coordinates, missionType model.MissionType) (model.Plan, error)
	GetMissionPlanByMissionID(ctx context.Context, missionID uint64) (model.Plan, error)
}
