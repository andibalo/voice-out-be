package constants

import "time"

const (
	APPLICATION_NAME        = "voiceout"
	JWT_EXPIRATION_DURATION = time.Hour * 3
)

const (
	BasePath        = "/api"
	V1BasePath      = "/api/v1"
	HealthAPIPath   = "/health"
	AuthAPIPath     = "/auth"
	RegisterAPIPath = "/register"
	LoginAPIPath    = "/login"
)

const (
	HeaderRequestID string = "X-Request-ID"
)
