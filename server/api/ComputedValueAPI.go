package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sam-berry/ecfr-analyzer/server/httpresponse"
	"github.com/sam-berry/ecfr-analyzer/server/service"
)

type ComputedValueAPI struct {
	Router               fiber.Router
	ComputedValueService *service.ComputedValueService
}

func (api *ComputedValueAPI) Register() {
	api.Router.Post(
		"/compute/title-metrics", func(c *fiber.Ctx) error {
			ctx := c.UserContext()

			err := api.ComputedValueService.ProcessTitleMetrics(ctx)

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, nil)
		},
	)

	api.Router.Post(
		"/compute/agency-metrics", func(c *fiber.Ctx) error {
			ctx := c.UserContext()

			err := api.ComputedValueService.ProcessAgencyMetrics(ctx, false)

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, nil)
		},
	)

	api.Router.Post(
		"/compute/sub-agency-metrics", func(c *fiber.Ctx) error {
			ctx := c.UserContext()

			err := api.ComputedValueService.ProcessAgencyMetrics(ctx, true)

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, nil)
		},
	)
}
