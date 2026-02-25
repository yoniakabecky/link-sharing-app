package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
)

// signToken is overridable in tests to simulate signing failures.
// var signToken = func(token *jwt.Token, key interface{}) (string, error) {
// 	return token.SignedString(key)
// }

const UserCtxKey = "userId"

func GenerateJWT(secret []byte, userID string) (string, error) {
	expiration := time.Second * time.Duration(config.Load().JWT.Exp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, errors.New("token is required")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Load().JWT.Key), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func permissionDenied(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(map[string]string{"error": "permission denied"})
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth == "" {
		return ""
	}
	tokenAuth = strings.TrimSpace(tokenAuth)

	const bearerPrefix = "Bearer "
	if len(tokenAuth) > len(bearerPrefix) && strings.EqualFold(tokenAuth[:len(bearerPrefix)], bearerPrefix) {
		return strings.TrimSpace(tokenAuth[len(bearerPrefix):])
	}
	return ""
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v\n", err)
			permissionDenied(w)
			return
		}
		if !token.Valid {
			log.Printf("invalid token\n")
			permissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["userId"]
		if !ok {
			log.Printf("failed to get user ID\n")
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserCtxKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
