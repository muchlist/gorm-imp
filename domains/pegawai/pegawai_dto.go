package pegawai

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/muchlist/gorm-imp/domains/pengeluaran"
	"github.com/muchlist/gorm-imp/domains/terapi"
)

type Pegawai struct {
	ID           uint
	Nama         string
	Kontak       string
	Username     string
	Password     string
	Level        int                       // 0 user , 1 admin, 2 superadmin
	Terapis      []terapi.Terapi           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PegawaiID"`
	Pengeluarans []pengeluaran.Pengeluaran `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PegawaiID"`
}

type PegawaiRequest struct {
	Nama     string
	Kontak   string
	Username string
	Password string
	Level    int
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
