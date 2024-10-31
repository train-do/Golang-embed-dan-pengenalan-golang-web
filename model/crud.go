package model

import "database/sql"

type CRUD interface {
	Create(*sql.Tx) (int, error)
	Read()
	Update(*sql.Tx)
	Delete(*sql.Tx)
}

func CreateInterface(s CRUD, tx *sql.Tx) (int, error) {
	return s.Create(tx)
}
