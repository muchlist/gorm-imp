package pegawai

import (
	"github.com/mashingan/smapping"
	"github.com/muchlist/gorm-imp/domains/dto"
)

func TranslateEntityToRes(data dto.Pegawai) (*dto.PegawaiResponse, error) {
	pegawaiResponse := dto.PegawaiResponse{}
	err := smapping.FillStruct(&pegawaiResponse, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &pegawaiResponse, err
}
