package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Você chamou a rota de forma incorreta")
	c.JSON(err.Code, err)
}
