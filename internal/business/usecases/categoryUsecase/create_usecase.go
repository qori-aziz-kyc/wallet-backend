package categoryusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
)

func (usecae *categoryUsecase) Create(ctx context.Context, categoryRequest request.CategoryRequest) (statusCode int, err error) {
	return 200, nil
}
