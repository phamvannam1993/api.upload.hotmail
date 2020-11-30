package routemiddleware

import (
	"github.com/labstack/echo/v4"
	"log.autofarmer.go/config"
	"log.autofarmer.go/util"
)

// Auth api before move to controllers
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			appCfg     = config.GetEnv()
			authHeader = c.Request().Header.Get(echo.HeaderAuthorization)
		)

		// Return if header not valid
		if authHeader != appCfg.Auth.SecretKey {
			return util.Response401(c, nil, "")
		}

		return next(c)
	}
}
