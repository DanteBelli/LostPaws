package models

import (
	"time"

	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Id               uint      `json:"id"`
	User_id          uint      `json:"user_id"`
	Description      string    `json:"description"`
	Location         string    `json:"location"`
	ImageUrl         string    `json:"image_url"`
	Status           bool      `json:"status"`
	Name             string    `json:"name"`
	Raza_id          int8      `json:"raza_id"`
	Tamanio          int8      `json:"tamanio"`
	Color_prim       int       `json:"color_prim"`
	Color_Sec        int       `json:"color_sec"`
	Color_Terc       int       `json:"color_terc"`
	Fecha_enc        time.Time `json:"fecha_enc"`
	Tipo_masc        int8      `json:"tipo_masc"`
	Sexo             int8      `json:"sexo"`
	Castrado         bool      `json:"castrado"`
	SociablePerro    bool      `json:"sociable_perro" gorm:"type:boolean"`
	SociablePersonas bool      `json:"sociable_personas" gorm:"type:boolean"`
	Ubic_Actual      string    `json:"ubic_actual"`
}
