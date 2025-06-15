package controller

import (
	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/dto"
	"Waste2Feed/service"
	"Waste2Feed/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	PickupController interface {
		Create(ctx *gin.Context)

		ViewHistory(ctx *gin.Context)
		ViewPickup(ctx *gin.Context)
	}

	pickupController struct {
		jwtService    config.JWTService
		pickupService service.PickupService
	}
)

func NewPickupController(ps service.PickupService, jwt config.JWTService) PickupController {
	return &pickupController{
		jwtService:    jwt,
		pickupService: ps,
	}
}

func (c *pickupController) Create(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	var req dto.PickupCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.pickupService.Create(ctx.Request.Context(), userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PICKUP, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PICKUP, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *pickupController) ViewHistory(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)
	role := ctx.MustGet(constants.CTX_KEY_ROLE_NAME).(string)

	var req dto.PaginationQuery
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Error": err.Error(),
		})
		return
	}

	result, err := c.pickupService.GetPaginated(ctx.Request.Context(), userID, role, req)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "history.tmpl", gin.H{
		"Data": result,
	})
}

func (c *pickupController) ViewPickup(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pick-up.tmpl", gin.H{})
}
