package schema

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Apelido    string    `gorm:"size:32;not null"`
	Nome       string    `gorm:"size:100;not null"`
	Nascimento time.Time `gorm:"type:date;not null"`
	Stack      *[]string `gorm:"type:json"`
}

type PersonResponse struct {
	ID         uint           `json: "id"`
	Apelido    string         `json: "apelido"`
	Nome       string         `json: "nome"`
	Nascimento time.Time      `json: "nascimento"`
	Stack      []string       `json: "stack"`
	CreatedAt  time.Time      `json: "createdAt"`
	UpdatedAt  time.Time      `json: "updatedAt"`
	DeletedAt  gorm.DeletedAt `json: "deletedAt, omitempty"`
}
