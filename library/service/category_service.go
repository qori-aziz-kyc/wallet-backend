package service

import (
	"github.com/qori-aziz-kyc/wallet-backend/library/models"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	FindOneBy(criteria map[string]interface{}) (*models.Category, error)
}

type categoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryRepo {
	return categoryService{db}
}

func (svc categoryService) FindOneBy(criteria map[string]interface{}) (*models.Category, error) {
	m := &models.Category{}
	result := svc.db.Where(criteria).First(&m)
	if err := result.Error; err != nil {
		return nil, err
	}

	return m, nil
}
