package repository

import (
	"errors"
	"log"
	"reflect"

	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/database"
)

type PgSQLRepository struct{}

// Global repository instance
var IPgSQLRepo *PgSQLRepository

// Initialize Repository
func PgSQLInit() {
	IPgSQLRepo = &PgSQLRepository{}
}

/*** Inserting Data ***/
func (r *PgSQLRepository) Insert(req interface{}) error {
	return database.PgSQLDB.Create(req).Error
}

func (r *PgSQLRepository) Save(req interface{}) error {
	return database.PgSQLDB.Save(req).Error
}

func (r *PgSQLRepository) InsertAndReturnID(req interface{}) (uint, error) {
	if err := database.PgSQLDB.Create(req).Error; err != nil {
		return 0, err
	}

	// Extract ID dynamically using reflection
	value := reflect.ValueOf(req).Elem()
	idField := value.FieldByName("ID")
	if !idField.IsValid() {
		return 0, errors.New("ID field not found")
	}

	return uint(idField.Uint()), nil
}

/*** Fetching Data ***/
func (r *PgSQLRepository) FindById(obj interface{}, id int) error {
	return database.PgSQLDB.Where("id = ?", id).First(obj).Error
}

/*** Fetching Data ***/
func (r *PgSQLRepository) FindByEmail(obj interface{}, email string) error {
	return database.PgSQLDB.Where("email = ? ", email).First(obj).Error
}
func (r *PgSQLRepository) FindAll(obj interface{}) error {
	return database.PgSQLDB.Find(obj).Error
}

func (r *PgSQLRepository) FindOneWhere(out interface{}, query string, args ...interface{}) error {
	return database.PgSQLDB.Where(query, args...).First(out).Error
}

func (r *PgSQLRepository) FindAllWhere(obj interface{}, query interface{}, args ...interface{}) error {
	return database.PgSQLDB.Where(query, args...).Find(obj).Error
}

/*** Fetching with Sorting ***/
func (r *PgSQLRepository) FindAllSorted(obj interface{}, orderBy string, ascending bool) error {
	query := database.PgSQLDB.Order(orderBy)
	if !ascending {
		query = query.Order(orderBy + " DESC")
	}
	return query.Find(obj).Error
}

/*** Fetching with Pagination ***/
func (r *PgSQLRepository) FindPaginated(obj interface{}, page, pageSize int, orderBy string) error {
	offset := (page - 1) * pageSize
	return database.PgSQLDB.Order(orderBy).Offset(offset).Limit(pageSize).Find(obj).Error
}

/*** Updating Data ***/
func (r *PgSQLRepository) Update(obj interface{}, id int, update interface{}) error {
	log.Println("Updating object:", obj)
	return database.PgSQLDB.Where("id = ?", id).First(obj).Updates(update).Error
}

func (r *PgSQLRepository) UpdateByFields(obj interface{}, id int, fields map[string]interface{}) error {
	log.Println("Updating fields:", fields)
	return database.PgSQLDB.Model(obj).Where("id = ?", id).Updates(fields).Error
}

/*** Deleting Data ***/
func (r *PgSQLRepository) Delete(obj interface{}, id int) error {
	return database.PgSQLDB.Where("id = ?", id).First(obj).Delete(obj).Error
}

func (r *PgSQLRepository) HardDelete(obj interface{}) error {
	return database.PgSQLDB.Unscoped().Delete(obj).Error
}

/*** Fetching Distinct Values ***/
func (r *PgSQLRepository) FindDistinct(obj interface{}, field string, query interface{}, args ...interface{}) error {
	return database.PgSQLDB.Model(obj).Distinct(field).Where(query, args...).Find(obj).Error
}

/*** Execute Raw SQL Queries ***/
func (r *PgSQLRepository) Raw(query string, args ...interface{}) *gorm.DB {
	return database.PgSQLDB.Raw(query, args...)
}
