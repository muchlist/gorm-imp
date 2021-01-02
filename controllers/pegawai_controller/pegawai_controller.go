package pegawai_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/pegawai"
	"github.com/muchlist/gorm-imp/services/pegawai_services"
)

func FindPegawai(c *fiber.Ctx) error {
	pegawaiData := pegawai_services.PegawaiService.Find()
	return c.JSON(pegawaiData)
}

func CreatePegawai(c *fiber.Ctx) error {
	var pegawaiFromBody pegawai.PegawaiRequest
	err := c.BodyParser(&pegawaiFromBody)

	err = pegawaiFromBody.Validate()
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	pegawaiData := pegawai.Pegawai{
		Nama:     pegawaiFromBody.Nama,
		Kontak:   pegawaiFromBody.Kontak,
		Username: pegawaiFromBody.Username,
		Password: pegawaiFromBody.Password,
		Level:    pegawaiFromBody.Level,
	}

	pegawaiResp, err := pegawai_services.PegawaiService.Create(pegawaiData)
	if err != nil {
		apiErr := rest_err.NewBadRequestError("Input Salah")
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	return c.JSON(pegawaiResp)
}
