package pasien

import (
	"github.com/mashingan/smapping"
	"github.com/muchlist/gorm-imp/domains/dto"
)

func TranslateReqToEntity(data dto.PasienRequest) (*dto.Pasien, error) {
	pasienEntity := dto.Pasien{}
	err := smapping.FillStruct(&pasienEntity, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &pasienEntity, err
}
