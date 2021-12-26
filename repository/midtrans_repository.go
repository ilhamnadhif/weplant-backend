package repository

import (
	"errors"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"os"
	"weplant-backend/config"
)

type MidtransRepository interface {
	SendBalance(request snap.Request) (*snap.Response, *midtrans.Error)
}

type midtransRepositoryImpl struct {
}

func NewMidtransRepository() MidtransRepository {
	return &midtransRepositoryImpl{}
}

func (repository *midtransRepositoryImpl) SendBalance(request snap.Request) (*snap.Response, *midtrans.Error) {
	key := config.GetMidtransKey()

	var s snap.Client
	s.New(key, midtransEnvType(os.Getenv("MIDTRANS_ENV_TYPE")))

	res, err := s.CreateTransaction(&request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func midtransEnvType(env string) midtrans.EnvironmentType {
	if env == "production" {
		return midtrans.Production
	} else if env == "sandbox" {
		return midtrans.Sandbox
	} else {
		panic(errors.New("Error Env Type").Error())
	}
}
