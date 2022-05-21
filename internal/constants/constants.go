package constants

import "time"

const (
	APPLICATION_NAME        = "voiceout"
	JWT_EXPIRATION_DURATION = time.Hour * 3
)

const (
	BasePath                  = "/api"
	V1BasePath                = "/api/v1"
	HealthAPIPath             = "/health"
	PostAPIPath               = "/post"
	AuthAPIPath               = "/auth"
	UserAPIPath               = "/user"
	RegisterAPIPath           = "/register"
	LoginAPIPath              = "/login"
	FetchPostsByUserIDAPIPath = "/user/:id"
	FetchCurrentUserAPIPath   = "/me"
)

const (
	HeaderRequestID string = "X-Request-ID"
)
