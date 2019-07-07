package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/gavv/httpexpect.v2"

	"notez/config"
	"notez/core"
	"notez/database"
	"notez/modules/auth"
	"notez/modules/notez"
	"notez/modules/users"
)

var conf = config.Init("../config/test.yml")

var server = core.NewServer(
	core.NewRouter(),
	database.NewDatabase(conf),
	conf,
)

func NewExpect(t *testing.T, path string) *httpexpect.Expect {

	routesGroups := []core.Routes{
		users.Routes,
		auth.Routes,
		notez.Routes,
	}

	var routes core.Routes

	for _, r := range routesGroups {
		routes = append(routes, r...)
	}

	server.Wire(routes)

	var s = httptest.NewServer(server.Router)

	e := httpexpect.New(t, s.URL+path)

	return e
}

type TokenResponse struct {
	Token string `json:"idToken"`
	UID   string `json:"localId"`
}

func GetTestToken(email string) *TokenResponse {

	reqBody, _ := json.Marshal(map[string]interface{}{
		"email":             email,
		"password":          conf.Firebase.TestPsw,
		"returnSecureToken": true,
	})

	res, err := http.Post(
		"https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key="+conf.Firebase.ApiKey,
		"application/json",
		bytes.NewBuffer(reqBody),
	)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	tokenRes := &TokenResponse{}

	if err := json.Unmarshal(body, tokenRes); err != nil {
		log.Fatalln(err)
	}

	return tokenRes
}

func GetUserTestToken() *TokenResponse {
	return GetTestToken("user@test.com")
}

func GetAdminTestToken() *TokenResponse {
	return GetTestToken("admin@test.com")
}
