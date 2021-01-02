package pengeluaran

import (
	"github.com/muchlist/gorm-imp/domains/pegawai"
	"time"
)

type Pengeluaran struct {
	ID          uint
	Tanggal     time.Time
	Pegawai     pegawai.Pegawai `gorm:"foreignKey:PegawaiID"`
	PegawaiID   int
	Deskripsi   string
	BiayaSatuan float64
	Qty         float64
	Lampiran    string
}
