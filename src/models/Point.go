package models

import "github.com/google/uuid"

type Point struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;AutoIncrement:false"`
	Latitude  int64     `json:"latitude" gorm:"type:int64"`
	Longitude int64     `json:"longitude" gorm:"type:int64"`
}
