package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/preemptcordon/server/cache"
	"github.com/preemptcordon/server/config"
	"github.com/preemptcordon/server/db"
	"github.com/preemptcordon/server/discover"
)

var visits uint64
var loadwatch chan uint64
var wg sync.WaitGroup

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

func apiserve() {
	r := router()
	log.Fatal(http.ListenAndServe(":8080", r))
	wg.Done()
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
	wg.Add(1)
	go apiserve()
	fmt.Println("started API server")
	wg.Wait()
	fmt.Println("exiting")
}
