package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest, file interface{}) web.ProductCreateRequest
	FindById(ctx context.Context, productId string) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponseAll
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductUpdateRequest
	UpdateMainImage(ctx context.Context, request web.ProductUpdateImageRequest, file interface{}) web.ProductUpdateImageRequest
	PushImageIntoImages(ctx context.Context, productId string, request web.ImageCreateRequest, file interface{}) web.ImageCreateRequest
	PullImageFromImages(ctx context.Context, productId string, imageId string)
	Delete(ctx context.Context, productId string)
}

type productServiceImpl struct {
	ProductRepository    repository.ProductRepository
	CloudinaryRepository repository.CloudinaryRepository
	CategoryRepository   repository.CategoryRepository
	MerchantRepository   repository.MerchantRepository
	CustomerRepository   repository.CustomerRepository
}

func NewProductService(productRepository repository.ProductRepository, cloudinaryRepository repository.CloudinaryRepository, categoryRepository repository.CategoryRepository, merchantRepository repository.MerchantRepository, customerRepository repository.CustomerRepository) ProductService {
	return &productServiceImpl{
		ProductRepository:    productRepository,
		CloudinaryRepository: cloudinaryRepository,
		CategoryRepository:   categoryRepository,
		MerchantRepository:   merchantRepository,
		CustomerRepository:   customerRepository,
	}
}

func (service *productServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest, file interface{}) web.ProductCreateRequest {

	merchant, errMerchant := service.MerchantRepository.FindById(ctx, request.MerchantId)
	helper.PanicIfError(errMerchant)

	var categoriesCreateRequest []*domain.ProductCategory
	for _, category := range request.Categories {
		ctgry, errCtgry := service.CategoryRepository.FindById(ctx, category.CategoryId)
		helper.PanicIfError(errCtgry)
		categoriesCreateRequest = append(categoriesCreateRequest, &domain.ProductCategory{
			CategoryId: ctgry.Id.Hex(),
		})
	}

	res, err := service.ProductRepository.Create(ctx, domain.Product{
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		MerchantId:  merchant.Id.Hex(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		MainImage: &domain.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
		},
		Categories: categoriesCreateRequest,
	})
	helper.PanicIfError(err)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, res.MainImage.FileName, file)
	helper.PanicIfError(errUpload)

	return request
}

func (service *productServiceImpl) FindById(ctx context.Context, productId string) web.ProductResponse {
	res, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfError(err)

	imgUrl, errImg := service.CloudinaryRepository.GetImage(ctx, res.MainImage.FileName)
	helper.PanicIfError(errImg)

	var imagesResponse []*web.ImageResponse
	for _, img := range res.Images {
		url, errUrl := service.CloudinaryRepository.GetImage(ctx, img.FileName)
		helper.PanicIfError(errUrl)
		imagesResponse = append(imagesResponse, &web.ImageResponse{
			Id:       img.Id.Hex(),
			FileName: img.FileName,
			URL:      url,
		})
	}

	var categoriesResponse []*web.CategoryResponse
	for _, category := range res.Categories {
		ctgry, errCtgry := service.CategoryRepository.FindById(ctx, category.CategoryId)
		helper.PanicIfError(errCtgry)
		url, errUrl := service.CloudinaryRepository.GetImage(ctx, ctgry.MainImage.FileName)
		helper.PanicIfError(errUrl)
		categoriesResponse = append(categoriesResponse, &web.CategoryResponse{
			Id:        ctgry.Id.Hex(),
			CreatedAt: ctgry.CreatedAt,
			UpdatedAt: ctgry.UpdatedAt,
			Name:      ctgry.Name,
			MainImage: &web.ImageResponse{
				Id:       ctgry.MainImage.Id.Hex(),
				FileName: ctgry.MainImage.FileName,
				URL:      url,
			},
		})
	}
	return web.ProductResponse{
		Id:          res.Id.Hex(),
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
		MerchantId:  res.MerchantId,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Stock:       res.Stock,
		MainImage: &web.ImageResponse{
			Id:       res.MainImage.Id.Hex(),
			FileName: res.MainImage.FileName,
			URL:      imgUrl,
		},
		Images:     imagesResponse,
		Categories: categoriesResponse,
	}
}

func (service *productServiceImpl) FindAll(ctx context.Context) []web.ProductResponseAll {
	res, err := service.ProductRepository.FindAll(ctx)
	helper.PanicIfError(err)

	var productsResponse []web.ProductResponseAll
	for _, product := range res {
		imgUrl, errImg := service.CloudinaryRepository.GetImage(ctx, product.MainImage.FileName)
		helper.PanicIfError(errImg)
		productsResponse = append(productsResponse, web.ProductResponseAll{
			Id:          product.Id.Hex(),
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			MainImage: &web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      imgUrl,
			},
		})
	}

	return productsResponse
}

func (service *productServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductUpdateRequest {
	product, err := service.ProductRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	var categoriesUpdateRequest []*domain.ProductCategory
	for _, category := range request.Categories {
		ctgry, errCtgry := service.CategoryRepository.FindById(ctx, category.CategoryId)
		helper.PanicIfError(errCtgry)
		categoriesUpdateRequest = append(categoriesUpdateRequest, &domain.ProductCategory{
			CategoryId: ctgry.Id.Hex(),
		})
	}

	_, errUpdate := service.ProductRepository.Update(ctx, domain.Product{
		Id:          product.Id,
		UpdatedAt:   request.UpdatedAt,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Categories:  categoriesUpdateRequest,
	})
	helper.PanicIfError(errUpdate)
	return request
}

func (service *productServiceImpl) UpdateMainImage(ctx context.Context, request web.ProductUpdateImageRequest, file interface{}) web.ProductUpdateImageRequest {
	product, err := service.ProductRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	res, errUpdate := service.ProductRepository.Update(ctx, domain.Product{
		Id:        product.Id,
		UpdatedAt: request.UpdatedAt,
		MainImage: &domain.Image{
			Id:       product.MainImage.Id,
			FileName: request.MainImage.FileName,
		},
	})
	helper.PanicIfError(errUpdate)

	errDelete := service.CloudinaryRepository.DeleteImage(ctx, product.MainImage.FileName)
	helper.PanicIfError(errDelete)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, res.MainImage.FileName, file)
	helper.PanicIfError(errUpload)

	return request
}

func (service *productServiceImpl) PushImageIntoImages(ctx context.Context, productId string, request web.ImageCreateRequest, file interface{}) web.ImageCreateRequest {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfError(err)
	res, errPush := service.ProductRepository.PushImageIntoImages(ctx, product.Id.Hex(), domain.Image{
		Id:       primitive.NewObjectID(),
		FileName: request.FileName,
	})
	helper.PanicIfError(errPush)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, res.FileName, file)
	helper.PanicIfError(errUpload)
	return request
}

func (service *productServiceImpl) PullImageFromImages(ctx context.Context, productId string, imageId string) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfError(err)

	res, errPull := service.ProductRepository.PullImageFromImages(ctx, product.Id.Hex(), imageId)
	helper.PanicIfError(errPull)

	errDeleteImg := service.CloudinaryRepository.DeleteImage(ctx, res.FileName)
	helper.PanicIfError(errDeleteImg)

}

func (service *productServiceImpl) Delete(ctx context.Context, productId string) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfError(err)

	errDelete := service.ProductRepository.Delete(ctx, product.Id.Hex())
	helper.PanicIfError(errDelete)

	errPull := service.CustomerRepository.PullAllProductFromCart(ctx, product.Id.Hex())
	helper.PanicIfError(errPull)

	errDeleteImg := service.CloudinaryRepository.DeleteImage(ctx, product.MainImage.FileName)
	helper.PanicIfError(errDeleteImg)

	for _, image := range product.Images {
		errDeleteImage := service.CloudinaryRepository.DeleteImage(ctx, image.FileName)
		helper.PanicIfError(errDeleteImage)
	}
}
