package controller

import (
	"errors"
	"net/http"
	"os"
	"start/models"
	"start/services"
	"start/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoginController(ctx *gin.Context) {

	var login models.UserLogin

	if err := ctx.BindJSON(&login); err != nil {
		log.Err(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := services.GetUser(login)
	if err != nil {
		log.Err(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	valid := utils.CompareHash(login.Password, user.Password)
	if !valid {
		log.Err(errors.New("password dont match with value stored on database"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	signed, err := utils.GenerateJwt(*user)

	if err != nil {
		log.Err(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.SetCookie("Token", signed, 3600*128, "/", os.Getenv("APP_HOST"), false, false)

	ctx.Status(http.StatusOK)
}
