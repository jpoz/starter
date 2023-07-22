package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jpoz/starter/pkg/config"
)

func NewServer(cfg config.Server) *Server {
	return &Server{
		Config: cfg,
	}
}

type Server struct {
	Config config.Server
}

func (s *Server) Router() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// JSON
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Cookie", "X-Zendesk-Host", "X-Zendesk-Subdomain"},
		ExposedHeaders:   []string{"Link", "SetCookie"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		Debug:            (log.GetLevel() == log.DebugLevel),
	}))

	// Health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("{\"status\": \"ok\"}")) })

	// Routes
	r.Route("/graphql", func(r chi.Router) {
		r.Get("/", playground.Handler("GraphQL playground", "/graphql/query"))
		r.Mount("/query", s.GraphQLServer())
	})

	return r
}

func (s *Server) ListenAndServe() error {
	log.Infof("Server running on %s", s.Config.Addr)
	router := s.Router()
	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Debugf("[%s]:\t%s has %d middlewares", method, route, len(middlewares))
		return nil
	})
	return http.ListenAndServe(s.Config.Addr, router)
}
