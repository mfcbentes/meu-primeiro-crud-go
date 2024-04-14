package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String(
			"journey",
			"createUser",
		),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info.",
			err,
			zap.String(
				"journey",
				"createUser",
			),
		)

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully.",
		zap.String(
			"journey",
			"createUser",
		),
	)

	c.String(http.StatusOK, "")

}
