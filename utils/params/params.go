package params

import (
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

func GetIdParam(r *http.Request) (*int, error) {
	
	idP := mux.Vars(r)["id"]
	
	id, err := strconv.Atoi(idP)
	if err != nil {
		return nil, err
	}
	
	return &id, nil
}
