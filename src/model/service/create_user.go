package service

import (
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/mfcbentes/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
