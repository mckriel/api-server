package redis

import "time"

type CacheEntry struct {
	Key        string    `json:"key" redis:"key"`
	Value      string    `json:"value" redis:"value"`
	TTL        int64     `json:"ttl" redis:"ttl"`
	Created_At time.Time `json:"created_at" redis:"created_at"`
}