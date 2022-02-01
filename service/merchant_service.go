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
	Create(ctx context.Context, request web.MerchantCreateRequest) web.MerchantCreateRequest
	FindById(ctx context.Context, merchantId string) web.MerchantResponse
	FindManageOrderById(ctx context.Context, merchantId string) web.ManageOrderResponse
	Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantUpdateRequest
	UpdateMainImage(ctx context.Context, request web.MerchantUpdateImageRequest) web.MerchantUpdateImageRequest
	Delete(ctx context.Context, merchantId string)
}

type merchantServiceImpl struct {
	MerchantRepository   repository.MerchantRepository
	CloudinaryRepository repository.CloudinaryRepository
	ProductRepository    repository.ProductRepository
}

func NewMerchantService(merchantRepository repository.MerchantRepository, cloudinaryRepository repository.CloudinaryRepository, productRepository repository.ProductRepository) MerchantService {
	return &merchantServiceImpl{
		MerchantRepository:   merchantRepository,
		CloudinaryRepository: cloudinaryRepository,
		ProductRepository:    productRepository,
	}
}

func (service *merchantServiceImpl) Create(ctx context.Context, request web.MerchantCreateRequest) web.MerchantCreateRequest {

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.MerchantRepository.Create(ctx, domain.Merchant{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Email:     request.Email,
		Password:  helper.HashPassword(request.Password),
		Name:      request.Name,
		Slug:      request.Slug,
		Phone:     request.Phone,
		MainImage: &domain.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
		Address: &domain.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			Country:    request.Address.Country,
			PostalCode: request.Address.PostalCode,
		},
	})
	if err != nil {
		errUpload := service.CloudinaryRepository.DeleteImage(ctx, request.MainImage.FileName)
		helper.PanicIfError(errUpload)
		panic(err.Error())
	}

	request.MainImage.URL = url
	return request
}

func (service *merchantServiceImpl) FindById(ctx context.Context, merchantId string) web.MerchantResponse {

	res, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfError(err)

	return web.MerchantResponse{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Email:     res.Email,
		Name:      res.Name,
		Slug:      res.Slug,
		Phone:     res.Phone,
		Balance:   res.Balance,
		MainImage: &web.ImageResponse{
			Id:       res.MainImage.Id.Hex(),
			FileName: res.MainImage.FileName,
			URL:      res.MainImage.URL,
		},
		Address: &web.AddressResponse{
			Address:    res.Address.Address,
			City:       res.Address.City,
			Province:   res.Address.Province,
			Country:    res.Address.Country,
			PostalCode: res.Address.PostalCode,
		},
	}
}

func (service *merchantServiceImpl) FindManageOrderById(ctx context.Context, merchantId string) web.ManageOrderResponse {
	merchant, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfError(err)

	var productsResponse []*web.ManageOrderProductResponse
	for _, v := range merchant.Orders {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)

		productsResponse = append(productsResponse, &web.ManageOrderProductResponse{
			Id:          v.Id.Hex(),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			ProductId:   product.Id.Hex(),
			Name:        product.Name,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       v.Price,
			Quantity:    v.Quantity,
			SubTotal:    v.Price * v.Quantity,
			MainImage: &web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      product.MainImage.URL,
			},
			Address: &web.AddressResponse{
				Address:    v.Address.Address,
				City:       v.Address.City,
				Province:   v.Address.Province,
				Country:    v.Address.Country,
				PostalCode: v.Address.PostalCode,
			},
		})
	}

	return web.ManageOrderResponse{
		MerchantId: merchant.Id.Hex(),
		Products:   productsResponse,
	}
}

func (service *merchantServiceImpl) Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantUpdateRequest {
	merchant, err := service.MerchantRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	_, err = service.MerchantRepository.Update(ctx, domain.Merchant{
		Id:        merchant.Id,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      merchant.Slug,
		Phone:     request.Phone,
		Address: &domain.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			Country:    request.Address.Country,
			PostalCode: request.Address.PostalCode,
		},
	})
	helper.PanicIfError(err)
	return request
}

func (service *merchantServiceImpl) UpdateMainImage(ctx context.Context, request web.MerchantUpdateImageRequest) web.MerchantUpdateImageRequest {
	merchant, err := service.MerchantRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.MerchantRepository.Update(ctx, domain.Merchant{
		Id:        merchant.Id,
		UpdatedAt: request.UpdatedAt,
		Slug:      merchant.Slug,
		MainImage: &domain.Image{
			Id:       merchant.MainImage.Id,
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	})
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, merchant.MainImage.FileName)
	helper.PanicIfError(err)

	request.MainImage.URL = url
	return request
}

func (service *merchantServiceImpl) Delete(ctx context.Context, merchantId string) {
	merchant, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfError(err)

	products, err := service.ProductRepository.FindByMerchantId(ctx, merchant.Id.Hex())
	helper.PanicIfError(err)

	if products != nil {
		panic("tidak dapat menghapus toko ini karena didalamnya masih terdapat produk")
	}

	err = service.MerchantRepository.Delete(ctx, merchant.Id.Hex())
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, merchant.MainImage.FileName)
	helper.PanicIfError(err)
}
