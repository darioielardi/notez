package core

import (
	"net/http"
	
	"github.com/gorilla/mux"
	
	"notez/utils/enums"
)

type ServerFunc func(s *Server) http.HandlerFunc

// Route is the route struct that holds info on the route to bind to the router
type Route struct {
	Name    string
	Method  string
	Path    string
	Auth    bool
	Roles   []enums.Role
	Handler ServerFunc
}

// Routes is a Route slice
type Routes []Route

func NewRouter() *mux.Router {
	return mux.NewRouter()
}
