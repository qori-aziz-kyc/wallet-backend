package categoryusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
)

func (uc *categoryUsecase) Create(ctx context.Context, req request.CategoryRequest) (int, *models.Category, error) {
	category := &models.Category{
		Name:  req.Name,
		Color: req.Color,
	}

	tx := uc.db.Begin()
	defer tx.Rollback()
	cat, err := uc.categoryRepo.Create([]*models.Category{category}, tx)
	if err != nil {
		return 422, nil, err
	}

	tx.Commit()
	return 200, cat[0], nil
}
