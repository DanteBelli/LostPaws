package models

import (
	"time"

	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	UserId            uint      `json:"user_id"`
	Description       string    `json:"description"`
	ImageUrl          string    `json:"image_url"`
	Location          string    `json:"location"`
	Status            bool      `json:"status"`
	Nombre            string    `json:"nombre"`
	RazaId            bool      `json:"raza"`
	Tamanio           int8      `json:"tamanio"`
	Color_Principal   int       `json:"color_principal"`
	Color_Secundario  int       `json:"color_secundario"`
	Color_Tercero     int       `json:"color_tercero"`
	Dir               string    `string:"dir"`
	FechaSucesi       time.Time `json:"Fecha"`
	Tipo_Mascota      int8      `json:"TipoMascota"`
	Sexo              int8      `json:"sexo"`
	Estado            int8      `json:"estado"`
	Castrado          bool      `json:"castrado"`
	Sociable_Perros   bool      `json:"sociablePerro"`
	Sociable_Personas bool      `json:"sociablePersona"`
	UbicacionActual   string    `json:"UbicacionActual"`
}
