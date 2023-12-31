package models

import (
	"time"

	"github.com/google/uuid"
)

type LiveConsumption struct {
	HomeId                      uuid.UUID `gorm:"type:symbol" json:"homeId"`
	Timestamp                   time.Time `gorm:"type:timestamp" json:"timestamp"`
	Power                       float64   `gorm:"type:double" json:"power"`
	MinPower                    float64   `gorm:"type:double" json:"minPower"`
	MaxPower                    float64   `gorm:"type:double" json:"maxPower"`
	AveragePower                float64   `gorm:"type:double" json:"averagePower"`
	LastMeterConsumption        float64   `gorm:"type:double" json:"lastMeterConsuption"`
	LastMeterProduction         float64   `gorm:"type:double" json:"lastMeterProduction"`
	AccumulatedConsumptionToday float64   `gorm:"type:double" json:"accumulatedConsumptionToday"`
	AccumulatedProductionToday  float64   `gorm:"type:double" json:"accumulatedProductionToday"`
	AccumulatedConsumptionHour  float64   `gorm:"type:double" json:"accumulatedConsumptionHour"`
	AccumulatedProductionHour   float64   `gorm:"type:double" json:"accumulatedProductionHour"`
	CurrentPrice                float64   `gorm:"type:double" json:"CurrentPrice"`
	AccumulatedCostToday        float64   `gorm:"type:double" json:"accumulatedCostToday"`
}

type ElectricityPrice struct {
	HomeId          uuid.UUID `gorm:"type:symbol" json:"homeId"`
	Timestamp       time.Time `gorm:"type:timestamp" json:"timestamp"`
	TotalByProvider float64   `gorm:"type:double" json:"totalProvider"`
	Spot            float64   `gorm:"type:double" json:"energy"`
	Tax             float64   `gorm:"type:double" json:"tax"`
	Calculated      float64   `gorm:"type:double" json:"calculated"`
	Grid            float64   `gorm:"type:double" json:"grid"`
	Currency        string    `gorm:"type:symbol" json:"currency"`
}
