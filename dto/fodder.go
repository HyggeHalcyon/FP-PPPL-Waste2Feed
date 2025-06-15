package dto

import "errors"

const (
	MESSAGE_SUCCESS_CREATE_FODDER  = "Successfully created fodder request"
	MESSAGE_SUCCESS_PAYMENT_FODDER = "Successfully paid fodder request"

	MESSAGE_FAILED_CREATE_FODDER  = "Failed to create fodder request"
	MESSAGE_FAILED_PAYMENT_FODDER = "Failed to pay fodder request"
)

var (
	ErrInvalidFodderType         = errors.New("invalid fodder type")
	ErrFodderNotOwnedByRequester = errors.New("fodder not owned by requester")
	ErrFodderAlreadyPaid         = errors.New("fodder already paid")
)

type (
	FodderRequest struct {
		Address     string  `json:"address" form:"address" binding:"required"`
		Coordinates string  `json:"coordinates" form:"coordinates"`
		Amount      float64 `json:"amount" form:"amount" binding:"required"`
		Date        string  `json:"date" form:"date" binding:"required"`
		Time        string  `json:"time" form:"time" binding:"required"`
	}

	FodderPaymentRequest struct {
		ID      string   `json:"id" form:"id" binding:"required"`
		Bank    string   `json:"bank" form:"bank" binding:"required"`
		Account string   `json:"account" form:"account" binding:"required"`
		Nominal *float64 `json:"nominal" form:"nominal" binding:"required"`
	}

	FodderResponse struct {
		ID        string `json:"id" form:"id"`
		FarmerID  string `json:"farmer_id" form:"farmer_id"`
		CourierID string `json:"courier_id" form:"courier_id"`

		Date        string  `json:"date" form:"date"`
		Time        string  `json:"time" form:"time"`
		Address     string  `json:"address" form:"address"`
		Coordinates string  `json:"coordinates" form:"coordinates"`
		Amount      float64 `json:"amount" form:"amount"`
		Status      string  `json:"status" form:"status" gorm:"default:'progress'"`
		Type        string  `json:"type" form:"type"`
		Paid        bool    `json:"paid" form:"paid" gorm:"default:false"`
	}
)
