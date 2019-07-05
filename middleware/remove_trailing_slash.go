package middleware

import (
	"net/http"
	"strings"
)

func RemoveTrailingSlash(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	}
}
