package handlers

import (
	"github.com/Rest_Savings/apps/models"
	"github.com/Rest_Savings/pkg/request"
	"github.com/Rest_Savings/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// Create Adds Saving godoc
//
//	@Summary		create Add Transaction Saving
//	@Description	Create Add Transaction Saving
//		@Accept			json
//		@Produce		json
//		@Param			payload	body		request.CreateTrasaction	true	"Account Request"
//		@Success		200		{object}  response.Response  "OK"
//		@Failure		400		{object}  response.Response  "Bad Request"
//
// @Router /tabung [post]
func (hand *Handler) CreateTractionAdds(ctx *fiber.Ctx) error {
	acctReq := new(request.CreateTrasaction)
	err := ctx.BodyParser(acctReq)
	res := response.Response{}
	status := fiber.StatusOK
	if err != nil {
		res.Message = "Gagal"
		status = fiber.StatusBadRequest
	}
	row, balace := hand.Ser.Repository.CreateTransaction("C", acctReq)
	if row == 0 {

		res.Message = "Gagal"
		status = fiber.StatusBadRequest
	} else {
		res.Message = "Berhasil"
		res.Body = fiber.Map{"Saldo": balace}
	}
	return ctx.Status(status).JSON(res)
}

// Create Widraw Saving godoc
//
//	@Summary		create Add Transaction Widraw
//	@Description	Create Add Transaction Widraw
//		@Accept			json
//		@Produce		json
//		@Param			payload	body		request.CreateTrasaction	true	"Account Request"
//		@Success		200		{object}  response.Response  "OK"
//		@Failure		400		{object}  response.Response  "Bad Request"
//
// @Router /tarik [post]
func (hand *Handler) CreateTractionWidraw(ctx *fiber.Ctx) error {
	acctReq := new(request.CreateTrasaction)
	err := ctx.BodyParser(acctReq)
	res := response.Response{}
	status := fiber.StatusOK
	if err != nil {
		res.Message = "Gagal"
		status = fiber.StatusBadRequest
	}
	tmpTrx := new(models.Trasaction)
	err = hand.Ser.Repository.GetBalace(acctReq.NoRek, tmpTrx)
	if err != nil {
		res.Message = "Gagal"
		status = fiber.StatusBadRequest
	}

	if tmpTrx.Balance >= acctReq.Nominal {
		row, balace := hand.Ser.Repository.CreateTransaction("D", acctReq)
		if row == 0 {
			res.Message = "Gagal"
			status = fiber.StatusBadRequest
		} else {
			res.Message = "Berhasil"
			res.Body = fiber.Map{"Saldo": balace}
		}
	} else {
		res.Message = "Saldo Tidak Cukup"
		res.Body = fiber.Map{"Saldo": tmpTrx.Balance}
		status = fiber.StatusBadRequest
	}

	return ctx.Status(status).JSON(res)
}

// Saldo godoc
//
//	@Summary		Get Detail Saldo
//	@Description	Get Detail of Saldo
//	@Accept			json
//	@Produce		json
//	@Param			no_rekening	path		string	false	"no_rekening"
//
//		@Success		200		{object}  response.Response  "OK"
//		@Failure		400		{object}  response.Response  "Bad Request"
//	 @Failure		404		{object}  response.Response  "Province Not Found"
//
// @Router /saldo/{no_rekening} [Get]
func (hand *Handler) GetBalaceAccount(ctx *fiber.Ctx) error {
	no_rekening := ctx.Params("no_rekening")
	res := response.Response{}
	status := fiber.StatusOK
	tmpTrx := new(models.Trasaction)
	err := hand.Ser.Repository.GetBalace(no_rekening, tmpTrx)
	if err != nil {
		res.Message = "Gagal"
		status = fiber.StatusBadRequest
		res.Body = "No Rekening tidak di temukan"
	} else {
		res.Body = fiber.Map{"Saldo": tmpTrx.Balance}
	}

	return ctx.Status(status).JSON(res)
}

// Mutasi godoc
//
//	@Summary		Get List Transaction Account
//	@Description	Get List of Transaction Account
//	@Accept			json
//	@Produce		json
//	@Param			no_rekening	path		string	false	"no_rekening"
//
//		@Success		200		{object}  response.Response  "OK"
//		@Failure		400		{object}  response.Response  "Bad Request"
//	 @Failure		404		{object}  response.Response  "Province Not Found"
//
// @Router /mutasi/{no_rekening} [Get]
func (hand *Handler) GetListTransaction(ctx *fiber.Ctx) error {
	no_rekening := ctx.Params("no_rekening")
	res := response.Response{}
	status := fiber.StatusOK

	result, err := hand.Ser.Repository.GetListTransaction(no_rekening)
	if err == nil {
		bodyRes := []response.GetListTransaction{}
		layout := "2006-01-02 15:04:05"
		for _, attr := range result {
			bodyRes = append(bodyRes, response.GetListTransaction{
				TransactionTime: attr.CreatedAt.Format(layout),
				TransactionCode: attr.TransactionCode,
				Balance:         attr.Nominal,
			})
		}
		res.Body = bodyRes
	} else {
		res.Message = "Gagal"
		status = fiber.StatusBadRequest
		res.Body = "No Rekening tidak di temukan"
	}

	return ctx.Status(status).JSON(res)
}
