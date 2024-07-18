package recordusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/service"
	"gorm.io/gorm"
)

type RecordUsecase interface {
	Create(ctx context.Context, req request.RecordCreateRequest) (int, *models.Record, error)
	Find(ctx context.Context, req request.RecordFindRequest) (int, []*models.Record, error)
	Update(ctx context.Context, req request.RecordCreateRequest) (int, *models.Record, error)
}

type recordUsecase struct {
	db         *gorm.DB
	recordRepo service.RecordRepo
}

func NewRecordUsecase(db *gorm.DB, recordRepo service.RecordRepo) RecordUsecase {
	return &recordUsecase{
		db:         db,
		recordRepo: recordRepo,
	}
}
