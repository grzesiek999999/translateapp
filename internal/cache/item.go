package cache

import "time"

type Item struct {
	data  string
	ttl   time.Time
	timer time.Duration
}
