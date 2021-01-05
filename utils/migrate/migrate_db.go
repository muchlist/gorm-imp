package migrate

import (
	"github.com/muchlist/gorm-imp/database"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"log"
)

func DbMigrate() {
	// Migrate the schema
	err := database.DbConn.AutoMigrate(&dto2.Terapi{})
	if err != nil {
		log.Printf("database terapi gagal dimigrasi. err : %s", err.Error())
	}

	err = database.DbConn.AutoMigrate(&dto2.Pasien{})
	if err != nil {
		log.Printf("database pasien gagal dimigrasi. err : %s", err.Error())
	}

	err = database.DbConn.AutoMigrate(&dto2.Pengeluaran{})
	if err != nil {
		log.Printf("database pengeluaran gagal dimigrasi. err : %s", err.Error())
	}
}
