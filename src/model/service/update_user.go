package service

import (
	"fmt"

	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	fmt.Println(u)

	return nil
}
