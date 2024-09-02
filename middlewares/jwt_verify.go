package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mframadann/gourl/utils"
)

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JWT_SIGNING_METHOD := jwt.SigningMethodHS256
		ACCESS_SECRET_KEY := os.Getenv("ACCESS_SECRET_KEY")
		authorizationHeader := r.Header.Get("Authorization")

		if r.URL.Path == "/api/v1/sign-in" || r.URL.Path == "/api/v1/register" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.Contains(authorizationHeader, "Bearer") {
			response := &utils.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)

			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(ACCESS_SECRET_KEY), nil
		})

		if err != nil {
			response := &utils.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := &utils.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid or token has expired",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

			return
		}

		ctx := context.WithValue(context.Background(), "user", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
