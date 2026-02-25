package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

const testSecret = "test-jwt-secret-key-at-least-32-bytes-long"

func TestGenerateJWT(t *testing.T) {
	tsc := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				token, err := GenerateJWT([]byte(testSecret), "user-123")
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsed, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
					return []byte(testSecret), nil
				})
				require.NoError(t, err)
				require.True(t, parsed.Valid)
				claims := parsed.Claims.(jwt.MapClaims)
				require.Equal(t, "user-123", claims["userId"])
			},
		},
		{
			name: "empty user ID",
			test: func(t *testing.T) {
				token, err := GenerateJWT([]byte(testSecret), "")
				require.NoError(t, err)
				require.NotEmpty(t, token)

				parsed, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
					return []byte(testSecret), nil
				})
				require.NoError(t, err)
				require.True(t, parsed.Valid)

				claims := parsed.Claims.(jwt.MapClaims)
				require.Equal(t, "", claims["userId"])
			},
		},
		{
			name: "different user IDs produce different tokens",
			test: func(t *testing.T) {
				t1, err := GenerateJWT([]byte(testSecret), "user-a")
				require.NoError(t, err)
				t2, err := GenerateJWT([]byte(testSecret), "user-b")
				require.NoError(t, err)
				require.NotEqual(t, t1, t2)
			},
		},
		{
			name: "token includes expiredAt claim",
			test: func(t *testing.T) {
				token, err := GenerateJWT([]byte(testSecret), "user-123")
				require.NoError(t, err)
				parsed, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
					return []byte(testSecret), nil
				})
				require.NoError(t, err)
				claims := parsed.Claims.(jwt.MapClaims)
				require.Contains(t, claims, "expiredAt")
				expAt, ok := claims["expiredAt"].(float64)
				require.True(t, ok)
				require.Greater(t, expAt, float64(time.Now().Unix()))
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, tc.test)
	}
}

func TestValidateJWT(t *testing.T) {
	// Use same secret as config so validateJWT (which uses config.Load().JWT.Key) accepts tokens we generate.
	t.Setenv("JWT_KEY", testSecret)

	tsc := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				token, err := GenerateJWT([]byte(testSecret), "user-123")
				require.NoError(t, err)
				require.NotEmpty(t, token)

				valid, err := validateJWT(token)
				require.NoError(t, err)
				require.True(t, valid.Valid)
				claims := valid.Claims.(jwt.MapClaims)
				require.Equal(t, "user-123", claims["userId"])
			},
		},
		{
			name: "empty token returns error",
			test: func(t *testing.T) {
				_, err := validateJWT("")
				require.Error(t, err)
				require.Contains(t, err.Error(), "token is required")
			},
		},
		{
			name: "malformed token returns error",
			test: func(t *testing.T) {
				_, err := validateJWT("not.a.valid.jwt")
				require.Error(t, err)
			},
		},
		{
			name: "tampered token returns error",
			test: func(t *testing.T) {
				token, err := GenerateJWT([]byte(testSecret), "user-123")
				require.NoError(t, err)
				// Tamper: change one character in the payload (middle part)
				parts := splitToken(token)
				require.Len(t, parts, 3)
				tampered := parts[0] + "." + parts[1] + "x." + parts[2]
				_, err = validateJWT(tampered)
				require.Error(t, err)
			},
		},
		{
			name: "wrong secret invalidates signature",
			test: func(t *testing.T) {
				// Generate with one secret
				token, err := GenerateJWT([]byte("other-secret-key"), "user-123")
				require.NoError(t, err)
				// Validate with JWT_KEY=testSecret (set above)
				_, err = validateJWT(token)
				require.Error(t, err)
				require.Contains(t, err.Error(), "signature")
			},
		},
		{
			name: "unexpected signing method returns error",
			test: func(t *testing.T) {
				// Token signed with RS256 so key func rejects non-HMAC
				key, err := rsa.GenerateKey(rand.Reader, 2048)
				require.NoError(t, err)
				token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"userId": "user-123"})
				tokenStr, err := token.SignedString(key)
				require.NoError(t, err)
				_, err = validateJWT(tokenStr)
				require.Error(t, err)
				require.Contains(t, err.Error(), "unexpected signing method")
			},
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, tc.test)
	}
}

func TestGetTokenFromRequest(t *testing.T) {
	tsc := []struct {
		name        string
		authHeader  string
		wantToken   string
	}{
		{
			name:       "empty header returns empty",
			authHeader: "",
			wantToken:  "",
		},
		{
			name:       "Bearer prefix returns token",
			authHeader: "Bearer my-token-here",
			wantToken:  "my-token-here",
		},
		{
			name:       "Bearer prefix case insensitive",
			authHeader: "bearer my-token",
			wantToken:  "my-token",
		},
		{
			name:       "no Bearer prefix returns empty",
			authHeader: "Basic dXNlcjpwYXNz",
			wantToken:  "",
		},
		{
			name:       "Bearer only no space returns empty",
			authHeader: "Bearer",
			wantToken:  "",
		},
		{
			name:       "Authorization with leading and trailing spaces",
			authHeader: "  Bearer  trimmed-token  ",
			wantToken:  "trimmed-token",
		},
	}
	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.authHeader != "" {
				req.Header.Set("Authorization", tc.authHeader)
			}
			got := getTokenFromRequest(req)
			require.Equal(t, tc.wantToken, got)
		})
	}
}

// splitToken splits a JWT into header.payload.signature (no verification).
func splitToken(s string) []string {
	return strings.Split(s, ".")
}

func TestMiddleware(t *testing.T) {
	t.Setenv("JWT_KEY", testSecret)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(UserCtxKey)
		w.WriteHeader(http.StatusOK)
		if userID != nil {
			w.Write([]byte(userID.(string)))
		}
	})
	handler := Middleware(next)

	tsc := []struct {
		name           string
		authHeader     string
		wantStatus     int
		wantBodySubstr string
	}{
		{
			name:           "no Authorization header returns 403",
			authHeader:     "",
			wantStatus:     http.StatusForbidden,
			wantBodySubstr: "permission denied",
		},
		{
			name:           "invalid token returns 403",
			authHeader:     "Bearer invalid.jwt.token",
			wantStatus:     http.StatusForbidden,
			wantBodySubstr: "permission denied",
		},
		{
			name:           "valid Bearer token passes and sets user ID in context",
			authHeader:     "", // set in test to valid token
			wantStatus:     http.StatusOK,
			wantBodySubstr: "user-456",
		},
	}

	for _, tc := range tsc {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.authHeader != "" {
				req.Header.Set("Authorization", tc.authHeader)
			}
			if tc.wantBodySubstr == "user-456" {
				token, err := GenerateJWT([]byte(testSecret), "user-456")
				require.NoError(t, err)
				req.Header.Set("Authorization", "Bearer "+token)
			}

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			require.Equal(t, tc.wantStatus, rec.Code)
			require.Contains(t, rec.Body.String(), tc.wantBodySubstr)
		})
	}

	t.Run("token without userId claim returns 403", func(t *testing.T) {
		// Token signed with testSecret but has no "userId" claim
		tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "user-789"})
		tokenStr, err := tok.SignedString([]byte(testSecret))
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenStr)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		require.Equal(t, http.StatusForbidden, rec.Code)
		require.Contains(t, rec.Body.String(), "permission denied")
	})

	t.Run("expired token returns 403", func(t *testing.T) {
		// Token with exp in the past so token.Valid is false after parse
		tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": "user-exp",
			"exp":    time.Now().Add(-1 * time.Hour).Unix(),
		})
		tokenStr, err := tok.SignedString([]byte(testSecret))
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenStr)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		require.Equal(t, http.StatusForbidden, rec.Code)
		require.Contains(t, rec.Body.String(), "permission denied")
	})
}
