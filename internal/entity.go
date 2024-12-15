package internal

import "gorm.io/gorm"

type Developer struct {
	gorm.Model
	Nome string
	Email string
	Task  []Task `gorm:"foreignKey:DeveloperID"`
}

type Task struct {
	gorm.Model
	Nome string
	Descricao string
	Dificuldade string
	Status string
	DeveloperID uint64
}