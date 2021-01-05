package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mashingan/smapping"
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

type TerapiRequestByRange struct {
	StarDate time.Time `json:"star_date"`
	EndDate  time.Time `json:"end_date"`
}

func (b TerapiRequestByRange) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.StarDate, validation.Required, validation.Date(time.ANSIC)), //2021-01-02T19:49:58.828478+08:00
		validation.Field(&b.EndDate, validation.Required, validation.Date(time.ANSIC)),
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

func (b TerapiRequest) TranslateToEntity() (*Terapi, error) {
	terapiEntity := Terapi{}
	err := smapping.FillStruct(&terapiEntity, smapping.MapFields(&b))
	if err != nil {
		return nil, err
	}
	return &terapiEntity, err
}

func (b Terapi) TranslateToRes() (*TerapiResponse, error) {
	terapiResponse := TerapiResponse{}
	err := smapping.FillStruct(&terapiResponse, smapping.MapFields(&b))
	if err != nil {
		return nil, err
	}
	return &terapiResponse, err
}
