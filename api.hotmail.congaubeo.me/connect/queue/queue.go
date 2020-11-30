package queue

import (
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type QueueContext struct {
}

var (
	AppName   = "log.autofarmer.go"
	RedisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
	MyWorkerPool *work.WorkerPool
)

// Init ...
func Init() {
	MyWorkerPool = work.NewWorkerPool(QueueContext{}, 10, AppName, RedisPool)
}
