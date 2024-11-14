package middleware

import (
	"context"
	"go/projcet-Adv/configs"
	"go/projcet-Adv/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnathued(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authedHeader := req.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer") {
			writeUnathued(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnathued(w)
			return
		}
		ctx := context.WithValue(req.Context(), ContextEmailKey, data.Email)
		r := req.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
