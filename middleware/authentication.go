package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/dto"
	"Waste2Feed/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticateBearer(jwtService config.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, dto.ErrTokenNotFound.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			abortTokenInvalid(ctx)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		userID, userRole, err := jwtService.GetPayloadInsideToken(authHeader)
		if err != nil {
			if err.Error() == dto.ErrTokenExpired.Error() {
				response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, dto.ErrTokenExpired.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set(constants.CTX_KEY_TOKEN, authHeader)
		ctx.Set(constants.CTX_KEY_USER_ID, userID)
		ctx.Set(constants.CTX_KEY_ROLE_NAME, userRole)
		ctx.Next()
	}
}

func AuthenticateCookies(jwtService config.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieToken, err := ctx.Cookie("jwt")
		if err != nil || strings.TrimSpace(cookieToken) == "" {
			ctx.SetCookie("jwt", "", -1, "/", "", false, true)
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
			return
		}

		if cookieToken == "" {
			ctx.SetCookie("jwt", "", -1, "/", "", false, true)
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
			return
		}

		userID, userRole, err := jwtService.GetPayloadInsideToken(cookieToken)
		if err != nil {
			ctx.SetCookie("jwt", "", -1, "/", "", false, true)
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
			return
		}

		fmt.Println("REACHED")
		ctx.Set(constants.CTX_KEY_TOKEN, cookieToken)
		ctx.Set(constants.CTX_KEY_USER_ID, userID)
		ctx.Set(constants.CTX_KEY_ROLE_NAME, userRole)
		ctx.Next()
	}
}

func abortTokenInvalid(ctx *gin.Context) {
	response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_TOKEN, dto.ErrTokenInvalid.Error(), nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
