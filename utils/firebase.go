package utils

import (
	"context"
	"log"
	
	firebase "firebase.google.com/go"
)

var app *firebase.App

// InitFirebase initialize a new firebase admin app instance and must be called in main
func InitFirebase() *firebase.App {
	fbApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error initializing firebase app: %v\n", err)
	}
	app = fbApp
	return app
}

// GetFb gets the current firebase app instance
func GetFb() *firebase.App {
	if app == nil {
		return InitFirebase()
	}
	return app
}
