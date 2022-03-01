package middleware

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	address  = "172.18.0.91:6379"
	password = "password123"
	dbname   = 0
)

//SessionDB is a wrapper for redis like DB
type SessionDB struct {
	ctx    context.Context
	client *redis.Client
	err    error
}

//New creates a singleton instance
func (sdb *SessionDB) New() {
	sdb.ctx = context.TODO()
	sdb.client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       dbname,
	})
}

//Set puts the key, value
func (sdb *SessionDB) Set(key, val string) {
	sdb.err = sdb.client.Set(sdb.ctx, key, val, 0).Err()
	if sdb.err != nil {
		log.Fatal(sdb.err)
	}
}

//Get fetches the key
func (sdb *SessionDB) Get(key string) (val string) {
	val, sdb.err = sdb.client.Get(sdb.ctx, key).Result()
	if sdb.err != nil {
		log.Fatal(sdb.err)
	}
	return val
}
