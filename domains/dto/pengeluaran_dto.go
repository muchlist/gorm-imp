package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type Pengeluaran struct {
	ID          uint      `json:"id"`
	PegawaiID   uint      `json:"pegawai_id"`
	Tanggal     time.Time `json:"tanggal"`
	Deskripsi   string    `json:"deskripsi"`
	BiayaSatuan float64   `json:"biaya_satuan"`
	Qty         float64   `json:"qty"`
	Lampiran    string    `json:"lampiran"`
}

type PengeluaranRequest struct {
	PegawaiID   uint    `json:"pegawai_id"`
	Deskripsi   string  `json:"deskripsi"`
	BiayaSatuan float64 `json:"biaya_satuan"`
	Qty         float64 `json:"qty"`
	Lampiran    string  `json:"lampiran"`
}

func (b PengeluaranRequest) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.PegawaiID, validation.Required),
		validation.Field(&b.Deskripsi, validation.Required),
		validation.Field(&b.BiayaSatuan, validation.Required),
		validation.Field(&b.Qty, validation.Required),
	)
}
