package missions

import "github.com/VeneLooool/missions-api/internal/pkg/ql"

const Table = "missions"

var (
	ID     = ql.NewField(Table, "id")
	Name   = ql.NewField(Table, "name")
	Type   = ql.NewField(Table, "type")
	Status = ql.NewField(Table, "status")

	CreatedBy = ql.NewField(Table, "created_by")
	UpdatedBy = ql.NewField(Table, "updated_by")

	CreatedAt = ql.NewField(Table, "created_at")
	StartedAt = ql.NewField(Table, "started_at")
	UpdatedAt = ql.NewField(Table, "updated_at")

	FieldID = ql.NewField(Table, "field_id")
	DroneID = ql.NewField(Table, "drone_id")
)
