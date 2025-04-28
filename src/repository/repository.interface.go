package repository

import "gorm.io/gorm"

type PgSQLRepoInterface interface {
	Insert(obj interface{}) error
	FindById(obj interface{}, id int) error
	FindByEmail(obj interface{}, email string) error
	Update(obj interface{}, id int, update interface{}) error
	Delete(obj interface{}, id int) error
	HardDelete(obj interface{}) error
	FindAll(obj interface{}) error
	FindAllWhere(obj interface{}, query interface{}, args ...interface{}) error
	UpdateByFields(obj interface{}, id int, fields map[string]interface{}) error
	InsertAndReturnID(req interface{}) (uint, error)
	FindDistinct(obj interface{}, fiels string, query interface{}, args ...interface{}) error
	Raw(sql string, values ...interface{}) *gorm.DB
	Save(req interface{}) error
}
