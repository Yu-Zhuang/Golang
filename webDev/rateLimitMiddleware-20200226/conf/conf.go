package conf

import "time"

const (
	// ServicePort : port of this service
	ServicePort = "80"
	// Database : database for rate limit middleware
	Database         = "redis"
	DatabaseName     = "0"
	DatabaseUser     = ""
	DatabasePassword = ""
	DatabaseAddr     = "redis://localhost:6379/"
	// RateLimit conf
	RateLimitNum      = 3
	RateLimitDuration = time.Minute
)
