package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	dto "waysbuck/dto/result"
	jwtToken "waysbuck/pkg/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
		}

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
