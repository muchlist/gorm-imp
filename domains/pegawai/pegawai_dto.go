package pegawai

import (
	"github.com/muchlist/gorm-imp/domains/pengeluaran"
	"github.com/muchlist/gorm-imp/domains/terapi"
)

type Pegawai struct {
	ID           uint
	Nama         string
	Kontak       string
	Username     string
	Password     string
	Level        int                       // 0 admin , 1 superAdmin
	Terapis      []terapi.Terapi           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PegawaiID"`
	Pengeluarans []pengeluaran.Pengeluaran `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:PegawaiID"`
}
