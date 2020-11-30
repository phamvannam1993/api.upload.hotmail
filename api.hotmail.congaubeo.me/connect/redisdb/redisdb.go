package redisdb

import (
	"context"
	"encoding/json"
	"github.com/kr/pretty"

	"github.com/go-redis/redis/v8"

	"log.autofarmer.go/config"
	"log.autofarmer.go/util"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// Init ...
func Init() {
	var (
		envVars  = config.GetEnv()
		redisURI = envVars.Redis.URI
		redisPwd = envVars.Redis.Password
	)

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPwd,
		DB:       0, // use default DB
	})

	// Test
	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		//log.Fatal("Cannot connect to redis", redisURI, err)
	}

	if !config.IsTest() {
		util.ConsolePrintServiceSuccess("Redis", redisURI+" - Password: "+redisPwd)
	}
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}) {
	storeByte, _ := json.Marshal(value)
	rdb.Set(ctx, key, storeByte, 0)
}

// GetValueByKey ...
func GetValueByKey(key string) interface{} {
	value, _ := rdb.Get(ctx, key).Result()
	return value
}

// DelKey ...
func DelKey(key string) {
	rdb.Del(ctx, key)
}

// DelPattern ...
func DelPattern(pattern string) {
	iter := rdb.Scan(ctx, 0, pattern, 10000).Iterator()
	for iter.Next(ctx) {
		rdb.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		pretty.Println("- Error when delete redis key", pattern, err)
	}
}

// CountKeys ...
func CountKeys(key string) (total int) {
	iter := rdb.Scan(ctx, 0, key, 10000).Iterator()
	for iter.Next(ctx) {
		total++
	}
	if err := iter.Err(); err != nil {
		pretty.Println("- Error when count redis key", key, err)
		return 0
	}
	return total
}

// AddMultipleKeys (map[string]interface{}{"key1": "value1", "key2": "value2"})
func AddMultipleKeys(values map[string]interface{}) {
	rdb.MSet(ctx, values)
}

// FlushAllAsync ...
func FlushAllAsync() {
	rdb.FlushAllAsync(ctx)
}

// CheckKeyExisted ...
func CheckKeyExisted(key string) bool {
	existed := rdb.Exists(ctx, key).Val()
	return existed == 1
}
