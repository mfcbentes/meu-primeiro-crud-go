package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string
	Name     string
	Age      int8
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) SetID(id string) {
	u.ID = id
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
