package middleware

import (
	"context"
	"net/http"

	"github.com/gauravst/go-api-template/internal/config"
	"github.com/gauravst/go-api-template/internal/utils/jwtToken"
	"github.com/gauravst/go-api-template/internal/utils/response"
)

type contextKey string

const userDataKey contextKey = "userData"

func Auth(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the request headers
			cookie, err := r.Cookie("accessToken")
			if err != nil {
				response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(err))
				return
			}
			token := cookie.Value

			// If the token is valid, call the next handler
			userData, err := jwtToken.VerifyJwtAndGetData(token, cfg.JwtPrivateKey)
			if err != nil {
				response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(err))
				return
			}

			ctx := context.WithValue(r.Context(), userDataKey, userData)
			r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
