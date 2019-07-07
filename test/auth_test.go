package test

import (
	"testing"
)

func TestGetMe(t *testing.T) {

	e := NewExpect(t, "/auth")

	// Unauthenticated => 401
	e.GET("/me").
		Expect().
		Status(401)

	utr := GetUserTestToken()

	// User auth => 200, User Data
	e.GET("/me").
		WithHeader("Authorization", "Bearer "+utr.Token).
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("data").
		Value("data").
		Object().
		ContainsKey("auth_id").
		ValueEqual("auth_id", utr.UID)

	atr := GetAdminTestToken()

	// Admin auth => 200, User Data
	e.GET("/me").
		WithHeader("Authorization", "Bearer "+atr.Token).
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("data").
		Value("data").
		Object().
		ContainsKey("auth_id").
		ValueEqual("auth_id", atr.UID)

}
