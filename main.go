package main

import (
	"notez/config"
	"notez/core"
	db "notez/database"
	"notez/modules/auth"
	"notez/modules/notez"
	"notez/modules/users"
	"notez/utils"
)

func main() {
	
	conf := config.Init("./config/dev.yml")
	
	utils.InitFirebase()
	
	server := core.NewServer(
		core.NewRouter(),
		db.NewDatabase(conf),
		conf,
	)
	
	rgs := []core.Routes{
		users.Routes,
		auth.Routes,
		notez.Routes,
	}
	
	server.Init(rgs)
}
