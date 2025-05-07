package missions

import "github.com/VeneLooool/missions-api/internal/model"

func ValidateNewMissionStatus(status model.MissionStatus) bool {
	return status == model.MissionStatusPending || status == model.MissionStatusScheduled
}
