package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sam-berry/ecfr-analyzer/server/httpresponse"
	"github.com/sam-berry/ecfr-analyzer/server/service"
)

type WordCountAPI struct {
	Router           fiber.Router
	WordCountService *service.WordCountService
}

func (api *WordCountAPI) Register() {
	api.Router.Get(
		"/word-count/:slug", func(c *fiber.Ctx) error {
			ctx := c.UserContext()
			slug := c.Params("slug")

			r, err := api.WordCountService.CountWordsForAgency(ctx, slug)

			if err != nil {
				return httpresponse.ApplyErrorToResponse(c, "Unexpected error", err)
			}

			return httpresponse.ApplySuccessToResponse(c, r)
		},
	)
}
