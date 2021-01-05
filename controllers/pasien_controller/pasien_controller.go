package pasien_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/erru_utils_go/rest_err"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"github.com/muchlist/gorm-imp/services/pasien_services"
	"strings"
)

func FindPasien(c *fiber.Ctx) error {
	gender := c.Query("gender")
	// validation
	if gender != "" {
		// jika gender query berisi selain l atau p maka dianggap kosong
		if !(strings.ToLower(gender) == "l") || (strings.ToLower(gender) == "p") {
			gender = ""
		}
	}

	println(gender)

	pasienData := pasien_services.PasienService.Find(gender)
	return c.JSON(pasienData)
}

func CreatePasien(c *fiber.Ctx) error {

	var pasienFromBody dto2.PasienRequest
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
