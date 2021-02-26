package models

import "time"

// RateLimit : struct to record a ip rate limit
type RateLimit struct {
	IP string
	RateLimitValue
}

// RateLimitValue ...
type RateLimitValue struct {
	RemainNum  int
	ExpireTime time.Time
}
