package main

import (
	"fmt"
	"os"
	router "start/routes"
	"start/utils"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var _ = loadEnv()

func loadEnv() error {

	if os.Getenv("APP_ENV") != "PROD" {

		if err := godotenv.Load(".env"); err != nil {
			log.Fatal().Msg("Error getting .env config")
		}

	}

	return nil
}

func main() {

	//SetGinLog()
	ConfigLogger()
	db := utils.InitDb()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Err(fmt.Errorf("error closing postgres connection %w", err))
		}
	}()

	utils.InitRedisCli()

	r := router.GetRouter()

	r.Run(fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
}
