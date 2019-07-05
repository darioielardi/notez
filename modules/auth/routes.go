package auth

import (
	"notez/core"
	"notez/utils/enums"
)

var basePath = "/auth"

var Routes = core.Routes{
	core.Route{
		Name:    "Get Me",
		Method:  "GET",
		Path:    basePath + "/me",
		Auth:    true,
		Roles:   enums.Roles{"admin", "user"},
		Handler: GetMe,
	},
}
