package models

import (
	"github.com/ikaro/usuarios/lib"
)

type Usuarios struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Email string `db:"email" json:"email`
}

var UsuarioModel = lib.Sess.Collection("usuarios")