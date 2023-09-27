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
