package db

import (
	"errors"

	"github.com/Lezard82/movies-api/src/infrastructure/logger"
	"gorm.io/gorm"
)

type GormDBAdapter struct {
	DB *gorm.DB
}

func NewGormDBAdapter(db *gorm.DB) *GormDBAdapter {
	return &GormDBAdapter{DB: db}
}

func (g *GormDBAdapter) Find(dest interface{}, filters map[string]interface{}) error {
	log := logger.GetLogger()
	log.WithField("filters", filters).Info("Executing Find")

	query := g.DB.Model(dest)

	allowedFields := map[string]bool{
		"title":        true,
		"release_date": true,
		"genre":        true,
	}

	for field, value := range filters {
		if !allowedFields[field] {
			errMsg := "invalid filter field"
			log.WithField("filters", filters).Error("Error executing Find: " + errMsg)
			return errors.New(errMsg)
		}
		query = query.Where(field+" = ?", value)
	}

	err := query.Find(dest).Error
	if err != nil {
		log.WithField("error", err).Error("Error executing Find")
		return err
	}

	log.Info("Find executed successfully")
	return nil
}

func (g *GormDBAdapter) First(dest interface{}, id int64) error {
	err := g.DB.First(dest, id).Error
	logError("First", err, map[string]interface{}{"id": id})
	return err
}

func (g *GormDBAdapter) FirstByField(dest interface{}, field string, value interface{}) error {
	err := g.DB.Where(field+" = ?", value).First(dest).Error
	logError("FirstByField", err, map[string]interface{}{"field": field, "value": value})
	return err
}

func (g *GormDBAdapter) Create(value interface{}) error {
	err := g.DB.Create(value).Error
	logError("Create", err, nil)
	return err
}

func (g *GormDBAdapter) Save(value interface{}) error {
	err := g.DB.Save(value).Error
	logError("Save", err, nil)
	return err
}

func (g *GormDBAdapter) Delete(value interface{}, id int64) error {
	err := g.DB.Delete(value, id).Error
	logError("Delete", err, map[string]interface{}{"id": id})
	return err
}

func (g *GormDBAdapter) CountByFields(model any, conditions map[string]interface{}, excludeID int64) (int64, error) {
	var count int64

	query := g.DB.Model(model)
	for field, value := range conditions {
		query = query.Where(field+" = ?", value)
	}

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	err := query.Count(&count).Error
	logError("CountByFields", err, nil)
	return count, err
}

func logError(operation string, err error, details map[string]interface{}) {
	if err != nil {
		log := logger.GetLogger()
		entry := log.WithField("error", err).WithField("operation", operation)
		for key, value := range details {
			entry = entry.WithField(key, value)
		}
		entry.Error("Database operation failed")
	}
}
