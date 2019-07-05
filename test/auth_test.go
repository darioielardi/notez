package test

import (
	"testing"
)

func TestGetMe(t *testing.T) {
	
	e := NewExpect(t, "/auth")
	
	// Unauthorized
	e.GET("/me").
		Expect().
		Status(401)
	
	tr := GetTestTokenRes()
	
	e.GET("/me").
		WithHeader("Authorization", "Bearer "+tr.Token).
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("data").
		Value("data").
		Object().
		ContainsKey("auth_id").
		ValueEqual("auth_id", tr.UID)
	
}
