package redis

import (
	"api-servers/internal/models/redis"
	"context"
	"encoding/json"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type sessionRepository struct {
	db *Database
}

func NewSessionRepository(db *Database) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) Create(ctx context.Context, session redis.Session) error {
	ttl := time.Until(session.Expires_At)
	if ttl <= 0 {
		return fmt.Errorf("session expired")
	}

	session_data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	key := "session:" + session.ID
	err = r.db.Connection.Set(ctx, key, session_data, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	token_key := "session:token:" + session.Token
	err = r.db.Connection.Set(ctx, token_key, session.ID, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to create token mapping: %w", err)
	}
	return nil
}

func (r *sessionRepository) GetByID(ctx context.Context, id string) (redis.Session, error) {
	var session redis.Session

	key := "session:" + id
	result := r.db.Connection.Get(ctx, key)

	if err := result.Err(); err != nil {
		if err == goredis.Nil {
			return session, fmt.Errorf("session with id %s not found", id)
		}
		return session, fmt.Errorf("failed to get session: %w", err)
	}
	session_data := result.Val()
	err := json.Unmarshal([]byte(session_data), &session)
	if err != nil {
		return session, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return session, nil

}

// Add the other 6 functions as stubs for now:
func (r *sessionRepository) GetByToken(ctx context.Context, token string) (redis.Session, error) {
	var session redis.Session

	token_key := "session:token:" + token
	result := r.db.Connection.Get(ctx, token_key)

	if err := result.Err(); err != nil {
		if err == goredis.Nil {
			return session, fmt.Errorf("session with token %s not found", token)
		}
		return session, fmt.Errorf("failed to get token mapping: %w", err)
	}

	// Get session ID from token mapping
	session_id := result.Val()
	return r.GetByID(ctx, session_id)
}

func (r *sessionRepository) GetByUserID(ctx context.Context, userID string) ([]redis.Session, error) {
	var sessions []redis.Session

	pattern := "session:*"
	keys := r.db.Connection.Keys(ctx, pattern)

	if err := keys.Err(); err != nil {
		return sessions, fmt.Errorf("failed to get session keys: %w", err)
	}

	for _, key := range keys.Val() {
		session_id := key[8:] // Remove "session:" prefix
		session, err := r.GetByID(ctx, session_id)
		if err != nil {
			continue // Skip sessions that can't be retrieved
		}
		if session.User_ID == userID {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

func (r *sessionRepository) GetActiveByUserID(ctx context.Context, userID string) ([]redis.Session, error) {
	sessions, err := r.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user sessions: %w", err)
	}

	var active_sessions []redis.Session
	now := time.Now()

	for _, session := range sessions {
		if session.Active && session.Expires_At.After(now) {
			active_sessions = append(active_sessions, session)
		}
	}

	return active_sessions, nil
}

func (r *sessionRepository) Update(ctx context.Context, id string, session redis.Session) error {
	// Check if session exists
	_, err := r.GetByID(ctx, id)
	if err != nil {
		return err // Session not found
	}

	// Set the ID to ensure consistency
	session.ID = id

	// Delete old session and token mapping
	err = r.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete old session: %w", err)
	}

	// Create new session
	return r.Create(ctx, session)
}

func (r *sessionRepository) Delete(ctx context.Context, id string) error {
	// Get session to find token for cleanup
	session, err := r.GetByID(ctx, id)
	if err != nil {
		return err // Session not found or other error
	}

	// Delete main session key
	key := "session:" + id
	err = r.db.Connection.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	// Delete token mapping
	token_key := "session:token:" + session.Token
	err = r.db.Connection.Del(ctx, token_key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete token mapping: %w", err)
	}

	return nil
}

func (r *sessionRepository) DeleteByUserID(ctx context.Context, userID string) error {
	sessions, err := r.GetByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user sessions: %w", err)
	}

	for _, session := range sessions {
		err := r.Delete(ctx, session.ID)
		if err != nil {
			// Log error but continue with other sessions
			continue
		}
	}

	return nil
}

func (r *sessionRepository) DeleteExpired(ctx context.Context) error {
	pattern := "session:*"
	keys := r.db.Connection.Keys(ctx, pattern)

	if err := keys.Err(); err != nil {
		return fmt.Errorf("failed to get session keys: %w", err)
	}

	now := time.Now()
	deleted_count := 0

	for _, key := range keys.Val() {
		session_id := key[8:] // Remove "session:" prefix
		session, err := r.GetByID(ctx, session_id)
		if err != nil {
			continue // Skip sessions that can't be retrieved
		}

		// Delete if expired
		if session.Expires_At.Before(now) {
			err := r.Delete(ctx, session_id)
			if err == nil {
				deleted_count++
			}
		}
	}

	return nil
}
