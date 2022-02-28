package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository   repository.CategoryRepository
	CloudinaryRepository repository.CloudinaryRepository
	ProductRepository    repository.ProductRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository, cloudinaryRepository repository.CloudinaryRepository, productRepository repository.ProductRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository:   categoryRepository,
		CloudinaryRepository: cloudinaryRepository,
		ProductRepository:    productRepository,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateRequestResponse {
	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	res, err := service.CategoryRepository.Create(ctx, schema.Category{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      request.Slug,
		MainImage: &schema.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	})
	if err != nil {
		err := service.CloudinaryRepository.DeleteImage(ctx, request.MainImage.FileName)
		helper.PanicIfError(err)
		panic(err.Error())
	}

	return web.CategoryCreateRequestResponse{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Slug:      res.Slug,
		MainImage: web.ImageResponse{
			Id:       res.MainImage.Id.Hex(),
			FileName: res.MainImage.FileName,
			URL:      res.MainImage.URL,
		},
	}
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryDetailResponse {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfErrorNotFound(err)

	products, err := service.ProductRepository.FindByCategoryId(ctx, category.Id.Hex())
	helper.PanicIfError(err)

	var productsResponse []web.ProductSimpleResponse
	for _, product := range products {
		productsResponse = append(productsResponse, web.ProductSimpleResponse{
			Id:          product.Id.Hex(),
			MerchantId:  product.MerchantId,
			Name:        product.Name,
			Slug:        product.Slug,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			MainImage: web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      product.MainImage.URL,
			},
		})
	}

	return web.CategoryDetailResponse{
		Id:        category.Id.Hex(),
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
		Slug:      category.Slug,
		MainImage: web.ImageResponse{
			Id:       category.MainImage.Id.Hex(),
			FileName: category.MainImage.FileName,
			URL:      category.MainImage.URL,
		},
		Products: productsResponse,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategorySimpleResponse {
	categories, err := service.CategoryRepository.FindAll(ctx)
	helper.PanicIfError(err)

	var categoriesResponse []web.CategorySimpleResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, web.CategorySimpleResponse{
			Id:   category.Id.Hex(),
			Name: category.Name,
			Slug: category.Slug,
			MainImage: web.ImageResponse{
				Id:       category.MainImage.Id.Hex(),
				FileName: category.MainImage.FileName,
				URL:      category.MainImage.URL,
			},
		})
	}
	return categoriesResponse
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest {
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	_, err = service.CategoryRepository.Update(ctx, schema.Category{
		Id:        category.Id,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      category.Slug,
	})
	helper.PanicIfError(err)
	return request
}

func (service *CategoryServiceImpl) UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest) web.CategoryUpdateImageRequestResponse {
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.CategoryRepository.Update(ctx, schema.Category{
		Id:        category.Id,
		UpdatedAt: request.UpdatedAt,
		Slug:      category.Slug,
		MainImage: &schema.Image{
			Id:       category.MainImage.Id,
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	})
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, category.MainImage.FileName)
	helper.PanicIfError(err)

	return web.CategoryUpdateImageRequestResponse{
		Id:        category.Id.Hex(),
		UpdatedAt: request.UpdatedAt,
		MainImage: web.ImageResponse{
			Id:       category.MainImage.Id.Hex(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	}
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfError(err)

	err = service.ProductRepository.PullCategoryIdFromProduct(ctx, category.Id.Hex())
	helper.PanicIfError(err)

	err = service.CategoryRepository.Delete(ctx, category.Id.Hex())
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, category.MainImage.FileName)
	helper.PanicIfError(err)
}
