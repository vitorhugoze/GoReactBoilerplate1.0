package middleware

import (
	"net/http"
	"os"
	"start/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AuthUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		jwtToken, err := ctx.Cookie("Token")

		if err != nil {
			log.Err(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {

			if claims, err := utils.CheckJwt(jwtToken); err != nil {
				log.Err(err)
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {

				signed, err := utils.GenerateJwt(claims.User)

				if err != nil {
					log.Err(err)
				} else {
					ctx.SetCookie("Token", signed, 3600*128, "/", os.Getenv("APP_HOST"), false, false)
				}

			}

		}

		ctx.Next()
	}

}
