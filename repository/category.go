package repository

import (
	"dibantuin-be/entity"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (cr *CategoryRepository) AddCategory(category *entity.Category) error {
	return cr.DB.Create(category).Error
}

func (cr *CategoryRepository) GetCategories() (*[]entity.Category, error) {
	var categories []entity.Category
	if err := cr.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (cr *CategoryRepository) GetCategoryById(id uint64) (*entity.Category, error) {
	var category entity.Category
	if err := cr.DB.Find(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (cr *CategoryRepository) UpdateCategory(category *entity.Category) error {
	return cr.DB.Save(category).Error
}

func (cr *CategoryRepository) DeleteCategory(id uint64) error {
	return cr.DB.Delete(&entity.Category{}, id).Error
}
