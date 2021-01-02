package book_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/gorm-imp/domains/book"
	"github.com/muchlist/gorm-imp/services/book_services"
	"strconv"
)

func Find(c *fiber.Ctx) error {
	books, err := book_services.BookService.GetBooks()
	if err != nil {
		return c.JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}
	return c.JSON(books)
}

func Get(c *fiber.Ctx) error {
	id, err:= strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	bookUnit, err := book_services.BookService.GetBookByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}
	return c.JSON(bookUnit)
}

func Create(c *fiber.Ctx) error {
	var bookFromBody book.Book
	err := c.BodyParser(&bookFromBody)
	if err!= nil {
		return c.Status(500).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	err = bookFromBody.Validate()
	if err!= nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	bookResp, err := book_services.BookService.CreateBook(bookFromBody)
	if err!= nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	return c.JSON(bookResp)
}

func Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	err = book_services.BookService.DeleteBook(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}
	return c.JSON(fiber.Map{"msg": "Buku berhasil dihapus"})

}

func Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	var bookFromBody book.Book
	err = c.BodyParser(&bookFromBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	err = bookFromBody.Validate()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	bookFromBody.ID = uint(id)

	bookUpdated, err := book_services.BookService.UpdateBook(bookFromBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": fmt.Sprintf("Error : %v", err.Error())})
	}

	return c.JSON(bookUpdated)

}