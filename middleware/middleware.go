package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// gzipStatic is a middleware that sets the Content-Encoding header to gzip if the requested file is a .gz file
func GzipStatic(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasSuffix(c.Request().URL.Path, ".gz") {
			c.Response().Header().Set("Content-Type", "application/javascript")
			c.Response().Header().Set("Content-Encoding", "gzip")
		}

		return next(c)
	}
}
