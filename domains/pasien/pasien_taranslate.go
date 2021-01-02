package pasien

import "github.com/mashingan/smapping"

func TranslateReqToEntity(data PasienRequest) (*Pasien, error) {
	pasienEntity := Pasien{}
	err := smapping.FillStruct(&pasienEntity, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &pasienEntity, err
}
