package CoreControllers

import (
	"github.com/labstack/echo/v4"
	"log.autofarmer.go/util"
)

type BaseControllers struct{}

func (controller *BaseControllers) Index(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "index",
		"query":      c.QueryParams(),
	}, "")
}
func (controller *BaseControllers) Show(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "show",
		"query":      c.QueryParams(),
	}, "")
}
func (controller *BaseControllers) Create(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "create",
		"query":      c.QueryParams(),
	}, "")
}
func (controller *BaseControllers) Store(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "store",
		"payload":    c.Get("payload"),
	}, "")
}
func (controller *BaseControllers) Edit(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "edit",
		"query":      c.QueryParams(),
	}, "")
}
func (controller *BaseControllers) Update(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "update",
		"query":      c.QueryParams(),
		"payload":    c.Get("payload"),
	}, "")
}
func (controller *BaseControllers) Delete(c echo.Context) error {
	return util.Response200(c, echo.Map{
		"controller": "delete",
		"query":      c.QueryParams(),
	}, "")
}
