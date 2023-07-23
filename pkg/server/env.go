package server

import (
	"net/http"

	"github.com/jpoz/starter/pkg/env"
)

func (s *Server) EnvMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := env.Attach(r.Context(), s.Query, s.Redis)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
