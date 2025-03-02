package db

type Database interface {
	Find(dest interface{}) error
	First(dest interface{}, id int64) error
	FirstByField(dest interface{}, field string, value interface{}) error
	CountByFields(model any, conditions map[string]interface{}, excludeID int64) (int64, error)
	Create(value interface{}) error
	Save(value interface{}) error
	Delete(value interface{}, id int64) error
}
