package terapi

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type Terapi struct {
	ID        uint      `json:"id"`
	PasienID  uint      `json:"pasien_id"`
	PegawaiID uint      `json:"pegawai_id"`
	Tglterapi time.Time `json:"tglterapi"`
	Layanan   string    `json:"layanan"`
	Biaya     float64   `json:"biaya"`
	Upah      float64   `json:"upah"`
}

type TerapiRequest struct {
	PasienID uint    `json:"pasien_id"`
	Layanan  string  `json:"layanan"`
	Biaya    float64 `json:"biaya"`
}

func (b TerapiRequest) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.PasienID, validation.Required),
		validation.Field(&b.Layanan, validation.Required),
	)
}

type TerapiResponse struct {
	ID        uint      `json:"id"`
	PasienID  uint      `json:"pasien_id"`
	PegawaiID uint      `json:"pegawai_id"`
	Tglterapi time.Time `json:"tglterapi"`
	Layanan   string    `json:"layanan"`
	Biaya     float64   `json:"biaya"`
	Upah      float64   `json:"upah"`
}
