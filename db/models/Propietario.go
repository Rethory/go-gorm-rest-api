package models

import "gorm.io/gorm"

type Propietario struct {
	gorm.Model

	PNombre string `gorm:"not null" json:"nombre"`
	PApellido string `gorm:"not null" json:"apellido"`
	Telefono int `gorm:"not null;unique_index" json:"telefono"`
	Mascotas []Mascota `json:"mascotas"`

}