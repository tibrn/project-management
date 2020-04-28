package models

import (
	"time"

	"github.com/go-redis/cache/v7"
)

func Cache(Key string, t time.Duration, item interface{}, f func() (interface{}, error)) (rez interface{}, err error) {
	//Try get user from Redis Cache Store

	if errr := Store.Get(Key, item); errr == nil {

		rez = item
		return
	}

	//Retrive item from other place than redis
	item, err = f()

	if err != nil {

		return
	}

	rez = item

	//Send user to redis cache store
	go func() {
		Store.Set(&cache.Item{
			Key:        Key,
			Object:     item,
			Expiration: t,
		})
	}()

	return
}
