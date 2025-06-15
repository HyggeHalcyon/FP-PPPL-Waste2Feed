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
	FodderController interface {
		Create(ctx *gin.Context)
		Pay(ctx *gin.Context)

		ViewPaymentPage(ctx *gin.Context)
		ViewPendingPayments(ctx *gin.Context)
		ViewFodder(ctx *gin.Context)
		ViewOrder(ctx *gin.Context)
		ViewFodderInformation(ctx *gin.Context)
		ViewNutrition(ctx *gin.Context)
		ViewDosage(ctx *gin.Context)
	}

	fodderController struct {
		jwtService    config.JWTService
		fodderService service.FodderService
	}
)

func NewFodderController(fs service.FodderService, jwt config.JWTService) FodderController {
	return &fodderController{
		jwtService:    jwt,
		fodderService: fs,
	}
}

func (c *fodderController) Create(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)
	_type := ctx.Query("type")

	var req dto.FodderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.fodderService.Order(ctx.Request.Context(), userID, _type, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_FODDER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_FAILED_CREATE_FODDER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *fodderController) Pay(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	var req dto.FodderPaymentRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.fodderService.Payment(ctx.Request.Context(), userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PAYMENT_FODDER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_PAYMENT_FODDER, result)
	ctx.JSON(http.StatusOK, res)

}

func (c *fodderController) ViewPaymentPage(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)
	id := ctx.Query("id")

	if id == "" {
		ctx.Redirect(http.StatusSeeOther, "/fodder/pending")
		return
	}

	result, err := c.fodderService.GetPendingDetails(ctx.Request.Context(), userID, id)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "fodder-payment.tmpl", gin.H{
		"Details": result,
	})
}

func (c *fodderController) ViewPendingPayments(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	result, err := c.fodderService.GetPendingPayments(ctx.Request.Context(), userID)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Error": err.Error(),
		})
	}

	ctx.HTML(http.StatusOK, "fodder-pending.tmpl", gin.H{
		"PendingPayments": result,
	})
}

func (c *fodderController) ViewFodder(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "fodder.tmpl", gin.H{})
}

func (c *fodderController) ViewOrder(ctx *gin.Context) {
	_type := ctx.Query("type")

	if _type == "" {
		ctx.Redirect(http.StatusSeeOther, "/fodder")
		return
	}

	if _type != constants.ENUM_FODDER_TYPE_CHICKEN && _type != constants.ENUM_FODDER_TYPE_FISH {
		ctx.Redirect(http.StatusSeeOther, "/fodder")
		return
	}

	ctx.HTML(http.StatusOK, "fodder-order.tmpl", gin.H{})
}

func (c *fodderController) ViewFodderInformation(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "fodder-information.tmpl", gin.H{})
}

func (c *fodderController) ViewNutrition(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "fodder-nutrition.tmpl", gin.H{})
}

func (c *fodderController) ViewDosage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "fodder-dosage.tmpl", gin.H{})
}
