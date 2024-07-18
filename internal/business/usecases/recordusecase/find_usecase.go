package recordusecase

import (
	"context"
	"fmt"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
)

func (uc *recordUsecase) Find(ctx context.Context, req request.RecordFindRequest) (int, []*models.Record, error) {
	if req.UserID == 0 {
		return 422, nil, fmt.Errorf("failed getting userID")
	}

	criteria := make(map[string]interface{}, 0)
	if len(req.CategoryID) > 0 {
		criteria["category_id"] = req.CategoryID
	}

	criteria["user_id"] = req.UserID

	records, err := uc.recordRepo.FindBy(criteria)
	if err != nil {
		return 422, nil, err
	}

	return 200, records, nil
}
