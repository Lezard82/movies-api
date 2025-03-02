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

func (g *GormDBAdapter) Find(dest interface{}) error {
	return g.DB.Find(dest).Error
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
