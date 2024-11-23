package rueidislimiter

import "time"

type RateLimit struct {
	limit  int
	window time.Duration
}
type RateLimitOption func(*RateLimit)

func WithLimit(limit int) RateLimitOption {
	return func(rl *RateLimit) {
		rl.limit = limit
	}
}

func WithWindow(window time.Duration) RateLimitOption {
	return func(rl *RateLimit) {
		rl.window = window
	}
}
