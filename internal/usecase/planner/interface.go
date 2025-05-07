package planner

import (
	"context"

	"github.com/VeneLooool/missions-api/internal/model"
)

type FieldsClient interface {
	GetFieldByID(ctx context.Context, fieldID uint64) (model.Field, error)
}

type Repo interface {
	Create(ctx context.Context, plan model.Plan) (model.Plan, error)
	Update(ctx context.Context, plan model.Plan) (model.Plan, error)
	GetByMissionID(ctx context.Context, missionID uint64) (model.Plan, error)
}
