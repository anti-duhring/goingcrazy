package handler

import (
	"fmt"
	"time"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errParamMaxLength(name, typ string, max int) error {
	return fmt.Errorf("param: %s (type: %s) max length: %d", name, typ, max)
}

func hasStringWithMaxLength(s []string, max int) bool {

	for _, v := range s {
		if len(v) > max {
			return true
		}
	}

	return false
}

type CreatePersonRequest struct {
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento time.Time `json:"nascimento"`
	Stack      []string  `json:"stack"`
}

func (r *CreatePersonRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("malformed request body")
	}

	if r.Nome == "" {
		return errParamIsRequired("nome", "string")
	}
	if len(r.Nome) > 100 {
		return errParamMaxLength("nome", "string", 100)
	}

	if r.Apelido == "" {
		return errParamIsRequired("apelido", "string")
	}
	if len(r.Apelido) > 32 {
		return errParamMaxLength("apelido", "string", 32)
	}

	if hasStringWithMaxLength(r.Stack, 32) {
		return errParamMaxLength("stack", "string", 32)
	}

	return nil
}
