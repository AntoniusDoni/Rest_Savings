package services

import (
	"github.com/Rest_Savings/apps/repository"
	"github.com/Rest_Savings/pkg/database"
	"github.com/Rest_Savings/pkg/validator"
)

type Services struct {
	Db         *database.GormDB
	Validator  *validator.Validator
	Repository *repository.Repository
}

func New() *Services {

	validator := validator.New()
	gormdb := database.New()
	repository := repository.New(gormdb)

	return &Services{
		Db:         gormdb,
		Validator:  validator,
		Repository: repository,
	}
}
