package users

import (
	"notez/core"
	"notez/utils/enums"
)

var basePath = "/users"

var Routes = core.Routes{
	core.Route{
		Name:    "Find One User",
		Method:  "GET",
		Path:    basePath + "/{id}",
		Auth:    true,
		Roles:   enums.Roles{"admin"},
		Handler: FindOne,
	},
}
