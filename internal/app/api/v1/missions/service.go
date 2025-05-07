package missions

import (
	desc "github.com/VeneLooool/missions-api/internal/pb/api/v1/missions"
)

// Implementation is a Service implementation
type Implementation struct {
	desc.UnimplementedMissionsServer

	missionUC MissionUC
}

// NewService return new instance of Implementation.
func NewService(missionUC MissionUC) *Implementation {
	return &Implementation{
		missionUC: missionUC,
	}
}
