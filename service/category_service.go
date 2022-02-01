package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/domain"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateRequest
	FindById(ctx context.Context, categoryId string) web.CategoryResponseWithProduct
	FindAll(ctx context.Context) []web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest
	UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest) web.CategoryUpdateImageRequest
	Delete(ctx context.Context, categoryId string)
}

type categoryServiceImpl struct {
	CategoryRepository   repository.CategoryRepository
	CloudinaryRepository repository.CloudinaryRepository
	ProductRepository    repository.ProductRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository, cloudinaryRepository repository.CloudinaryRepository, productRepository repository.ProductRepository) CategoryService {
	return &categoryServiceImpl{
		CategoryRepository:   categoryRepository,
		CloudinaryRepository: cloudinaryRepository,
		ProductRepository:    productRepository,
	}
}

func (service *categoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateRequest {
	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.CategoryRepository.Create(ctx, domain.Category{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      request.Slug,
		MainImage: &domain.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
			URL:      url,
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

func (service *categoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryResponseWithProduct {
	res, err := service.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfError(err)

	resProducts, err := service.ProductRepository.FindByCategoryId(ctx, res.Id.Hex())
	helper.PanicIfError(err)

	var products []*web.ProductResponseAll
	for _, product := range resProducts {
		products = append(products, &web.ProductResponseAll{
			Id:          product.Id.Hex(),
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			MerchantId:  product.MerchantId,
			Name:        product.Name,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			MainImage: &web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      product.MainImage.URL,
			},
		})
	}

	return web.CategoryResponseWithProduct{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Slug:      res.Slug,
		MainImage: &web.ImageResponse{
			Id:       res.MainImage.Id.Hex(),
			FileName: res.MainImage.FileName,
			URL:      res.MainImage.URL,
		},
		Products: products,
	}
}

func (service *categoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	res, err := service.CategoryRepository.FindAll(ctx)
	helper.PanicIfError(err)

	var categories []web.CategoryResponse
	for _, category := range res {
		categories = append(categories, web.CategoryResponse{
			Id:        category.Id.Hex(),
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			Name:      category.Name,
			Slug:      category.Slug,
			MainImage: &web.ImageResponse{
				Id:       category.MainImage.Id.Hex(),
				FileName: category.MainImage.FileName,
				URL:      category.MainImage.URL,
			},
		})
	}
	return categories
}

func (service *categoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest {
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	_, err = service.CategoryRepository.Update(ctx, domain.Category{
		Id:        category.Id,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      category.Slug,
	})
	helper.PanicIfError(err)
	return request
}

func (service *categoryServiceImpl) UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest) web.CategoryUpdateImageRequest {
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.CategoryRepository.Update(ctx, domain.Category{
		Id:        category.Id,
		UpdatedAt: request.UpdatedAt,
		Slug:      category.Slug,
		MainImage: &domain.Image{
			Id:       category.MainImage.Id,
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	})
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, category.MainImage.FileName)
	helper.PanicIfError(err)

	request.MainImage.URL = url
	return request
}

func (service *categoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfError(err)

	err = service.ProductRepository.PullCategoryIdFromProduct(ctx, category.Id.Hex())
	helper.PanicIfError(err)

	err = service.CategoryRepository.Delete(ctx, category.Id.Hex())
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, category.MainImage.FileName)
	helper.PanicIfError(err)
}
