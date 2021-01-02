package terapi

import (
	"time"
)

type Terapi struct {
	ID        uint
	PasienID  uint
	PegawaiID uint
	Tglterapi time.Time
	Layanan   string
	Biaya     float64
}
