package model

type Coordinates []Coordinate

type Coordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Plan struct {
	ID          uint64      `db:"id"`
	Coordinates Coordinates `db:"coordinates"`
	MissionID   uint64      `db:"mission_id"`
}
