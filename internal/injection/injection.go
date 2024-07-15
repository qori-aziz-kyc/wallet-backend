package injection

import (
	categoryusecase "github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/categoryUsecase"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/handlers/categoryhandler"
	"github.com/qori-aziz-kyc/wallet-backend/library/jwt"
)

type HandlerInjection struct {
	Category categoryhandler.CategoryHandler
}

func NewInitialInjection(jwt jwt.JWTService) HandlerInjection {
	return HandlerInjection{
		Category: categoryhandler.NewCategoryHandler(categoryusecase.NewCategoryUsecase(jwt)),
	}
}
