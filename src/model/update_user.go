package model

import (
	"fmt"

	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *UserDomain) UpdateUser() *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))

	fmt.Println(u)

	return nil
}
