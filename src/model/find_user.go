package model

import (
	"fmt"

	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *UserDomain) FindUser() (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))

	fmt.Println(u)

	return nil, nil
}
