package main

import (
	"net/http"
	"sync/atomic"

	"github.com/preemptcordon/server/cache"
	"github.com/preemptcordon/server/config"
	"github.com/preemptcordon/server/db"
	"github.com/preemptcordon/server/discover"
)

var visits uint64
var loadwatch chan uint64

const MAX_BACKGROUND = 1000

func loadmiddleware() {
	atomic.AddUint64(&visits, 1)
}

func emitresetload() {
	// once a second, emit the result of load and reset to zero
	loadwatch <- visits
	visits = 0
}

func loadlisten() {
	load := <-loadwatch
	if load < MAX_BACKGROUND {
		discover.DiscoverQueue()
	}
}

func main() {
	var err error
	config.ServerConfig, err = config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	cache.ConnectRedis(config.ServerConfig.RedisURI)
	loadwatch = make(chan uint64)
	go loadlisten()
	db.InitDB(config.ServerConfig)
	r := router()
	http.ListenAndServe(":8080", r)
}
