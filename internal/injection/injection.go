package injection

import (
	categoryusecase "github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/categoryUsecase"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/handlers/categoryhandler"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/jwt"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/service"
	"gorm.io/gorm"
)

type HandlerInjection struct {
	Category categoryhandler.CategoryHandler
}

func NewInitialInjection(jwt jwt.JWTService, DB *gorm.DB) HandlerInjection {
	categoryRepo := service.NewCategoryService(DB)
	return HandlerInjection{
		Category: categoryhandler.NewCategoryHandler(categoryusecase.NewCategoryUsecase(jwt, DB, categoryRepo)),
	}
}
