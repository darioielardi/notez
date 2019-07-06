package users

import (
	"notez/core"
	roles "notez/utils/enums"
)

var basePath = "/users"

var Routes = core.Routes{
	core.Route{
		Name:    "Find One User",
		Method:  "GET",
		Path:    basePath + "/{id}",
		Auth:    true,
		Roles:   roles.Roles{roles.Admin},
		Handler: FindOne,
	},
}
