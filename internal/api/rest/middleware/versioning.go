package middleware

import (
	"context"
	"net/http"
)

const (
	APIVersionHeader = "API-Version"
	VersionKey       = "api_version"

	Version20241001 = "2024-10-01"
	Version20241218 = "2024-12-18"
	DefaultVersion  = Version20241218
)

func VersioningMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		version := r.Header.Get(APIVersionHeader)
		if version == "" {
			version = DefaultVersion
		}

		ctx := context.WithValue(r.Context(), VersionKey, version)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetVersionFromContext(ctx context.Context) string {
	if version, ok := ctx.Value(VersionKey).(string); ok {
		return version
	}
	return DefaultVersion
}
