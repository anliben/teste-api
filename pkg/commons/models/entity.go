package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Address      string `json:"address"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state"`
	Document     string `json:"document"`
	Cpf          string `json:"cpf"`
	ImportantOf  string `json:"important_of"`
	ReferenceOf  string `json:"reference_of"`
	ReceivedOf   string `json:"received_of"`
	Value        string `json:"value"`
}
