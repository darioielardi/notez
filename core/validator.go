package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	
	"github.com/gorilla/schema"
)

type Query interface {
	Validate() error
}

type DTO interface {
	Validate() error
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

var decoder = schema.NewDecoder()

func (v *Validator) ValidateBody(dto DTO, r *http.Request) error {
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	
	if err := json.Unmarshal(body, &dto); err != nil {
		return errors.New("invalid empty body")
	}
	
	if err := dto.Validate(); err != nil {
		return err
	}
	
	return nil
}

func (v *Validator) ValidateQuery(query Query, r *http.Request) error {
	
	q := r.URL.Query()
	
	if len(q) == 0 {
		return nil
	}

	if err := decoder.Decode(query, r.URL.Query()); err != nil {
		return err
	}
	
	if err := query.Validate(); err != nil {
		return err
	}

	return nil
}
