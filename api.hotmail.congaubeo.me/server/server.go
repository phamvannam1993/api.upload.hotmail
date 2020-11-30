package server

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/gocraft/work"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log.autofarmer.go/config"
	"log.autofarmer.go/connect/queue"
	CoreModels "log.autofarmer.go/modules/core/models"
	logging "log.autofarmer.go/modules/log"
	"log.autofarmer.go/route"
	"log.autofarmer.go/util"
)

// Start ...
func Start() {
	// Bootstrap
	bootstrap()
	pool := queue.MyWorkerPool
	pool.Job("insertLog", func(job *work.Job) error {
		log.Println("InsertLog")
		loggingController := logging.LoggingController{}
		input := job.Args
		for _, record := range input {
			log.Printf(" [===>] Record: %s", record)
			log.Println(reflect.TypeOf(record))
			for _, v := range record.([]interface{}) { // use type assertion to loop over []interface{}
				hotmail := v.(map[string]interface{})
				hotmail["status"] = "live"
				hotmail["created_at"] = time.Now().Unix() * 1000
				loggingController.InsertLog(hotmail)
			}
		}

		log.Println(job.Args)
		if err := job.ArgError(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	pool.Middleware(func(job *work.Job, next work.NextMiddlewareFunc) error {
		fmt.Println("Starting job: ", job.Name)
		return next()
	})
	pool.Start()
	/*signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)*/

	// Get config
	appCfg := config.GetEnv()

	// Echo instance
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://log.congaubeo.me", "http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Middleware
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	apiGroup := e.Group("/api")
	// Init routes
	route.Api(apiGroup)
	route.Web(e)
	// Start server
	util.ConsolePrintServiceSuccess("Server environment", appCfg.Env)
	util.ConsolePrintServerRoutes(e.Routes())
	//e.Logger.Fatal(e.Start(appCfg.Port.App))
	e.Logger.Fatal(e.Start(":8888"))
}

type QueueContext struct{}

func (c *QueueContext) InsertLog(job *work.Job) error {
	fmt.Println("InsertLog")
	fmt.Println(job.Args)
	var (
		loggingModels = logging.LoggingModels{
			CoreModels.MongoDBModels{
				Collection: "hotmail",
			},
		}
	)
	_, err := loggingModels.Insert(context.Background(), job.Args)
	if err != nil {
	}

	addr := job.ArgString("address")
	subject := job.ArgString("subject")
	fmt.Println("ServiceCode: " + addr + " subject: " + subject)
	if err := job.ArgError(); err != nil {
		return err
	}
	return nil
}
