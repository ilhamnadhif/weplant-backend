package config

import (
	"github.com/cloudinary/cloudinary-go"
	"os"
	"weplant-backend/helper"
)

func GetCloud() *cloudinary.Cloudinary {

	cloudUrl := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudUrl)

	helper.PanicIfError(err)
	return cld
}
