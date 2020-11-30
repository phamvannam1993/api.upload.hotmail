package logging

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	loggingController  = LoggingController{}
	loggingValidations = LoggingValidations{}
)

func RouterApi(e *echo.Group) {
	e.GET("", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	}))

}

func RouterWeb(e *echo.Echo) {
	routes := e.Group("/logging")
	routes.GET("", loggingController.Index, loggingValidations.IndexValidation)                                       //GetALL
	routes.POST("", loggingController.Store, loggingValidations.CreateValidation)                                     //Insert
	routes.GET("/:id", loggingController.Show, loggingController.CheckExisted)                                        //GetOne
	routes.PUT("/:id", loggingController.Update, loggingController.CheckExisted, loggingValidations.UpdateValidation) //Update
	routes.DELETE("/:id", loggingController.Delete, loggingController.CheckExisted)                                   //DELETE
}
