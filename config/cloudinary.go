package config

import (
	"github.com/cloudinary/cloudinary-go"
	"os"
	"weplant-backend/helper"
)

func GetCloud() *cloudinary.Cloudinary {

	cloudName := os.Getenv("CLOUDINARY_NAME")
	cloudKey := os.Getenv("CLOUDINARY_KEY")
	cloudSecret := os.Getenv("CLOUDINARY_SECRET")

	cld, errorCloud := cloudinary.NewFromParams(cloudName, cloudKey, cloudSecret)
	helper.PanicIfError(errorCloud)
	return cld
}
