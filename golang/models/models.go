package models

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/vmihailenco/msgpack/v4"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection
var Store *cache.Codec

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", envy.Get("REDIS_SERVER", "localhost"), envy.Get("REDIS_PORT", "6309")),
		Password: envy.Get("REDIS_PASSWORD", ""),
		PoolSize: runtime.NumCPU() * 5,
	})

	err = client.Ping().Err()

	if err != nil {
		log.Fatal(err)
	}

	Store = &cache.Codec{
		Redis: client,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
}
