package redis

import (
	"context"
	"time"

	_ "github.com/go-redis/redis"
	rs "github.com/go-redis/redis/v8"
)

type Cacher interface {
	Set(context.Context, string, interface{}, time.Duration) *rs.StatusCmd
	Get(context.Context, string) *rs.StringCmd
	Del(context.Context, ...string) *rs.IntCmd
}

func New(addr string, pass string, db int) (Cacher, error) {
	var options = &rs.Options{
		Addr: 		addr,
		Password:	pass,
		DB:			db, 
	}

	rdb := rs.NewClient(options)

	return rdb, nil
}

