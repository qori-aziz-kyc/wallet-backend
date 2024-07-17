package categoryusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
)

func (uc *categoryUsecase) Update(ctx context.Context, req request.CategoryRequest) (int, *models.Category, error) {

	// find category
	category, err := uc.categoryRepo.FindOneBy(map[string]interface{}{"id": req.ID})
	if err != nil {
		return 422, nil, nil
	}

	category.Color = req.Color
	category.Name = req.Name

	tx := uc.db.Begin()
	defer tx.Rollback()
	err = uc.categoryRepo.Update(category, tx)
	if err != nil {
		return 422, nil, err
	}

	tx.Commit()
	return 200, category, nil
}
