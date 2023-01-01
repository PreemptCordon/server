package main

import "context"

func main() {
	ServerConfig := LoadConfig(".")
	ctx := context.Background()
	cache.connectRedis(ctx)
}
