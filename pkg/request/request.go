package request

type CreateAccountRequest struct {
	NIK   string `json:"nik" validate:"min=6,required"`
	Name  string `json:"nama" validate:"required"`
	Phone string `json:"no_hp" validate:"min=11,max=12,required"`
}

type CreateTrasaction struct {
	NoRek   string  `json:"no_rekening" validate:"required"`
	Nominal float64 `json:"nominal" validate:"required"`
}
