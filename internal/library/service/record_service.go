package service

import (
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"gorm.io/gorm"
)

type RecordRepo interface {
	FindOneBy(criteria map[string]interface{}) (*models.Record, error)
	Create(models []*models.Record, tx *gorm.DB) ([]*models.Record, error)
	FindBy(criteria map[string]interface{}) ([]*models.Record, error)
	Update(model *models.Record, tx *gorm.DB) error
	Delete(models []*models.Record, tx *gorm.DB) error
}

type recordService struct {
	db *gorm.DB
}

func NewRecordService(db *gorm.DB) RecordRepo {
	return &recordService{db}
}

func (svc *recordService) FindOneBy(criteria map[string]interface{}) (*models.Record, error) {
	m := &models.Record{}
	result := svc.db.Where(criteria).First(&m)
	if err := result.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *recordService) Create(models []*models.Record, tx *gorm.DB) ([]*models.Record, error) {
	err := tx.Create(models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (svc *recordService) FindBy(criteria map[string]interface{}) ([]*models.Record, error) {
	m := []*models.Record{}
	result := svc.db.Where(criteria).Find(&m)
	if err := result.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *recordService) Update(model *models.Record, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (svc *recordService) Delete(models []*models.Record, tx *gorm.DB) error {
	err := tx.Delete(&models, tx).Error
	if err != nil {
		return err
	}

	return nil
}
