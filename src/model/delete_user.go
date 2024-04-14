package model

import (
	"fmt"

	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *UserDomain) DeleteUser() *rest_err.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	fmt.Println(u)

	return nil
}
