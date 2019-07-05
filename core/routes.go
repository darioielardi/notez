package core

import (
	"net/http"
	
	"github.com/gorilla/mux"
	
	"notez/utils/enums"
)

type ServerFunc func(s *Server) http.HandlerFunc

type Route struct {
	Name    string
	Method  string
	Path    string
	Auth    bool
	Roles   []enums.Role
	Handler ServerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	return mux.NewRouter()
}
