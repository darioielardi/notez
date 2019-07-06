package notez

import (
	"notez/core"
	roles "notez/utils/enums"
)

var basePath = "/notez"

var Routes = core.Routes{
	core.Route{
		Name:    "Get My Notez",
		Method:  "GET",
		Path:    basePath,
		Auth:    true,
		Roles:   roles.Roles{roles.Admin, roles.User},
		Handler: GetByUser,
	},
	core.Route{
		Name:    "Get One",
		Method:  "GET",
		Path:    basePath + "/{id}",
		Auth:    true,
		Roles:   roles.Roles{roles.Admin, roles.User},
		Handler: GetOne,
	},
	core.Route{
		Name:    "Create New Note",
		Method:  "POST",
		Path:    basePath,
		Auth:    true,
		Roles:   roles.Roles{roles.Admin, roles.User},
		Handler: CreateNew,
	},
}
