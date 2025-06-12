package middleware

import (
	"context"
	"net/http"
	"strings"

	jwttoken "github.com/khandev-bac/lemon/utils/jwtToken"
)

type contextKey string

const userIDKey contextKey = "userID"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwttoken.VerifyJWTAccessToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}
		userId, ok := claims["id"].(string)
		if !ok {
			http.Error(w, "Unauthorized: Invalid claims", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(userIDKey)
	if id, ok := val.(string); ok {
		return id, nil
	}
	return "", nil
}
