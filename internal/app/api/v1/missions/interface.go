package missions

import (
	"context"

	"github.com/VeneLooool/missions-api/internal/model"
)

type MissionUC interface {
	Create(ctx context.Context, mission model.Mission, plan model.Plan) (model.Mission, error)
	Update(ctx context.Context, mission model.Mission, plan model.Plan) (model.Mission, error)
	UpdateStatus(ctx context.Context, id uint64, status model.MissionStatus) (model.Mission, error)

	Get(ctx context.Context, id uint64) (model.Mission, error)
	GetByAuthor(ctx context.Context, authorLogin string) ([]model.Mission, error)
	GetMissionsInStatuses(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error)

	Delete(ctx context.Context, id uint64) error
}
