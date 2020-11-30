package server

import (
	"fmt"
	"log.autofarmer.go/connect/queue"

	"github.com/logrusorgru/aurora"

	"log.autofarmer.go/config"
	"log.autofarmer.go/connect/mongodb"
)

// Bootstrap modules
func bootstrap() {
	fmt.Println("")
	fmt.Println(aurora.Bold(aurora.Green("- SERVICES:")))
	config.Init()
	//redisdb.Init()
	mongodb.Init()
	queue.Init()
}
