package main

import (
	"archv1/internal/pkg/casbin"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/router"
	"fmt"
	"log"
)

func main() {
	cfg := config.NewConfig()

	psql := postgres.NewDB(cfg)
	cache := redis.NewRedis(cfg)
	enforcer := casbin.NewEnforcer(cfg)

	engine := router.New(&router.Router{
		Conf:       cfg,
		PostgresDB: psql,
		RedisCache: cache,
		Enforcer:   enforcer,
	})

	if err := engine.Run(fmt.Sprintf("%s:%s", cfg.HttpHost, cfg.HttpPort)); err != nil {
		log.Fatal(err)
	}

	log.Fatal("server exit")
}
