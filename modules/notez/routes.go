package notez

import (
	"notez/core"
	"notez/utils/enums"
)

var basePath = "/notez"

var Routes = core.Routes{
	core.Route{
		Name:    "Get My Notez",
		Method:  "GET",
		Path:    basePath,
		Auth:    true,
		Roles:   enums.Roles{"user", "admin"},
		Handler: GetByUser,
	},
	core.Route{
		Name:    "Get One",
		Method:  "GET",
		Path:    basePath + "/{id}",
		Auth:    true,
		Roles:   enums.Roles{"user", "admin"},
		Handler: GetOne,
	},
	core.Route{
		Name:    "Create New Note",
		Method:  "POST",
		Path:    basePath,
		Auth:    true,
		Roles:   enums.Roles{"user", "admin"},
		Handler: CreateNew,
	},
}
