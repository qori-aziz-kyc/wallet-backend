package userusecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/response"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/helper"
)

func (uc *userUsecase) Login(ctx context.Context, req request.RegisterRequest) (int, *response.LoginResponse, error) {
	// Search Username
	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"username": req.Username,
	})
	if err != nil {
		return 422, nil, fmt.Errorf("failed to login, user not found")
	}

	isSame := helper.ValidateHash(req.Password, user.Password)
	if !isSame {
		return 422, nil, fmt.Errorf("password is wrong")
	}

	accessToken, err := uc.jwtService.GenerateToken(strconv.Itoa(user.ID), user.IsAdmin)
	if err != nil {
		return 422, nil, fmt.Errorf("failed generate token : %v", err)
	}

	resp := &response.LoginResponse{
		UserID:      user.ID,
		AccessToken: accessToken,
	}

	return 200, resp, nil
}
