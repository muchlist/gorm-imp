package dto

import (
	"time"
)

type Pengeluaran struct {
	ID          uint
	PegawaiID   uint
	Tanggal     time.Time
	Deskripsi   string
	BiayaSatuan float64
	Qty         float64
	Lampiran    string
}
