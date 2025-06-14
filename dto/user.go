package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_REGISTER_USER = "failed register user"
	MESSAGE_FAILED_GET_USER      = "failed get user"
	MESSAGE_FAILED_UPDATE_USER   = "failed update user"
	MESSAGE_FAILED_LOGIN         = "failed login"
	MESSAGE_FAILED_REDEEM_POINTS = "failed redeem points"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "success register user"
	MESSAGE_SUCCESS_GET_USER      = "success get user"
	MESSAGE_SUCCESS_UPDATE_USER   = "success update user"
	MESSAGE_SUCCESS_LOGIN         = "success login"
	MESSAGE_SUCCESS_REDEEM_POINTS = "success redeem points"
)

var (
	ErrRoleNotAllowed            = errors.New("denied access for \"%v\" role")
	ErrGetUserById               = errors.New("failed to get user by id")
	ErrUpdateUser                = errors.New("failed to update user")
	ErrEmailOrPhoneNumberMissing = errors.New("email or phone number is missing")
	ErrCredentialsNotMatched     = errors.New("credentials not matched")
	ErrUsernameAlreadyExists     = errors.New("username already exist")
	ErrCreateUser                = errors.New("failed to create user")
	ErrInvalidRoleName           = errors.New("invalid role name")
	ErrMismatchOldPassword       = errors.New("old password does not match")
	ErrNoPointsAvailable         = errors.New("no points available for redemption")
	ErrInvalidPointsAmount       = errors.New("invalid points amount for redemption")
	ErrInvalidGender             = errors.New("invalid gender value")
)

type (
	UserAuthRequest struct {
		Email       string `json:"email" form:"email"`
		PhoneNumber string `json:"phone_number" form:"phone_number"`
		Password    string `json:"password" form:"password" binding:"required"`
	}

	UserUpdateRequest struct {
		UserName    string `json:"username" form:"username"`
		OldPassword string `json:"old_password" form:"old_password"`
		NewPassword string `json:"new_password" form:"new_password"`
		Name        string `json:"name" form:"name"`
		Address     string `json:"address" form:"address"`
		Role        string `json:"role" form:"role"`
		Gender      string `json:"gender" form:"gender"`
	}

	UserRedeemRequest struct {
		Bank    string `json:"bank" form:"bank" binding:"required"`
		Account string `json:"account" form:"account" binding:"required"`
		Nominal *uint  `json:"nominal" form:"nominal" binding:"required"`
	}

	UserResponse struct {
		ID          string `json:"id" form:"id"`
		UserName    string `json:"username" form:"username" gorm:"uniqueIndex;not null"`
		Name        string `json:"name" form:"name"`
		Address     string `json:"address" form:"address"`
		Email       string `json:"email" form:"email"`
		PhoneNumber string `json:"phone_number" form:"phone_number"`
		Points      uint   `json:"points" form:"points"`
		Gender      string `json:"gender" form:"gender"`

		RoleID   string `json:"role_id,omitempty" form:"role_id"`
		Role     string `json:"role" form:"role"`
		Verified bool   `json:"verified" form:"verified"`
	}
)
