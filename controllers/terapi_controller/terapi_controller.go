package terapi_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/terapi"
	"github.com/muchlist/gorm-imp/services/terapi_services"
)

func CreateTerapi(c *fiber.Ctx) error {
	var terapiFromBody terapi.TerapiRequest
	err := c.BodyParser(&terapiFromBody)
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	err = terapiFromBody.Validate()
	if err != nil {
		apiErr := rest_err.NewBadRequestError(err.Error())
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	terapiResp, err := terapi_services.TerapiService.Create(terapiFromBody)
	if err != nil {
		apiErr := rest_err.NewInternalServerError("Kesalahan dalam insert terapi", err)
		return c.Status(apiErr.Status()).JSON(apiErr)
	}

	return c.JSON(terapiResp)
}
