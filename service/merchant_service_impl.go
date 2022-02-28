package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
	"weplant-backend/pkg"
	"weplant-backend/repository"
)

type MerchantServiceImpl struct {
	MerchantRepository   repository.MerchantRepository
	CloudinaryRepository repository.CloudinaryRepository
	ProductRepository    repository.ProductRepository
}

func NewMerchantService(merchantRepository repository.MerchantRepository, cloudinaryRepository repository.CloudinaryRepository, productRepository repository.ProductRepository) MerchantService {
	return &MerchantServiceImpl{
		MerchantRepository:   merchantRepository,
		CloudinaryRepository: cloudinaryRepository,
		ProductRepository:    productRepository,
	}
}

func (service *MerchantServiceImpl) Create(ctx context.Context, request web.MerchantCreateRequest) web.TokenResponse {
	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	res, err := service.MerchantRepository.Create(ctx, schema.Merchant{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Email:     request.Email,
		Password:  pkg.HashPassword(request.Password),
		Name:      request.Name,
		Slug:      request.Slug,
		Phone:     request.Phone,
		MainImage: &schema.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
		Address: &schema.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			PostalCode: request.Address.PostalCode,
		},
	})
	if err != nil {
		errUpload := service.CloudinaryRepository.DeleteImage(ctx, request.MainImage.FileName)
		helper.PanicIfError(errUpload)
		panic(err.Error())
	}

	token := pkg.GenerateToken(web.JWTPayload{
		Id:   res.Id.Hex(),
		Role: "merchant",
	})
	return web.TokenResponse{
		Id:    res.Id.Hex(),
		Role:  "merchant",
		Token: token,
	}
}

func (service *MerchantServiceImpl) FindById(ctx context.Context, merchantId string) web.MerchantDetailResponse {
	merchant, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfErrorNotFound(err)

	products, err := service.ProductRepository.FindByMerchantId(ctx, merchant.Id.Hex())
	helper.PanicIfError(err)

	var productsResponse []web.ProductSimpleResponse
	for _, p := range products {
		productsResponse = append(productsResponse, web.ProductSimpleResponse{
			Id:          p.Id.Hex(),
			MerchantId:  p.MerchantId,
			Name:        p.Name,
			Slug:        p.Slug,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			MainImage: web.ImageResponse{
				Id:       p.MainImage.Id.Hex(),
				FileName: p.MainImage.FileName,
				URL:      p.MainImage.URL,
			},
		})
	}

	return web.MerchantDetailResponse{
		Id:        merchant.Id.Hex(),
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		Email:     merchant.Email,
		Name:      merchant.Name,
		Slug:      merchant.Slug,
		Phone:     merchant.Phone,
		Balance:   merchant.Balance,
		MainImage: web.ImageResponse{
			Id:       merchant.MainImage.Id.Hex(),
			FileName: merchant.MainImage.FileName,
			URL:      merchant.MainImage.URL,
		},
		Address: web.AddressResponse{
			Address:    merchant.Address.Address,
			City:       merchant.Address.City,
			Province:   merchant.Address.Province,
			PostalCode: merchant.Address.PostalCode,
		},
		Products: productsResponse,
	}
}

func (service *MerchantServiceImpl) FindManageOrderById(ctx context.Context, merchantId string) web.ManageOrderResponse {
	merchant, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfErrorNotFound(err)

	var productsResponse []web.ManageOrderProductResponse
	for _, v := range merchant.Orders {
		product, err := service.ProductRepository.FindById(ctx, v.ProductId)
		helper.PanicIfError(err)

		productsResponse = append(productsResponse, web.ManageOrderProductResponse{
			Id:          v.Id.Hex(),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			ProductId:   product.Id.Hex(),
			Name:        product.Name,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       v.Price,
			Quantity:    v.Quantity,
			TotalPrice:  v.Price * v.Quantity,
			MainImage: web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      product.MainImage.URL,
			},
			Address: web.AddressResponse{
				Address:    v.Address.Address,
				City:       v.Address.City,
				Province:   v.Address.Province,
				PostalCode: v.Address.PostalCode,
			},
		})
	}

	return web.ManageOrderResponse{
		MerchantId: merchant.Id.Hex(),
		Products:   productsResponse,
	}
}

func (service *MerchantServiceImpl) Update(ctx context.Context, request web.MerchantUpdateRequest) web.MerchantUpdateRequest {
	merchant, err := service.MerchantRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	_, err = service.MerchantRepository.Update(ctx, schema.Merchant{
		Id:        merchant.Id,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      merchant.Slug,
		Phone:     request.Phone,
		Address: &schema.Address{
			Address:    request.Address.Address,
			City:       request.Address.City,
			Province:   request.Address.Province,
			PostalCode: request.Address.PostalCode,
		},
	})
	helper.PanicIfError(err)
	return request
}

func (service *MerchantServiceImpl) UpdateMainImage(ctx context.Context, request web.MerchantUpdateImageRequest) web.MerchantUpdateImageRequestResponse {
	merchant, err := service.MerchantRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.MerchantRepository.Update(ctx, schema.Merchant{
		Id:        merchant.Id,
		UpdatedAt: request.UpdatedAt,
		Slug:      merchant.Slug,
		MainImage: &schema.Image{
			Id:       merchant.MainImage.Id,
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	})
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, merchant.MainImage.FileName)
	helper.PanicIfError(err)

	return web.MerchantUpdateImageRequestResponse{
		Id:        merchant.Id.Hex(),
		UpdatedAt: request.UpdatedAt,
		MainImage: web.ImageResponse{
			Id:       merchant.Id.Hex(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	}
}

func (service *MerchantServiceImpl) Delete(ctx context.Context, merchantId string) {
	merchant, err := service.MerchantRepository.FindById(ctx, merchantId)
	helper.PanicIfErrorNotFound(err)

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
