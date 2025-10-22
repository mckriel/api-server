package redis

import (
	"api-servers/internal/models/redis"
	"context"
	"time"
)

type SessionRepository interface {
	Create(ctx context.Context, session redis.Session) error
	GetByID(ctx context.Context, id string) (redis.Session, error)
	GetByToken(ctx context.Context, token string) (redis.Session, error)
	GetByUserID(ctx context.Context, userID string) ([]redis.Session, error)
	GetActiveByUserID(ctx context.Context, userID string) ([]redis.Session, error)
	Update(ctx context.Context, id string, session redis.Session) error
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteExpired(ctx context.Context) error
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
