package utils

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var redisCli *redis.Client

func InitRedisCli() {
	redisAddr := fmt.Sprintf("%v:%v", os.Getenv("RD_HOST"), os.Getenv("RD_PORT"))
	db, _ := strconv.Atoi(os.Getenv("RD_DATA"))

	cli := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("RD_PASS"),
		DB:       db,
	})

	ctx := context.Background()

	if _, err := cli.Ping(ctx).Result(); err != nil {
		log.Fatal().Err(fmt.Errorf("erro ao criar cliente redis %w", err))
	}

	redisCli = cli
}

func RedisCli() (*redis.Client, error) {
	ctx := context.Background()

	if _, err := redisCli.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("cliente redis não está respondendo %w", err)
	}

	return redisCli, nil
}
