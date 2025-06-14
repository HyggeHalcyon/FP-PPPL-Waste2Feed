package entity

import (
	"Waste2Feed/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID     uuid.UUID `json:"id" form:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()" `
	RoleID string    `json:"role_id" form:"role_id" gorm:"foreignKey" `

	UserName    string `json:"username" form:"username"`
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Password    string `json:"password" form:"password"`
	Verified    bool   `json:"verified" form:"verified"`
	Points      *uint  `json:"points" form:"points" gorm:"default:0"`
	Gender      string `json:"gender" form:"gender"`

	Role *Role `json:"role,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" `

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
