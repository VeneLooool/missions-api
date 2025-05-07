package planner

import (
	"context"
	"errors"
	"github.com/VeneLooool/missions-api/internal/model"
	"github.com/VeneLooool/missions-api/internal/pkg/error_hub"
	"github.com/jackc/pgx/v4"
)

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) CalculateMissionPlan(ctx context.Context, borderCoordinates model.Coordinates, missionType model.MissionType) (model.Plan, error) {
	switch missionType {
	case model.MissionTypePatrol:
		return model.Plan{Coordinates: borderCoordinates}, nil
	case model.MissionTypeResearch:
		return model.Plan{Coordinates: GenerateLawnMowerPath(borderCoordinates)}, nil
	}

	return model.Plan{}, error_hub.ErrInvalidMissionType
}

func (uc *UseCase) CreateMissionPlan(ctx context.Context, missionID uint64, plan model.Plan) (model.Plan, error) {
	plan.MissionID = missionID

	return uc.repo.Create(ctx, plan)
}

func (uc *UseCase) UpdateMissionPlan(ctx context.Context, missionID uint64, plan model.Plan) (model.Plan, error) {
	plan.MissionID = missionID

	return uc.repo.Update(ctx, plan)
}

func (uc *UseCase) GetMissionPlanByMissionID(ctx context.Context, missionID uint64) (model.Plan, error) {
	plan, err := uc.repo.GetByMissionID(ctx, missionID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Plan{}, error_hub.ErrMissionPlanNotFound
		}
		return model.Plan{}, err
	}
	return plan, nil
}
