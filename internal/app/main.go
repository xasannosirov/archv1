package main

import (
	"archv1/internal/pkg/casbin"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/repo/cache"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/router"
	"archv1/internal/websocket"
	"fmt"
	"log"
)

func main() {
	cfg := config.NewConfig()

	hub := websocket.NewHub()

	redisClient, err := cache.NewRedis(cfg)
	fmt.Println(redisClient, err)
	if err != nil {
		log.Fatal(err)
	}

	go hub.Run()

	psql := postgres.NewDB(cfg)
	enforcer := casbin.NewEnforcer(cfg)

	engine := router.New(&router.Router{
		RedisCache: redisClient,
		Conf:       cfg,
		Hub:        hub,
		PostgresDB: psql,
		Enforcer:   enforcer,
	})

	if err := engine.Run(fmt.Sprintf("%s:%s", cfg.HttpHost, cfg.HttpPort)); err != nil {
		log.Fatal(err)
	}
}
