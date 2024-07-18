package recordusecase

import (
	"context"
	"fmt"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
)

func (uc *recordUsecase) Update(ctx context.Context, req request.RecordCreateRequest) (int, *models.Record, error) {
	if req.UserID == 0 {
		return 422, nil, fmt.Errorf("failed getting userID")
	}

	// find category
	record, err := uc.recordRepo.FindOneBy(map[string]interface{}{
		"id":      req.ID,
		"user_id": req.UserID,
	})
	if err != nil {
		return 422, nil, nil
	}

	record.Amount = req.Amount
	record.CategoryID = req.CategoryID
	record.Date = req.Date
	record.Note = req.Note
	record.Type = req.Type

	tx := uc.db.Begin()
	defer tx.Rollback()
	err = uc.recordRepo.Update(record, tx)
	if err != nil {
		return 422, nil, err
	}

	tx.Commit()
	return 200, record, nil
}
