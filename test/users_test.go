package test

import (
	"testing"
)

func TestFindOne(t *testing.T) {

	e := NewExpect(t, "/users")

	// Unauthenticated => 401
	e.GET("/1").
		Expect().
		Status(401)

	utr := GetUserTestToken()

	// User auth => 403
	e.GET("/1").
		WithHeader("Authorization", "Bearer "+utr.Token).
		Expect().
		Status(403)

	atr := GetAdminTestToken()

	// Admin auth => 200, User Data
	e.GET("/1").
		WithHeader("Authorization", "Bearer "+atr.Token).
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("data").
		Value("data").
		Object().
		ContainsKey("id").
		ValueEqual("id", 1)

	// Non existing user id => 404
	e.GET("/-1").
		WithHeader("Authorization", "Bearer "+atr.Token).
		Expect().
		Status(404)

	// Invalid user id => 400
	e.GET("/abc").
		WithHeader("Authorization", "Bearer "+atr.Token).
		Expect().
		Status(400)

}
