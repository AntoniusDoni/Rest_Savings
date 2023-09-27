package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	NoRek      string       `gorm:"primaryKey;size:190;unique"`
	Nik        string       `gorm:"unique;size:190;index"`
	Name       string       `gorm:"index"`
	Phone      string       `gorm:"unique;size:13"`
	Trasaction []Trasaction `gorm:"foreignKey:NoRek"`
	CreatedAt  time.Time    `gorm:"autoCreateTime"`
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (acc *Account) BeforeCreate(tx *gorm.DB) (err error) {
	temp := fmt.Sprint(time.Now().Nanosecond())
	randNum := fmt.Sprintf("%s%s", "ACC-", temp[:5])
	acc.NoRek = randNum
	return
}
