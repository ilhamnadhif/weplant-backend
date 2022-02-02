package helper

import (
	"errors"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CheckTransactionStatus(response coreapi.TransactionStatusResponse) string {
	if response.TransactionStatus == "capture" {
		if response.FraudStatus == "accept" {
			return "success"
		} else {
			return "failed"
		}
	} else if response.TransactionStatus == "settlement" {
		return "success"
	} else if response.TransactionStatus == "pending" {
		return "pending"
	} else {
		return "failed"
	}
}

func MidtransEnvType(env string) midtrans.EnvironmentType {
	if env == "production" {
		return midtrans.Production
	} else if env == "developer" {
		return midtrans.Sandbox
	} else {
		panic(errors.New("Error Env Type").Error())
	}
}
