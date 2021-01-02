package terapi

import "github.com/muchlist/gorm-imp/database"

var (
	TerapiDao terapiDaoInterface = &terapiDao{}
)

type terapiDao struct{}

type terapiDaoInterface interface {
	Create(data Terapi) (Terapi, error)
}

func (p *terapiDao) Create(data Terapi) (Terapi, error) {
	db := database.DbConn
	var terapiData = data
	result := db.Create(&terapiData)

	return terapiData, result.Error
}
