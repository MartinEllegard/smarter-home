package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID         `gorm:"type:uuid;primeryKey" json:"id"`
	FirstName        string            `json:"firstName"`
	LastName         string            `json:"lastName"`
	Homes            []Home            `gorm:"foreignKey:user_id;references:id" json:"homes"`
	ElectricityDeals []ElectricityDeal `gorm:"foreignKey:user_id;references:id" json:"electricityDeals"`
}

type Home struct {
	ID           uuid.UUID `gorm:"type:uuid;primeryKey" json:"id"`
	UserId       uuid.UUID `gorm:"type:uuid" json:"userId"`
	ProviderType string    `gorm:"type:text" json:"providerType"`
	Name         string    `gorm:"type:text" json:"name"`
	WsSupport    bool      `gorm:"type:boolean" json:"wsSupport"`
}
