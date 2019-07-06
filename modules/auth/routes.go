package auth

import (
	"notez/core"
	roles "notez/utils/enums"
)

var basePath = "/auth"

var Routes = core.Routes{
	core.Route{
		Name:    "Get Me",
		Method:  "GET",
		Path:    basePath + "/me",
		Auth:    true,
		Roles:   roles.Roles{roles.User, roles.Admin},
		Handler: GetMe,
	},
}
