package entity

import "github.com/google/uuid"

type Pickup struct {
	ID        uuid.UUID  `json:"id" form:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    string     `json:"user_id" form:"user_id" gorm:"foreignKey"`
	CourierID *uuid.UUID `json:"courier_id" form:"courier_id" gorm:"foreignKey;type:uuid;default:NULL"`

	Date        string  `json:"date" form:"date"`
	Time        string  `json:"time" form:"time"`
	Address     string  `json:"address" form:"address"`
	Coordinates string  `json:"coordinates" form:"coordinates"`
	Amount      float64 `json:"amount" form:"amount"`
	Status      string  `json:"status" form:"status" gorm:"default:'progress'"`

	User    *User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Courier *User `json:"courier,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Timestamp
}
