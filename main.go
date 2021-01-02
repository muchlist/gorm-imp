package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/muchlist/gorm-imp/controllers/book_controller"
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/domains/book"
	"log"
)

func DbMigrate(){
	// Migrate the schema
	err := database.DbConn.AutoMigrate(&book.Book{})
	if err != nil {
		log.Printf("Database tidak dimigrasi. err : %s", err.Error())
	}
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Init Database
	database.InitDB()
	DbMigrate()

	app.Get("/api/v1/book", book_controller.Find)
	app.Get("/api/v1/book/:id", book_controller.Get)
	app.Delete("/api/v1/book/:id", book_controller.Delete)
	app.Post("/api/v1/book", book_controller.Create)
	app.Put("/api/v1/book/:id", book_controller.Update)

	log.Fatal(app.Listen(":3000"))
}
