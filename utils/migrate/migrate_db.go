package migrate

import (
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/domains/dto"
	"log"
)

func DbMigrate() {
	// Migrate the schema
	err := database.DbConn.AutoMigrate(&dto.Terapi{})
	if err != nil {
		log.Printf("database terapi gagal dimigrasi. err : %s", err.Error())
	}

	err = database.DbConn.AutoMigrate(&dto.Pasien{})
	if err != nil {
		log.Printf("database pasien gagal dimigrasi. err : %s", err.Error())
	}

	err = database.DbConn.AutoMigrate(&dto.Pengeluaran{})
	if err != nil {
		log.Printf("database pengeluaran gagal dimigrasi. err : %s", err.Error())
	}
}
