package models

import "gorm.io/gorm"

type Mascota struct {
	gorm.Model

	Nombre string `gorm:"not null"`
	Especie string `gorm:"not null"`
	Raza string `gorm:"not null"`
	Genero string `gorm:"not null"`
	Peso int `gorm:"not null"`
	Temperatura int `gorm:"not null"`
	Observacion string
	MascotaID uint
	PropietarioID uint `gorm:"foreignKey:PropietarioID"`
	
}