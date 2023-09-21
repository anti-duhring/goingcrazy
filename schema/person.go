package schema

import (
	"time"

	"gorm.io/datatypes"
)

type Person struct {
	Base
	Apelido    string         `gorm:"size:32;not null"`
	Nome       string         `gorm:"size:100;not null"`
	Nascimento datatypes.Date `gorm:"type:date;not null"`
	Stack      datatypes.JSON
}

type PersonResponse struct {
	ID         string    `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento time.Time `json:"nascimento"`
	Stack      []string  `json:"stack"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
}
