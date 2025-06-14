package entity

type PickUp struct {
	Date        string `json:"date" form:"date"`
	Time        string `json:"time" form:"time"`
	Address     string `json:"address" form:"address"`
	Coordinates string `json:"coordinates" form:"coordinates"`
	Amount      uint   `json:"amount" form:"amount"`
}
