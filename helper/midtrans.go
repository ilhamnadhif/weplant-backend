package helper

import (
	"errors"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CheckTransactionStatus(response coreapi.TransactionStatusResponse) bool {
	if response.TransactionStatus == "capture" {
		if response.FraudStatus == "accept" {
			return true
		} else {
			return false
		}
	} else if response.TransactionStatus == "settlement" {
		return true
	} else {
		return false
	}
}

func MidtransEnvType(env string) midtrans.EnvironmentType {
	if env == "production" {
		return midtrans.Production
	} else if env == "sandbox" {
		return midtrans.Sandbox
	} else {
		panic(errors.New("Error Env Type").Error())
	}
}
