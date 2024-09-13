package models

import "github.com/google/uuid"

type Contact struct {
	Id      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;AutoIncrement:false"`
	Email   string    `json:"email" gorm:"type:string"`
	Phone   string    `json:"phone" gorm:"type:string"`
	Address string    `json:"address" gorm:"type:string"`
	Point   Point     `json:"point" gorm:"type:struct"`
}
