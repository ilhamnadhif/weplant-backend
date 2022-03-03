package service

import (
	"context"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	ProductRepository  repository.ProductRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository, productRepository repository.ProductRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository:   categoryRepository,
		ProductRepository:    productRepository,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateRequestResponse {
	res, err := service.CategoryRepository.Create(ctx, schema.Category{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		Slug:      request.Slug,
	})
	helper.PanicIfError(err)

	return web.CategoryCreateRequestResponse{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Slug:      res.Slug,
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
		Products:  productsResponse,
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
		})
	}
	return categoriesResponse
}
