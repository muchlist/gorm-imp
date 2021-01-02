package pasien_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/pasien"
	"github.com/muchlist/gorm-imp/services/pasien_services"
)

func FindPasien(c *fiber.Ctx) error {
	pasienData := pasien_services.PasienService.Find()
	return c.JSON(pasienData)
}

func CreatePasien(c *fiber.Ctx) error {

	var pasienFromBody pasien.PasienRequest
	err := c.BodyParser(&pasienFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	err = pasienFromBody.Validate()
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	pasienResp, err := pasien_services.PasienService.Create(pasienFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError("Input Salah")
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	return c.JSON(pasienResp)
}
