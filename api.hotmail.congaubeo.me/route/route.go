package route

import (
	"github.com/labstack/echo/v4"
	logging "log.autofarmer.go/modules/log"
	routemiddleware "log.autofarmer.go/route/middleware"
)

func Api(group *echo.Group) {
	group.Use(routemiddleware.CORSConfig())
	group.Use(routemiddleware.Auth)
	// Bootstrap routes

}

// Init ...
func Web(e *echo.Echo) {
	e.Use(routemiddleware.CORSConfig())
	logging.RouterWeb(e)
}
