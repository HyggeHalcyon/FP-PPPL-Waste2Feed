package entity

import "github.com/google/uuid"

type (
	Fodder struct {
		ID        uuid.UUID  `json:"id" form:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
		FarmerID  string     `json:"farmer_id" form:"farmer_id" gorm:"foreignKey"`
		CourierID *uuid.UUID `json:"courier_id" form:"courier_id" gorm:"foreignKey;type:uuid;default:NULL"`

		Date        string  `json:"date" form:"date"`
		Time        string  `json:"time" form:"time"`
		Address     string  `json:"address" form:"address"`
		Coordinates string  `json:"coordinates" form:"coordinates"`
		Amount      float64 `json:"amount" form:"amount"`
		Status      string  `json:"status" form:"status" gorm:"default:'progress'"`
		Type        string  `json:"type" form:"type"`
		Paid        bool    `json:"paid" form:"paid" gorm:"default:false"`

		Farmer  *User `json:"farmer,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Courier *User `json:"courier,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

		Timestamp
	}
)
