package service

import (
	"dibantuin-be/entity"
	"dibantuin-be/repository"
	"errors"
	"strings"
)

type CategoryService struct {
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryService(repository *repository.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: repository}
}

func generateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

func (cs *CategoryService) CreateCategory(category *entity.CategoryRequest) (*entity.Category, error) {
	if category.Name == "" {
		return nil, errors.New("category name is required")
	}
	newSlug := generateSlug(category.Name)

	newCategory := &entity.Category{
		Name: category.Name,
		Slug: newSlug,
	}
	if err := cs.CategoryRepository.AddCategory(newCategory); err != nil {
		return nil, errors.New("failed to create category")
	}

	return newCategory, nil
}

func (cs *CategoryService) UpdateCategory(id uint64, updated *entity.CategoryRequest) (*entity.Category, error) {
	category, err := cs.CategoryRepository.GetCategoryById(id)
	if err != nil {
		return nil, errors.New("category id was not found")
	}

	category.Name = updated.Name
	category.Slug = generateSlug(updated.Name)

	err = cs.CategoryRepository.UpdateCategory(category)
	if err != nil {
		return nil, errors.New("failed to update category")
	}

	return category, nil
}

func (cs *CategoryService) GetCategories() (*[]entity.Category, error) {
	return cs.CategoryRepository.GetCategories()
}

func (cs *CategoryService) DeleteCategory(id uint64) error {
	return cs.CategoryRepository.DeleteCategory(id)
}
