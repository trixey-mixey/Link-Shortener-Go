package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authedHeader := req.Header.Get("Authorization")
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		fmt.Println(token)
		next.ServeHTTP(w, req)
	})
}
