package missions

import (
	"context"

	"github.com/VeneLooool/missions-api/internal/model"
	"github.com/VeneLooool/missions-api/internal/pkg/db"
	"github.com/VeneLooool/missions-api/internal/pkg/ql"
	common "github.com/VeneLooool/missions-api/internal/repository"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/huandu/go-sqlbuilder"
)

type Repo struct {
	db db.DataBase
}

func New(db db.DataBase) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, mission model.Mission) (model.Mission, error) {
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder().InsertInto(Table).Cols(
		Name.Short(),
		Type.Short(),
		Status.Short(),
		CreatedBy.Short(),
		UpdatedBy.Short(),

		CreatedAt.Short(),
		StartedAt.Short(),
		UpdatedAt.Short(),

		FieldID.Short(),
		DroneID.Short(),
	).Values(
		mission.Name,
		mission.Type,
		mission.Status,
		mission.CreatedBy,
		mission.UpdatedBy,

		mission.CreatedAt,
		mission.UpdatedAt,
		mission.StartedAt,

		mission.FieldID,
		mission.DroneID,
	)
	query, args := common.ReturningAll(ib).Build()

	var result model.Mission
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Mission{}, err
	}
	return result, nil
}

func (r *Repo) Update(ctx context.Context, mission model.Mission) (model.Mission, error) {
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder().Update(Table)
	ub = ub.Set(ql.Fields{
		Name,
		Status,
		UpdatedBy,
		UpdatedAt,
		StartedAt,
	}.ToAssignments(ub,
		mission.Name,
		mission.Status,
		mission.UpdatedBy,
		mission.UpdatedAt,
		mission.StartedAt,
	)...)
	ub = ub.Where(ub.Equal(ID.Short(), mission.ID))
	query, args := common.ReturningAll(ub).Build()

	var result model.Mission
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Mission{}, err
	}
	return result, nil
}

func (r *Repo) Get(ctx context.Context, id uint64) (model.Mission, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(ID.Short(), id))
	query, args := sb.Build()

	var result model.Mission
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Mission{}, err
	}
	return result, nil
}

func (r *Repo) GetByAuthor(ctx context.Context, authorLogin string) ([]model.Mission, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(CreatedBy.Short(), authorLogin))
	query, args := sb.Build()

	var result []model.Mission
	if err := pgxscan.Select(ctx, r.db, &result, query, args...); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repo) GetMissionsInStatuses(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error) {
	if len(statuses) == 0 {
		return []model.Mission{}, nil
	}

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.In(Status.Short(), sqlbuilder.Flatten(statuses)...))
	query, args := sb.Build()

	var result []model.Mission
	if err := pgxscan.Select(ctx, r.db, &result, query, args...); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repo) Delete(ctx context.Context, id uint64) error {
	db := sqlbuilder.PostgreSQL.NewDeleteBuilder().DeleteFrom(Table)
	query, args := db.Where(db.Equal(ID.Short(), id)).Build()

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
