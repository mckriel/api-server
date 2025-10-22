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
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	GetMultiple(ctx context.Context, keys []string) (map[string]string, error)
	Delete(ctx context.Context, key string) error
	DeleteMultiple(ctx context.Context, keys []string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetTTL(ctx context.Context, key string, ttl time.Duration) error
	GetTTL(ctx context.Context, key string) (time.Duration, error)
	Clear(ctx context.Context) error
}
