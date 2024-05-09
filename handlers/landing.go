package handlers

import (
	"github.com/bamsy/go-echo-templ/utils"
	"github.com/bamsy/go-echo-templ/views/pages"
	"github.com/labstack/echo/v4"
)

type LandingHandler struct{}

func (h *LandingHandler) LandingPage(ctx echo.Context) error {
	return utils.Render(ctx, pages.Landing())
}
