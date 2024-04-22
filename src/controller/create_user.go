package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (u *userControllerInterface) CreateUser(c *gin.Context) {

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

	domainResult, err := u.service.CreateUser(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully.",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))

}
