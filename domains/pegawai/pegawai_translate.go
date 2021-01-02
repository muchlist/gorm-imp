package pegawai

import "github.com/mashingan/smapping"

func TranslateEntityToRes(data Pegawai) (*PegawaiResponse, error) {
	pegawaiResponse := PegawaiResponse{}
	err := smapping.FillStruct(&pegawaiResponse, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &pegawaiResponse, err
}
