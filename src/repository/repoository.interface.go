package repository

type PgSQLRepoInterface interface {
	Insert(obj interface{}) error
}
