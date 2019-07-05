package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	
	"notez/config"
	"notez/middleware"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
	Config *config.Config
	V      *Validator
}

func NewServer(router *mux.Router, DB *gorm.DB, config *config.Config) *Server {
	return &Server{
		Router: router,
		DB:     DB,
		Config: config,
		V:      NewValidator(),
	}
}

func (s *Server) Init(routesGroups []Routes) {
	
	var routes Routes
	
	for _, r := range routesGroups {
		routes = append(routes, r...)
	}
	
	s.Wire(routes)
	s.Serve()
}

func (s *Server) Serve() {
	port := fmt.Sprintf(":%s", s.Config.Server.Port)
	
	if s.Config.Debug {
		fmt.Printf("Server listening on port: %s\n", port)
	}
	
	var router http.Handler = s.Router
	
	// Custom Global Middlewares
	
	router = middleware.RemoveTrailingSlash(router)
	
	// Gorilla Global Middleware
	
	router = handlers.LoggingHandler(os.Stdout, router)
	handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)(s.Router)
	
	log.Fatal(http.ListenAndServe(port, router))
}

func (s *Server) Wire(routes Routes) *Server {
	for _, route := range routes {
		
		handler := route.Handler(s)
		
		// Custom Local Middlewares
		
		handler = middleware.CheckAuth(handler, route.Auth, route.Roles)
		
		s.Router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	
	return s
}
