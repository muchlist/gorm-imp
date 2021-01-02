package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/muchlist/gorm-imp/controllers/pasien_controller"
	"github.com/muchlist/gorm-imp/controllers/pegawai_controller"
	"github.com/muchlist/gorm-imp/controllers/terapi_controller"
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

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Init Database
	database.InitDB()
	DbMigrate()

	app.Get("/api/pasien", pasien_controller.FindPasien)
	app.Post("/api/pasien", pasien_controller.CreatePasien)

	app.Get("/api/pegawai", pegawai_controller.FindPegawai)
	app.Post("/api/pegawai", pegawai_controller.CreatePegawai)

	app.Get("/api/terapi", terapi_controller.FindTerapi)
	app.Post("/api/terapi-range", terapi_controller.FindTerapiByRange)
	app.Post("/api/terapi", terapi_controller.CreateTerapi)

	log.Fatal(app.Listen(":3000"))
}
