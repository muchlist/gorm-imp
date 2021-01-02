package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/muchlist/gorm-imp/controllers/pasien_controller"
	"github.com/muchlist/gorm-imp/controllers/pegawai_controller"
	"github.com/muchlist/gorm-imp/controllers/pengeluaran_controller"
	"github.com/muchlist/gorm-imp/controllers/terapi_controller"
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/utils/migrate"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Init Database
	database.InitDB()
	migrate.DbMigrate()

	app.Get("/api/pasien", pasien_controller.FindPasien)
	app.Post("/api/pasien", pasien_controller.CreatePasien)

	app.Get("/api/pegawai", pegawai_controller.FindPegawai)
	app.Post("/api/pegawai", pegawai_controller.CreatePegawai)

	app.Get("/api/terapi", terapi_controller.FindTerapi)
	app.Post("/api/terapi-range", terapi_controller.FindTerapiByRange)
	app.Post("/api/terapi", terapi_controller.CreateTerapi)

	app.Get("/api/pengeluaran", pengeluaran_controller.FindPengeluaran)
	app.Post("/api/pengeluaran", pengeluaran_controller.CreatePengeluaran)

	log.Fatal(app.Listen(":3000"))
}
