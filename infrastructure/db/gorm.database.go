package db

import (
	"gorm.io/gorm"
)

type GormDBAdapter struct {
	DB *gorm.DB
}

func NewGormDBAdapter(db *gorm.DB) *GormDBAdapter {
	return &GormDBAdapter{DB: db}
}

func (g *GormDBAdapter) Find(dest interface{}, filters map[string]interface{}) error {
	query := g.DB.Model(dest)

	for field, value := range filters {
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
