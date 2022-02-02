package repository

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"os"
)

type CloudinaryRepository interface {
	UploadImage(ctx context.Context, filename string, image interface{}) (string, error)
	//GetImage(ctx context.Context, filename string) (string, error)
	DeleteImage(ctx context.Context, filename string) error
}

type cloudinaryRepositoryImpl struct {
	Cloud *cloudinary.Cloudinary
}

func NewCloudinaryRepository(cloud *cloudinary.Cloudinary) CloudinaryRepository {
	return &cloudinaryRepositoryImpl{
		Cloud: cloud,
	}
}

func (repository *cloudinaryRepositoryImpl) UploadImage(ctx context.Context, filename string, file interface{}) (string, error) {
	cloudFolder := os.Getenv("CLOUDINARY_FOLDER")
	var url string
	res, err := repository.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: cloudFolder + "/" + filename,
	})
	url = res.URL
	if err != nil {
		return url, err
	}
	return url, nil
}
//func (repository *cloudinaryRepositoryImpl) GetImage(ctx context.Context, filename string) (string, error) {
//	cloudFolder := os.Getenv("CLOUDINARY_FOLDER")
//	var resURL string
//	res, err := repository.Cloud.Admin.Asset(ctx, admin.AssetParams{
//		PublicID: cloudFolder + "/" + filename,
//	})
//	if err != nil {
//		return resURL, err
//	}
//	resURL = res.SecureURL
//	return resURL, nil
//}

func (repository *cloudinaryRepositoryImpl) DeleteImage(ctx context.Context, filename string) error {
	cloudFolder := os.Getenv("CLOUDINARY_FOLDER")
	_, err := repository.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: cloudFolder + "/" + filename,
	})
	if err != nil {
		return err
	}
	return nil
}
