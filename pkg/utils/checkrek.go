package utils

import (
	"github.com/Rest_Savings/apps/models"
	"github.com/Rest_Savings/apps/services"
	"github.com/Rest_Savings/pkg/request"
	"github.com/gofiber/fiber/v2"
)

func New() *AuthorizeUser {
	return &AuthorizeUser{}
}

type AuthorizeUser struct {
	Ser *services.Services
}

func (authorizeUser *AuthorizeUser) CheckRek(ctx *fiber.Ctx) error {
	acctReq := new(request.CreateTrasaction)
	ctx.BodyParser(acctReq)
	acc := new(models.Account)
	check, err := authorizeUser.Ser.Repository.GetDetailAccont(acctReq.NoRek, 0, acc)
	if check > 0 && err == nil {
		return ctx.Next()
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "No Rekening tidak ditemukan"})
	}

}
