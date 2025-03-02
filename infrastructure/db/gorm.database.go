package db

import (
	"errors"

	"github.com/Lezard82/movies-api/infrastructure/logger"
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
			return errors.New("invalid filter field")
		}
		query = query.Where(field+" = ?", value)
	}

	err := query.Find(dest).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *GormDBAdapter) First(dest interface{}, id int64) error {
	return g.DB.First(dest, id).Error
}

func (g *GormDBAdapter) FirstByField(dest interface{}, field string, value interface{}) error {
	return g.DB.Where(field+" = ?", value).First(dest).Error
}

func (g *GormDBAdapter) Create(value interface{}) error {
	return g.DB.Create(value).Error
}

func (g *GormDBAdapter) Save(value interface{}) error {
	return g.DB.Save(value).Error
}

func (g *GormDBAdapter) Delete(value interface{}, id int64) error {
	return g.DB.Delete(value, id).Error
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
	return count, err
}
