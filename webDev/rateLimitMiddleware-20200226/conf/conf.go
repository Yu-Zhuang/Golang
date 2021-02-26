package conf

import "time"

const (
	ServicePort = "80"
	// DataBase
	Database         = "redis"
	DatabaseName     = "0"
	DatabaseUser     = ""
	DatabasePassword = ""
	DatabaseAddr     = "redis://localhost:6379/"
	// RateLimit
	RateLimitNum      = 4
	RateLimitDuration = time.Minute
)
