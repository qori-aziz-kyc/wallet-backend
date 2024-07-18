package service

import (
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindOneBy(criteria map[string]interface{}) (*models.User, error)
	Create(models []*models.User, tx *gorm.DB) ([]*models.User, error)
	FindBy(criteria map[string]interface{}) ([]*models.User, error)
	Update(model *models.User, tx *gorm.DB) error
	Delete(models []*models.User, tx *gorm.DB) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserRepo {
	return userService{db}
}

func (svc userService) FindOneBy(criteria map[string]interface{}) (*models.User, error) {
	m := &models.User{}
	result := svc.db.Where(criteria).First(&m)
	if err := result.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (svc userService) Create(models []*models.User, tx *gorm.DB) ([]*models.User, error) {
	err := tx.Create(models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (svc userService) FindBy(criteria map[string]interface{}) ([]*models.User, error) {
	m := []*models.User{}
	result := svc.db.Where(criteria).Find(&m)
	if err := result.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (svc userService) Update(model *models.User, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (svc userService) Delete(models []*models.User, tx *gorm.DB) error {
	err := tx.Delete(&models, tx).Error
	if err != nil {
		return err
	}

	return nil
}
