package repository

import (
	"fmt"
	"time"

	"github.com/Rest_Savings/apps/models"
	"github.com/Rest_Savings/pkg/request"
)

func (rep *Repository) CreateTransaction(transType string, trx *request.CreateTrasaction) (int64, float64) {
	db, _ := rep.Gormdb.GetInstance()
	trxAdd := new(models.Trasaction)
	rep.GetBalace(trx.NoRek, trxAdd)
	temp := fmt.Sprint(time.Now().Nanosecond())
	randNum := fmt.Sprintf("%s%s", transType, temp[:5])
	trxAdd.TransactionCode = randNum
	trxAdd.NoRek = trx.NoRek
	trxAdd.Status = transType
	trxAdd.Nominal = trx.Nominal
	var row int64

	if transType == "C" {
		trxAdd.Balance = trxAdd.Balance + trx.Nominal
	} else {
		trxAdd.Balance = trxAdd.Balance - trx.Nominal
	}
	result := db.Create(&trxAdd)
	row = result.RowsAffected

	close, _ := db.DB()
	close.Close()
	return row, trxAdd.Balance
}

func (rep *Repository) GetBalace(no_rek string, trx *models.Trasaction) error {
	db, _ := rep.Gormdb.GetInstance()
	err := db.Select("balance").Where("no_rek = ?", no_rek).Order("created_at DESC").Last(&trx).Error
	close, _ := db.DB()
	close.Close()
	return err
}
