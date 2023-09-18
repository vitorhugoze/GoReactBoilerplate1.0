package controller

import (
	"fmt"
	"net/http"
	"start/models"
	"start/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SignupController(ctx *gin.Context) {

	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		log.Err(fmt.Errorf("error binding to user variable %w", err))
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sucess, err := services.CreateUser(user)
	if err != nil {
		log.Err(fmt.Errorf("user created ?: %t", sucess))

		if !sucess {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

	}

	ctx.Status(http.StatusOK)
}
