package test

import (
	"testing"
)

func TestFindOne(t *testing.T) {
	
	e := NewExpect(t, "/users")
	
	e.GET("/1").
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("data").
		Value("data").
		Object().
		ContainsKey("id").
		ValueEqual("id", 1)
	
}
