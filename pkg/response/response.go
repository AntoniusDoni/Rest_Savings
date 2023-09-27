package response

type CreateAccountResponse struct {
	NIK   string `json:"nik,omitempty" `
	Phone string `json:"no_hp,omitempty" `
	NoRek string `json:"no_rekening,omitempty"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
}

type GetListTransaction struct {
	TransactionTime string  `json:"waktu,omitempty"`
	TransactionCode string  `json:"kode_transaksi,omitempty"`
	Balance         float64 `json:"nominal"`
}
