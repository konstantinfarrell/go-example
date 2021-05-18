package redis

import (
	rs "github.com/go-redis/redis"
	_ "github.com/go-redis/redis/v8"
)


func New(addr string, pass string, db int) (*rs.Client, error) {
	var options = &rs.Options{
		Addr: 		addr,
		Password:	pass,
		DB:			db, 
	}

	rdb := rs.NewClient(options)

	return rdb, nil
}

