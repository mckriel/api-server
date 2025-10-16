package redis

import (
	"api-servers/internal/models/redis"
	"time"
)

type SessionRepository interface {
	Create(session redis.Session) error
	GetByID(id string) (redis.Session, error)
	GetByToken(token string) (redis.Session, error)
	GetByUserID(userID string) ([]redis.Session, error)
	GetActiveByUserID(userID string) ([]redis.Session, error)
	Update(id string, session redis.Session) error
	Delete(id string) error
	DeleteByUserID(userID string) error
	DeleteExpired() error
}

type CacheRepository interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	GetMultiple(keys []string) (map[string]string, error)
	Delete(key string) error
	DeleteMultiple(keys []string) error
	Exists(key string) (bool, error)
	SetTTL(key string, ttl time.Duration) error
	GetTTL(key string) (time.Duration, error)
	Clear() error
}
