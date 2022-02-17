package pkg

import (
	"github.com/joho/godotenv"
	"weplant-backend/helper"
)

func GoDotENV() {
	err := godotenv.Load()
	helper.PanicIfError(err)
}
