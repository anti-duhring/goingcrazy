package schema

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Apelido    string
	Nome       string
	Nascimento time.Time
	Stack      []string
}
