package logging

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gocraft/work"
	"github.com/labstack/echo/v4"
	"log"
	"log.autofarmer.go/connect/queue"
	"log.autofarmer.go/util"
)

type LoggingValidations struct{}

func (validation *LoggingValidations) IndexValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query IndexQueryFields
		c.Bind(&query)
		if err := query.IndexCheckValidate(); err != nil {
			return util.ResponseRouteValidation(c, err)
		}
		c.Set("query", query)
		return next(c)
	}
}

// Validate ...
func (query IndexQueryFields) IndexCheckValidate() error {
	return validation.ValidateStruct(
		&query,
		validation.Field(&query.Page, validation.Min(0).Error("số trang không hợp lệ")),
	)
}

func (validation *LoggingValidations) CreateValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queueData := work.Q{}
		if err := c.Bind(&queueData); err != nil {
			log.Println(err)
		}
		MyEnqueuer := work.NewEnqueuer(queue.AppName, queue.RedisPool)
		_, err := MyEnqueuer.Enqueue(
			"insertLog",
			queueData,
		)
		if err != nil {
			log.Fatal(err)
		}
		var payload struct{}
		c.Bind(&payload)
		c.Set("payload", payload)
		return next(c)
	}
}

func (validation *LoggingValidations) UpdateValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload struct{}
		c.Bind(&payload)
		c.Set("payload", payload)
		return next(c)
	}
}

func (validation *LoggingValidations) UpdateMaxLikedValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload struct {
			MaxLiked int `json:"maxLiked"`
		}
		c.Bind(&payload)
		c.Set("payload", payload)
		return next(c)
	}
}
