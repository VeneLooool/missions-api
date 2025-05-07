package planner

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

func (r *Repo) Create(ctx context.Context, plan model.Plan) (model.Plan, error) {
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder().InsertInto(Table).Cols(
		Coordinates.Short(),
		MissionID.Short(),
	).Values(
		plan.Coordinates,
		plan.MissionID,
	)
	query, args := common.ReturningAll(ib).Build()

	var result model.Plan
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Plan{}, err
	}
	return result, nil
}

func (r *Repo) Update(ctx context.Context, plan model.Plan) (model.Plan, error) {
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder().Update(Table)
	ub = ub.Set(ql.Fields{
		Coordinates,
	}.ToAssignments(ub,
		plan.Coordinates,
	)...)
	ub = ub.Where(ub.Equal(MissionID.Short(), plan.MissionID))
	query, args := common.ReturningAll(ub).Build()

	var result model.Plan
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Plan{}, err
	}
	return result, nil
}

func (r *Repo) Get(ctx context.Context, id uint64) (model.Plan, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(ID.Short(), id))
	query, args := sb.Build()

	var result model.Plan
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Plan{}, err
	}
	return result, nil
}

func (r *Repo) GetByMissionID(ctx context.Context, missionID uint64) (model.Plan, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(MissionID.Short(), missionID))
	query, args := sb.Build()

	var result model.Plan
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Plan{}, err
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
