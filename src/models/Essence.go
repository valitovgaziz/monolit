package models

import "github.com/google/uuid"

type Essence struct {
	Id               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;AutoIncrement:false"`
	Name             string    `json:"name" gorm:"type:string"`
	Type             string    `json:"type" gorm:"type:string"`
	Contact          Contact   `json:"contact" gorm:"type:struct"`
	ShortDescription string    `json:"shortDesc" gorm:"type:string"`
	Description      string    `json:"description" gorm:"type:string"`
	AverageBill      int       `json:"number" gorm:"type:int"`
}
