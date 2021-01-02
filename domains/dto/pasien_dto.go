package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Pasien struct {
	ID           uint
	NoPasien     string   `json:"no_pasien"`
	Nama         string   `json:"nama"`
	Jk           int      `json:"jk"` //0 perempuan 1 laki-laki
	NoHp         string   `json:"no_hp"`
	NoWa         string   `json:"no_wa"`
	Alamat       string   `json:"alamat"`
	Terapis      []Terapi `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PasienID" json:"terapis"`
	JumlahTerapi int      `json:"jumlah_terapi"`
}

type PasienRequest struct {
	Nama   string `json:"nama"`
	Jk     int    `json:"jk"` //0 perempuan 1 laki-laki
	NoHp   string `json:"no_hp"`
	NoWa   string `json:"no_wa"`
	Alamat string `json:"alamat"`
}

func (b PasienRequest) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Nama, validation.Required),
		validation.Field(&b.Jk, validation.Max(1)),
		validation.Field(&b.NoHp, validation.Required),
		validation.Field(&b.NoWa, validation.Required),
		validation.Field(&b.Alamat, validation.Required),
	)
}
