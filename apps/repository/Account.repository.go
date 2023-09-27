package repository

import (
	"github.com/Rest_Savings/apps/models"
	"github.com/Rest_Savings/pkg/request"
)

func (rep *Repository) CreateAccount(acct *request.CreateAccountRequest) (string, int64) {
	db, _ := rep.Gormdb.GetInstance()
	acc := new(models.Account)
	acc.Nik = acct.NIK
	acc.Name = acct.Name
	acc.Phone = acct.Phone
	result := db.Create(&acc)
	close, _ := db.DB()
	close.Close()
	return acc.NoRek, result.RowsAffected
}

func (rep *Repository) GetDetailAccont(search string, types int8, acct *models.Account) (int64, error) {
	db, _ := rep.Gormdb.GetInstance()
	query := db.Model(&models.Account{})
	if types == 1 {
		query = query.Where("nik=?", search)
	} else if types == 2 {
		query = query.Where("phone=?", search)
	} else {
		query = query.Where("no_rek=?", search)
	}
	query = query.First(&acct)
	row := query.RowsAffected
	err := query.Error
	close, _ := db.DB()
	close.Close()
	return row, err
}
func (rep *Repository) GetListAccount(search string) ([]models.Account, error) {
	db, _ := rep.Gormdb.GetInstance()
	var prov []models.Account
	query := db.Model(&models.Account{})
	if search != "" {
		query = query.Where("no_rek like ?", search+"%")
	}
	err := query.Find(&prov).Error
	close, _ := db.DB()
	close.Close()
	return prov, err
}
