package repository

import (
	"github.com/fathimasithara01/ecommerce/database"
)

type PgSQLRepository struct {
}

var IPgSQLrepo PgSQLRepoInterface

func PgSQLInit() {
	IPgSQLrepo = &PgSQLRepository{}

}

func (r *PgSQLRepository) Insert(obj interface{}) error {
	return database.PgSQLDB.Create(obj).Error
}
