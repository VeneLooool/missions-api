package planner

import "github.com/VeneLooool/missions-api/internal/pkg/ql"

const Table = "mission_plan"

var (
	ID          = ql.NewField(Table, "id")
	Coordinates = ql.NewField(Table, "coordinates")
	MissionID   = ql.NewField(Table, "mission_id")
)
