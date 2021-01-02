package terapi

import (
	"github.com/muchlist/gorm-imp/domains/pasien"
	"github.com/muchlist/gorm-imp/domains/pegawai"
	"time"
)

type Terapi struct {
	ID        uint
	Pasien    pasien.Pasien
	PasienID  int
	Tglterapi time.Time
	Layanan   string
	Biaya     float64
	Pegawai   pegawai.Pegawai `gorm:"foreignKey:PegawaiID"`
	PegawaiID int
}
