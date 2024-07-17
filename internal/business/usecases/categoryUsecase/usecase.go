package categoryusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/jwt"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/service"
	"gorm.io/gorm"
)

type CategoryUsecase interface {
	Create(ctx context.Context, req request.CategoryRequest) (int, *models.Category, error)
	Update(ctx context.Context, req request.CategoryRequest) (int, *models.Category, error)
	Find(ctx context.Context, req request.CategoryRequest) (int, []*models.Category, error)
}

type categoryUsecase struct {
	jwtService   jwt.JWTService
	db           *gorm.DB
	categoryRepo service.CategoryRepo
}

func NewCategoryUsecase(jwtService jwt.JWTService, db *gorm.DB, categoryRepo service.CategoryRepo) CategoryUsecase {
	return &categoryUsecase{
		jwtService:   jwtService,
		db:           db,
		categoryRepo: categoryRepo,
	}
}
