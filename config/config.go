package config

import (
	"github.com/redis/go-redis/v9"
)

type Configuration struct {
	//on redis server
	RedisAddress  string
	RedisPassword string
	RedisDb       int

	//DevelopMode should never be true in production
	DevelopMode   bool
	JwtSecret     string
	JwtFieldsKept string
	MaxBufferSize int64
	CORS          string

	//ServiceBatchSize is the number of tasks that a service can read from redis at the same time
	ServiceBatchSize int64
}

var Cfg Configuration = Configuration{
	JwtSecret:        "",
	MaxBufferSize:    32 << 20,
	CORS:             "*",
	ServiceBatchSize: 256,
}

// Parameter Server Should be Memory Only server, with high bandwidth and low latency.
// All parameter from web client are post to this redis server first
var ParamRds *redis.Client

const AppModeFRAMEWROK, AppModeSERVICE = "framework", "service"

// DataRds usually slower But with Flash Storage support ,such as Pikadb, and later may be KeyDB or DragonflyDB
// Default redis server to read data from and write data to web client
var DataRds *redis.Client
