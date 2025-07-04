package controller

import (
	"net/http"

	"Waste2Feed/config"
	"Waste2Feed/constants"
	"Waste2Feed/dto"
	"Waste2Feed/entity"
	"Waste2Feed/service"
	"Waste2Feed/utils"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Update(ctx *gin.Context)
		Login(ctx *gin.Context)
		Me(ctx *gin.Context)
		Redeem(ctx *gin.Context)

		NewCourier(ctx *gin.Context)
		DeleteCourier(ctx *gin.Context)
		EditCourier(ctx *gin.Context)

		ViewCourier(ctx *gin.Context)
		ViewIndex(ctx *gin.Context)
		ViewLogin(ctx *gin.Context)
		ViewRegister(ctx *gin.Context)
		ViewChooseRole(ctx *gin.Context)
		ViewDashboard(ctx *gin.Context)
		ViewRedeem(ctx *gin.Context)
		ViewProfile(ctx *gin.Context)
		ViewFAQ(ctx *gin.Context)
		RenderAddCourierPage(c *gin.Context)
	}

	userController struct {
		jwtService  config.JWTService
		userService service.UserService
	}
)

func NewUserController(us service.UserService, jwt config.JWTService) UserController {
	return &userController{
		jwtService:  jwt,
		userService: us,
	}
}

func (c *userController) NewCourier(ctx *gin.Context) {
	var req dto.UserAuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.NewCourier(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_COURIER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_COURIER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) DeleteCourier(ctx *gin.Context) {
	courierID := ctx.Query("id")

	if err := c.userService.DeleteCourier(ctx.Request.Context(), courierID); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_COURIER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_COURIER, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) EditCourier(ctx *gin.Context) {
	courierID := ctx.Query("id")

	var req dto.UserUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.userService.EditCourier(ctx.Request.Context(), courierID, req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_COURIER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_COURIER, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Register(ctx *gin.Context) {
	role := ctx.Query("role")

	var req dto.UserAuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.Register(ctx.Request.Context(), role, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Update(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	var req dto.UserUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err := c.userService.Update(ctx.Request.Context(), userID, req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserAuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.Login(ctx.Request.Context(), req)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token := c.jwtService.GenerateToken(user.ID.String(), user.Role.Name)
	userResponse := entity.Authorization{
		Token: token,
		Role:  user.Role.Name,
	}

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, userResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) Me(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	result, err := c.userService.Me(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Redeem(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	var req dto.UserRedeemRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err := c.userService.Redeem(ctx, userID, req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REDEEM_POINTS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REDEEM_POINTS, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) ViewIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func (c *userController) ViewLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func (c *userController) ViewRegister(ctx *gin.Context) {
	role := ctx.Query("role")
	if role == "" || (role != constants.ENUM_ROLE_USER && role != constants.ENUM_ROLE_FARMER) {
		ctx.Redirect(http.StatusSeeOther, "/choose-role")
		return
	}

	ctx.HTML(http.StatusOK, "register.tmpl", gin.H{})
}

func (c *userController) ViewChooseRole(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "choose-role.tmpl", gin.H{})
}

func (c *userController) ViewDashboard(ctx *gin.Context) {
	role := ctx.MustGet(constants.CTX_KEY_ROLE_NAME).(string)

	if role == constants.ENUM_ROLE_USER {
		ctx.HTML(http.StatusOK, "dashboard-user.tmpl", gin.H{})
	} else if role == constants.ENUM_ROLE_FARMER {
		ctx.HTML(http.StatusOK, "dashboard-farmer.tmpl", gin.H{})
	} else if role == constants.ENUM_ROLE_ADMIN {
		users, err := c.userService.GetAllByRoleID(ctx.Request.Context(), constants.ENUM_ROLE_COURIER)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "dashboard-admin.tmpl", gin.H{
			"Users": users,
		})
	}
}

func (c *userController) ViewRedeem(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "redeem.tmpl", gin.H{})
}

func (c *userController) ViewProfile(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "profile.tmpl", gin.H{})
}

func (c *userController) ViewFAQ(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "faq.tmpl", gin.H{})
}

func (c *userController) ViewCourier(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}

	courier, err := c.userService.GetCourierByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"Error": err.Error(),
		})
	}

	ctx.HTML(http.StatusOK, "courier.tmpl", gin.H{
		"Courier": courier,
	})
}

func (c *userController) RenderAddCourierPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "courier-add.tmpl", gin.H{
		"Title": "Add Courier",
	})
}
