package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sam-berry/ecfr-analyzer/server/httpresponse"
	"github.com/sam-berry/ecfr-analyzer/server/service"
)

type ECFRImport struct {
	Router            fiber.Router
	ECFRImportService service.ECFRImportService
}

func (api *ECFRImport) Register() {
	api.Router.Get(
		"/ecfr-import", func(c *fiber.Ctx) error {
			ctx := c.UserContext()

			data, err := api.ECFRImportService.GetData(ctx)

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, data)
		},
	)
}
