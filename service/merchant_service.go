package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type MerchantService interface {
	Create(ctx context.Context, request web.MerchantCreateRequest, file interface{}) web.MerchantCreateRequest
	FindById(ctx context.Context, merchantId string) web.MerchantResponse
	UpdateById(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantUpdateRequest
	UpdateMainImageById(ctx context.Context, request web.MerchantUpdateImageRequest, file interface{}) web.MerchantUpdateImageRequest
	Delete(ctx context.Context, merchantId string)
}

type merchantServiceImpl struct {
	MerchantRepository   repository.MerchantRepository
	CloudinaryRepository repository.CloudinaryRepository
}

func NewMerchantService(merchantRepository repository.MerchantRepository, cloudinaryRepository repository.CloudinaryRepository) MerchantService {
	return &merchantServiceImpl{
		MerchantRepository:   merchantRepository,
		CloudinaryRepository: cloudinaryRepository,
	}
}

func (service *merchantServiceImpl) Create(ctx context.Context, request web.MerchantCreateRequest, file interface{}) web.MerchantCreateRequest {

	res, err := service.MerchantRepository.Create(ctx, domain.Merchant{
		CreatedAt: helper.GetTimeNow(),
		UpdatedAt: helper.GetTimeNow(),
		Email:     request.Email,
		Password:  helper.HashPassword(request.Password),
		Name:      request.Name,
		Phone:     request.Phone,
		MainImage: &domain.Image{
			Id:        primitive.NewObjectID(),
			CreatedAt: helper.GetTimeNow(),
			UpdatedAt: helper.GetTimeNow(),
			FileName:  request.MainImage.FileName,
		},
		Address: &domain.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			Country:    request.Address.Country,
			PostalCode: request.Address.PostalCode,
			Latitude:   request.Address.Latitude,
			Longitude:  request.Address.Longitude,
		},
	})
	helper.PanicIfError(err)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, file)
	helper.PanicIfError(errUpload)

	return web.MerchantCreateRequest{
		Email:    res.Email,
		Password: res.Password,
		Name:     res.Name,
		Phone:    res.Phone,
		MainImage: &web.ImageCreateRequest{

			FileName: res.MainImage.FileName,
		},
		Address: &web.AddressCreateRequest{
			Address:    res.Address.Address,
			City:       request.Address.City,
			Province:   res.Address.Province,
			Country:    res.Address.Country,
			PostalCode: res.Address.PostalCode,
			Latitude:   res.Address.Latitude,
			Longitude:  res.Address.Longitude,
		},
	}
}

func (service *merchantServiceImpl) FindById(ctx context.Context, merchantId string) web.MerchantResponse {
	res, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfError(err)

	imgUrl, errImg := service.CloudinaryRepository.GetImage(ctx, res.MainImage.FileName)
	helper.PanicIfError(errImg)

	return web.MerchantResponse{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Phone:     res.Phone,
		MainImage: &web.ImageResponse{
			Id:        res.MainImage.Id.Hex(),
			CreatedAt: res.MainImage.CreatedAt,
			UpdatedAt: res.MainImage.UpdatedAt,
			FileName:  res.MainImage.FileName,
			URL:       imgUrl,
		},
		Address: &web.AddressResponse{
			Address:    res.Address.Address,
			City:       res.Address.City,
			Province:   res.Address.Province,
			Country:    res.Address.Country,
			PostalCode: res.Address.PostalCode,
			Latitude:   res.Address.Latitude,
			Longitude:  res.Address.Longitude,
		},
	}
}

func (service *merchantServiceImpl) UpdateById(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantUpdateRequest {
	merchant, err := service.MerchantRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	res, errUpdate := service.MerchantRepository.UpdateById(ctx, domain.Merchant{
		Id:        merchant.Id,
		UpdatedAt: helper.GetTimeNow(),
		Name:      request.Name,
		Phone:     request.Phone,
		Address: &domain.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			Country:    request.Address.Country,
			PostalCode: request.Address.PostalCode,
			Latitude:   request.Address.Latitude,
			Longitude:  request.Address.Longitude,
		},
	})
	helper.PanicIfError(errUpdate)
	return web.MerchantUpdateRequest{
		Id:    res.Id.Hex(),
		Name:  res.Name,
		Phone: res.Phone,
		Address: &web.AddressUpdateRequest{
			Address:    res.Address.Address,
			City:       res.Address.City,
			Province:   res.Address.Province,
			Country:    res.Address.Country,
			PostalCode: res.Address.PostalCode,
			Latitude:   res.Address.Latitude,
			Longitude:  res.Address.Longitude,
		},
	}
}

func (service *merchantServiceImpl) UpdateMainImageById(ctx context.Context, request web.MerchantUpdateImageRequest, file interface{}) web.MerchantUpdateImageRequest {
	merchant, err := service.MerchantRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	res, errUpdate := service.MerchantRepository.UpdateById(ctx, domain.Merchant{
		Id:        merchant.Id,
		UpdatedAt: helper.GetTimeNow(),
		MainImage: &domain.Image{
			CreatedAt: helper.GetTimeNow(),
			UpdatedAt: helper.GetTimeNow(),
			FileName:  request.MainImage.FileName,
		},
	})
	helper.PanicIfError(errUpdate)

	errDelete := service.CloudinaryRepository.DeleteImage(ctx, merchant.MainImage.FileName)
	helper.PanicIfError(errDelete)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, file)
	helper.PanicIfError(errUpload)

	return web.MerchantUpdateImageRequest{
		Id: res.Id.Hex(),
		MainImage: &web.ImageUpdateRequest{
			FileName: res.MainImage.FileName,
		},
	}
}

func (service *merchantServiceImpl) Delete(ctx context.Context, merchantId string) {
	merchant, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfError(err)

	errDelete := service.MerchantRepository.DeleteById(ctx, merchant.Id.Hex())
	helper.PanicIfError(errDelete)

	errDeleteImg := service.CloudinaryRepository.DeleteImage(ctx, merchant.MainImage.FileName)
	helper.PanicIfError(errDeleteImg)
}
