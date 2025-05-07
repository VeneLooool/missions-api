package missions

import (
	"context"

	"github.com/VeneLooool/missions-api/internal/model"
)

type Repo interface {
	Create(ctx context.Context, mission model.Mission) (model.Mission, error)
	Update(ctx context.Context, mission model.Mission) (model.Mission, error)
	Get(ctx context.Context, id uint64) (model.Mission, error)
	GetByAuthor(ctx context.Context, authorLogin string) ([]model.Mission, error)
	GetMissionsInStatuses(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error)
	Delete(ctx context.Context, id uint64) error
}

type PlannerUC interface {
	CreateMissionPlan(ctx context.Context, missionID uint64, plan model.Plan) (model.Plan, error)
	UpdateMissionPlan(ctx context.Context, missionID uint64, plan model.Plan) (model.Plan, error)
	GetMissionPlanByMissionID(ctx context.Context, missionID uint64) (model.Plan, error)
}
