package config

import (
	"github.com/cloudinary/cloudinary-go"
	"os"
)

func GetCloud() *cloudinary.Cloudinary {

	cloudUrl := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudUrl)

	if err != nil {
		panic(err)
	}
	return cld
}
