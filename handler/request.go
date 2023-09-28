package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

type CustomTime struct {
	time.Time
}

var dateFormat = "2006-01-02"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ct.Format(dateFormat))), nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	t, err := time.Parse(dateFormat, s[1:len(s)-1]) // Remove quotes
	if err != nil {
		return err
	}

	*ct = CustomTime{t}
	return nil
}

type CreatePersonRequest struct {
	Apelido    string     `json:"apelido"`
	Nome       string     `json:"nome"`
	Nascimento CustomTime `json:"nascimento"`
	Stack      []string   `json:"stack"`
}

func (r *CreatePersonRequest) Validate(c *gin.Context) error {
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
