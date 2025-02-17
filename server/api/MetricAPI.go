package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sam-berry/ecfr-analyzer/server/httpresponse"
	"github.com/sam-berry/ecfr-analyzer/server/service"
)

type MetricAPI struct {
	Router              fiber.Router
	AgencyMetricService *service.AgencyMetricService
	TitleMetricService  *service.TitleMetricService
}

func (api *MetricAPI) Register() {
	api.Router.Get(
		"/agency-metrics/:slug", func(c *fiber.Ctx) error {
			ctx := c.UserContext()
			slug := c.Params("slug")

			r, err := api.AgencyMetricService.CountWordsAndSections(ctx, slug, "")

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, r)
		},
	)

	api.Router.Get(
		"/title-metrics", func(c *fiber.Ctx) error {
			ctx := c.UserContext()

			r, err := api.TitleMetricService.CountAllWordsAndSections(ctx)

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, r)
		},
	)
}
