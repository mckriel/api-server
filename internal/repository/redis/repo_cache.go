package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	db *Database
}

func NewCacheRepository(db *Database) CacheRepository {
	return &cacheRepository{
		db: db,
	}
}

func (r *cacheRepository) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	err := r.db.Connection.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache key %s: %w", key, err)
	}
	return nil
}

func (r *cacheRepository) Get(ctx context.Context, key string) (string, error) {
	result := r.db.Connection.Get(ctx, key)

	if err := result.Err(); err != nil {
		if err == goredis.Nil {
			return "", fmt.Errorf("cache key %s not found", key)
		}
		return "", fmt.Errorf("failed to get cache key %s: %w", key, err)
	}

	return result.Val(), nil
}

func (r *cacheRepository) GetMultiple(ctx context.Context, keys []string) (map[string]string, error) {
	if len(keys) == 0 {
		return make(map[string]string), nil
	}

	result := r.db.Connection.MGet(ctx, keys...)
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to get multiple cache keys: %w", err)
	}

	cache_map := make(map[string]string)
	values := result.Val()

	for i, value := range values {
		if value != nil {
			cache_map[keys[i]] = value.(string)
		}
	}

	return cache_map, nil
}

func (r *cacheRepository) Delete(ctx context.Context, key string) error {
	err := r.db.Connection.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete cache key %s: %w", key, err)
	}
	return nil
}

func (r *cacheRepository) DeleteMultiple(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	err := r.db.Connection.Del(ctx, keys...).Err()
	if err != nil {
		return fmt.Errorf("failed to delete multiple cache keys: %w", err)
	}
	return nil
}

func (r *cacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	result := r.db.Connection.Exists(ctx, key)
	if err := result.Err(); err != nil {
		return false, fmt.Errorf("failed to check if cache key %s exists: %w", key, err)
	}

	return result.Val() > 0, nil
}

func (r *cacheRepository) SetTTL(ctx context.Context, key string, ttl time.Duration) error {
	result := r.db.Connection.Expire(ctx, key, ttl)
	if err := result.Err(); err != nil {
		return fmt.Errorf("failed to set TTL for cache key %s: %w", key, err)
	}

	if !result.Val() {
		return fmt.Errorf("cache key %s does not exist, cannot set TTL", key)
	}

	return nil
}

func (r *cacheRepository) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	result := r.db.Connection.TTL(ctx, key)
	if err := result.Err(); err != nil {
		return 0, fmt.Errorf("failed to get TTL for cache key %s: %w", key, err)
	}

	ttl := result.Val()

	if ttl == -2*time.Second {
		return 0, fmt.Errorf("cache key %s does not exist", key)
	}
	if ttl == -1*time.Second {
		return -1, nil
	}

	return ttl, nil
}

func (r *cacheRepository) Clear(ctx context.Context) error {
	err := r.db.Connection.FlushDB(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to clear cache: %w", err)
	}
	return nil
}
