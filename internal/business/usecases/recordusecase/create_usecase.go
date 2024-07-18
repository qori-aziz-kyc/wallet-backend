package recordusecase

import (
	"context"
	"fmt"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
)

func (uc *recordUsecase) Create(ctx context.Context, req request.RecordCreateRequest) (int, *models.Record, error) {
	if req.UserID == 0 {
		return 422, nil, fmt.Errorf("failed getting userID")
	}

	if req.Amount <= 0 {
		return 422, nil, fmt.Errorf("amount must more than 0")
	}

	record := &models.Record{
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Date:       req.Date,
		Type:       req.Type,
		Note:       req.Note,
	}

	tx := uc.db.Begin()
	defer tx.Rollback()
	records, err := uc.recordRepo.Create([]*models.Record{record}, tx)
	if err != nil {
		return 422, nil, fmt.Errorf("failed to create record : %v", err)
	}

	tx.Commit()

	return 200, records[0], nil
}
