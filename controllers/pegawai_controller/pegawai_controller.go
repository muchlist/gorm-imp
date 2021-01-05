package pegawai_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/erru_utils_go/rest_err"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"github.com/muchlist/gorm-imp/services"
)

func FindPegawai(c *fiber.Ctx) error {
	pegawaiData, apiErr := services.PegawaiService.Find()
	if apiErr != nil {
		return c.Status(apiErr.Status()).JSON(apiErr)
	}
	return c.JSON(pegawaiData)
}

func CreatePegawai(c *fiber.Ctx) error {
	var pegawaiFromBody dto2.PegawaiRequest
	err := c.BodyParser(&pegawaiFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	err = pegawaiFromBody.Validate()
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	pegawaiResp, err := services.PegawaiService.Create(pegawaiFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError("Input Salah")
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	return c.JSON(pegawaiResp)
}
