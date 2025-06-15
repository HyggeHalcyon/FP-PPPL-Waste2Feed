package dto

const (
	MESSAGE_FAILED_CREATE_PICKUP = "failed create pick-up"
	MESSAGE_FAILED_GET_PICKUP    = "failed get pick-up"

	MESSAGE_SUCCESS_CREATE_PICKUP = "success create pick-up"
	MESSAGE_SUCCESS_GET_PICKUP    = "success get pick-up"
)

type (
	PickupCreateRequest struct {
		Address     string  `json:"address" form:"address" binding:"required"`
		Coordinates string  `json:"coordinates" form:"coordinates"`
		Amount      float64 `json:"amount" form:"amount" binding:"required"`
		Date        string  `json:"date" form:"date" binding:"required"`
		Time        string  `json:"time" form:"time" binding:"required"`
	}

	PickupResponse struct {
		ID        string `json:"id" form:"id"`
		UserID    string `json:"user_id" form:"user_id"`
		CourierID string `json:"courier_id" form:"courier_id"`

		Date        string  `json:"date" form:"date" binding:"required"`
		Time        string  `json:"time" form:"time" binding:"required"`
		Address     string  `json:"address" form:"address" binding:"required"`
		Coordinates string  `json:"coordinates" form:"coordinates" binding:"required"`
		Amount      float64 `json:"amount" form:"amount" binding:"required"`
		Status      string  `json:"status" form:"status" binding:"required"`
	}

	PickupPaginationResponse struct {
		Data []PickupResponse `json:"data"`
		PaginationMetadata
	}
)
