package error_hub

import "errors"

var (
	ErrMissionNotFound      = errors.New("mission not found")
	ErrInvalidMissionStatus = errors.New("invalid mission status")
	ErrInvalidMissionType   = errors.New("invalid mission type")
)
