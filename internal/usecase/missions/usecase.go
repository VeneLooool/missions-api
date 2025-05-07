package missions

import (
	"context"
	"errors"

	"github.com/VeneLooool/missions-api/internal/model"
	"github.com/VeneLooool/missions-api/internal/pkg/error_hub"
	"github.com/jackc/pgx/v4"
)

type UseCase struct {
	repo Repo

	planner PlannerUC
}

func New(repo Repo, planner PlannerUC) *UseCase {
	return &UseCase{
		repo:    repo,
		planner: planner,
	}
}

func (u *UseCase) Create(ctx context.Context, mission model.Mission, plan model.Plan) (model.Mission, error) {
	if valid := ValidateNewMissionStatus(model.MissionStatusPending); !valid {
		return model.Mission{}, error_hub.ErrInvalidMissionStatus
	}

	mission.SetTimes()
	mission, err := u.repo.Create(ctx, mission)
	if err != nil {
		return model.Mission{}, err
	}

	_, err = u.planner.CreateMissionPlan(ctx, mission.ID, plan)
	if err != nil {
		return model.Mission{}, err
	}
	return mission, nil
}

func (u *UseCase) Update(ctx context.Context, mission model.Mission, plan model.Plan) (model.Mission, error) {
	mission.SetTimes()
	mission, err := u.repo.Update(ctx, mission)
	if err != nil {
		return model.Mission{}, err
	}

	_, err = u.planner.UpdateMissionPlan(ctx, mission.ID, plan)
	if err != nil {
		return model.Mission{}, err
	}
	return mission, nil
}

func (u *UseCase) UpdateStatus(ctx context.Context, id uint64, status model.MissionStatus) (model.Mission, error) {
	mission, err := u.Get(ctx, id)
	if err != nil {
		return model.Mission{}, err
	}
	mission.SetTimes()
	mission.SetStatus(status)

	mission, err = u.repo.Update(ctx, mission)
	if err != nil {
		return model.Mission{}, err
	}
	return mission, nil
}

func (u *UseCase) Get(ctx context.Context, id uint64) (model.Mission, error) {
	mission, err := u.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Mission{}, error_hub.ErrMissionNotFound
		}
		return model.Mission{}, err
	}
	return mission, nil
}

func (u *UseCase) GetByAuthor(ctx context.Context, authorLogin string) ([]model.Mission, error) {
	return u.repo.GetByAuthor(ctx, authorLogin)
}

func (u *UseCase) GetMissionsInStatuses(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error) {
	return u.repo.GetMissionsInStatuses(ctx, statuses)
}

func (u *UseCase) Delete(ctx context.Context, id uint64) error {
	return u.repo.Delete(ctx, id)
}
