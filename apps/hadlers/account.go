package handlers

import (
	"fmt"

	"github.com/Rest_Savings/apps/models"
	"github.com/Rest_Savings/pkg/request"
	"github.com/Rest_Savings/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// Create Account godoc
//
//	@Summary		create Account
//	@Description	Create Account
//		@Accept			json
//		@Produce		json
//		@Param			payload	body		request.CreateAccountRequest	true	"Account Request"
//		@Success		200		{object}  response.CreateAccountResponse  "OK"
//		@Failure		400		{object}  response.CreateAccountResponse  "Bad Request"
//
// @Router /daftar [post]
func (hand *Handler) CreateAccount(ctx *fiber.Ctx) error {
	acctReq := new(request.CreateAccountRequest)
	err := ctx.BodyParser(acctReq)
	res := response.CreateAccountResponse{}
	status := fiber.StatusOK
	message := "Berhasil"
	if err != nil {
		status = fiber.StatusBadRequest
		message = "Gagal"
	}

	errs := hand.Ser.Validator.ValidateRequest(acctReq)
	fmt.Println(errs)
	if errs != nil {
		status = fiber.StatusBadRequest
		message = "Gagal"
		return ctx.Status(status).JSON(errs)

	}
	acc := new(models.Account)
	checknik, err := hand.Ser.Repository.GetDetailAccont(acctReq.NIK, 1, acc)
	if checknik > 0 && err == nil {
		res.NIK = "NIK Sudah digunakan"
		status = fiber.StatusBadRequest
		message = "Gagal"
	}
	checkphone, err := hand.Ser.Repository.GetDetailAccont(acctReq.Phone, 2, acc)
	if checkphone > 0 && err == nil {
		res.Phone = "No Handpone Sudah digunakan"
		status = fiber.StatusBadRequest
		message = "Gagal"
	}
	if status != 400 {
		result, row := hand.Ser.Repository.CreateAccount(acctReq)
		if row == 0 {
			status = fiber.StatusBadRequest
			message = "Gagal"
		} else {
			res.NoRek = result
		}
	}

	resp := response.Response{Message: message, Body: res}
	return ctx.Status(status).JSON(resp)
}
