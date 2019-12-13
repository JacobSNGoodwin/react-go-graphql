package inmem

import (
	"github.com/go-redis/redis"
	"github.com/maxbrain0/react-go-graphql/server/logger"
)

// Conn holds globallay available redis-client connection
var Conn *redis.Client
var ctxLogger = logger.CtxLogger

// Redis holds data for initializing globally available redis conenction
type Redis struct {
	Addr     string
	Password string
}

// Connect used to create globally available redis connection
func (r *Redis) Connect() {
	// create REDIS client
	Conn = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       0,
	})

	_, err := Conn.Ping().Result()

	if err != nil {
		ctxLogger.Fatalf("Failed to create redis client connection: %v\n", err)
	}
}
