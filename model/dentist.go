package model

import (
	_ "github.com/godror/godror"
)

type Dentist struct {
	ID       uint32 `db:"id" json:"dentist_id"`
	Name     string `db:"nombre" json:"name"`
	LastName string `db:"apellido" json:"last_name"`
	DNI      string `db:"matricula" json:"dni"`
}
