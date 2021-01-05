package pengeluaran_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/erru_utils_go/rest_err"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"github.com/muchlist/gorm-imp/services/pengeluaran_services"
)

func FindPengeluaran(c *fiber.Ctx) error {
	pengeluaranData := pengeluaran_services.PengeluaranService.Find()
	return c.JSON(pengeluaranData)
}

func CreatePengeluaran(c *fiber.Ctx) error {
	var pengeluaranFromBody dto2.PengeluaranRequest
	err := c.BodyParser(&pengeluaranFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	err = pengeluaranFromBody.Validate()
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	pengeluaranResp, err := pengeluaran_services.PengeluaranService.Create(pengeluaranFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError("Input Salah")
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	return c.JSON(pengeluaranResp)
}
