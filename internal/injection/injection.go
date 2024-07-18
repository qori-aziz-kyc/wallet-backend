package injection

import (
	"github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/categoryusecase"
	"github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/recordusecase"
	"github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/userusecase"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/handlers/categoryhandler"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/handlers/recordhandler"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/handlers/userhandler"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/jwt"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/service"
	"gorm.io/gorm"
)

type HandlerInjection struct {
	Category categoryhandler.CategoryHandler
	User     userhandler.UserHandler
	Record   recordhandler.RecordHandler
}

func NewInitialInjection(jwt jwt.JWTService, DB *gorm.DB) HandlerInjection {
	categoryRepo := service.NewCategoryService(DB)
	userRepo := service.NewUserService(DB)
	recordRepo := service.NewRecordService(DB)
	return HandlerInjection{
		Category: categoryhandler.NewCategoryHandler(categoryusecase.NewCategoryUsecase(jwt, DB, categoryRepo)),
		User:     userhandler.NewUserHandler(userusecase.NewUserUsecase(jwt, DB, userRepo)),
		Record:   recordhandler.NewRecordHandler(recordusecase.NewRecordUsecase(DB, recordRepo)),
	}
}
