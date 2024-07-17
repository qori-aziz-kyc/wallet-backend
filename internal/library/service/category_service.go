package service

import (
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	FindOneBy(criteria map[string]interface{}) (*models.Category, error)
	Create(models []*models.Category, tx *gorm.DB) ([]*models.Category, error)
	FindBy(criteria map[string]interface{}) ([]*models.Category, error)
	Update(model *models.Category, tx *gorm.DB) error
	Delete(models []*models.Category, tx *gorm.DB) error
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

func (svc categoryService) Create(models []*models.Category, tx *gorm.DB) ([]*models.Category, error) {
	err := tx.Create(models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (svc categoryService) FindBy(criteria map[string]interface{}) ([]*models.Category, error) {
	m := []*models.Category{}
	result := svc.db.Where(criteria).Find(&m)
	if err := result.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (svc categoryService) Update(model *models.Category, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (svc categoryService) Delete(models []*models.Category, tx *gorm.DB) error {
	err := tx.Delete(&models, tx).Error
	if err != nil {
		return err
	}

	return nil
}
