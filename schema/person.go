package schema

import (
	"gorm.io/datatypes"
)

type Person struct {
	Base
	Apelido    string         `gorm:"size:32;not null;unique"`
	Nome       string         `gorm:"size:100;not null"`
	Nascimento datatypes.Date `gorm:"type:date;not null"`
	Stack      datatypes.JSON
}
