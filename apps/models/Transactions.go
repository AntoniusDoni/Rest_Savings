package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Trasaction struct {
	Id              string `gorm:"primaryKey;"`
	TransactionCode string
	NoRek           string `gorm:"size:190"`
	Nominal         float64
	Status          string
	Ref             string
	TrasactionDate  datatypes.Date
	Balance         float64
	gorm.Model
}

func (trx *Trasaction) BeforeCreate(tx *gorm.DB) (err error) {
	trx.Id = uuid.New().String()
	currentTime := time.Now()
	trx.TrasactionDate = datatypes.Date(currentTime)
	return
}
