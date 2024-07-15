package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/constants"
	"github.com/qori-aziz-kyc/wallet-backend/library/helper"
	"github.com/qori-aziz-kyc/wallet-backend/library/jwt"
)

type AuthMiddleware struct {
	jwtService jwt.JWTService
	isAdmin    bool
}

func NewAuthMiddleware(jwtService jwt.JWTService, isAdmin bool) gin.HandlerFunc {
	return (&AuthMiddleware{
		jwtService: jwtService,
		isAdmin:    isAdmin,
	}).Handle
}

func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		helper.NewAbortResponse(ctx, "missing authorization header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		helper.NewAbortResponse(ctx, "invalid header format")
		return
	}

	if headerParts[0] != "Bearer" {
		helper.NewAbortResponse(ctx, "token must content bearer")
		return
	}

	user, err := m.jwtService.ParseToken(headerParts[1])
	if err != nil {
		helper.NewAbortResponse(ctx, "invalid token")
		return
	}

	if user.IsAdmin != m.isAdmin && !user.IsAdmin {
		helper.NewAbortResponse(ctx, "you don't have access for this action")
		return
	}

	ctx.Set(constants.CtxAuthenticatedUserKey, user)
	ctx.Next()
}
