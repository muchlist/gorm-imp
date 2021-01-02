package pengeluaran

import (
	"github.com/mashingan/smapping"
	"github.com/muchlist/gorm-imp/domains/dto"
)

func TranslateReqToEntity(data dto.PengeluaranRequest) (*dto.Pengeluaran, error) {
	pengeluaranEntity := dto.Pengeluaran{}
	err := smapping.FillStruct(&pengeluaranEntity, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &pengeluaranEntity, err
}
