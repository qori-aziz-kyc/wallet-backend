package categoryusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/library/jwt"
)

type CategoryUsecase interface {
	Create(ctx context.Context, categoryRequest request.CategoryRequest) (statusCode int, err error)
}

type categoryUsecase struct {
	jwtService jwt.JWTService
}

func NewCategoryUsecase(jwtService jwt.JWTService) CategoryUsecase {
	return &categoryUsecase{
		jwtService: jwtService,
	}
}
