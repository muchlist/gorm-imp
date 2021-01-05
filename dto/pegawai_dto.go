package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mashingan/smapping"
)

type Pegawai struct {
	ID           uint
	Nama         string
	Kontak       string
	Username     string
	Password     string
	Level        int           // 0 user , 1 admin, 2 superadmin
	Terapis      []Terapi      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PegawaiID"`
	Pengeluarans []Pengeluaran `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PegawaiID"`
}

type PegawaiRequest struct {
	Nama     string `json:"nama"`
	Kontak   string `json:"kontak"`
	Username string `json:"username"`
	Password string `json:"password"`
	Level    int    `json:"level"`
}

type PegawaiResponse struct {
	ID           uint   `json:"id"`
	Nama         string `json:"nama"`
	Kontak       string `json:"kontak"`
	Username     string `json:"username"`
	Level        int    `json:"level"`
	Terapis      []Terapi
	Pengeluarans []Pengeluaran
}

func (b PegawaiRequest) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Nama, validation.Required),
		validation.Field(&b.Kontak, validation.Required),
		validation.Field(&b.Username, validation.Required),
		validation.Field(&b.Password, validation.Required),
		validation.Field(&b.Level, validation.Max(2)),
	)
}

func (b Pegawai) TranslateToResponse() (*PegawaiResponse, error) {
	pegawaiResponse := PegawaiResponse{}
	err := smapping.FillStruct(&pegawaiResponse, smapping.MapFields(&b))
	if err != nil {
		return nil, err
	}
	return &pegawaiResponse, err
}

func (b PegawaiRequest) TranslateToEntity() (*Pegawai, error) {
	pegawaiEntity := Pegawai{}
	err := smapping.FillStruct(&pegawaiEntity, smapping.MapFields(&b))
	if err != nil {
		return nil, err
	}
	return &pegawaiEntity, err
}
