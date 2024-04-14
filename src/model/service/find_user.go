package service

import (
	"fmt"

	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	fmt.Println(u)

	return nil, nil
}
