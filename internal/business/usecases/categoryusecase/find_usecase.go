package categoryusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
)

func (uc *categoryUsecase) Find(ctx context.Context, req request.CategoryRequest) (int, []*models.Category, error) {
	criteria := make(map[string]interface{}, 0)
	if len(req.CategoryID) > 0 {
		criteria["id"] = req.CategoryID
	}
	cat, err := uc.categoryRepo.FindBy(criteria)
	if err != nil {
		return 422, nil, err
	}

	return 200, cat, nil
}
