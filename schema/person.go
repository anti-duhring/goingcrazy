package schema

import (
	"fmt"
	"strings"
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

func (t *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(`"2006-01-02"`)), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		date, err = time.Parse("2006-01-02T15:04:05Z", strings.Trim(string(b), "\""))
		if err != nil {
			return fmt.Errorf("error on converting date: %w", err)
		}
	}
	t.Time = date
	return
}

type Person struct {
	Base
	Apelido     string     `gorm:"size:32;not null;unique" json:"apelido"`
	Nome        string     `gorm:"size:100;not null" json:"nome"`
	SearchIndex string     `gorm:"not null"`
	Nascimento  CustomTime `gorm:"type:date;not null" json:"nascimento"`
	Stack       []string   `gorm:"type:jsonb" json:"stack"`
}

func (r *Person) Validate(c *gin.Context) error {
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
