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
	Create(ctx context.Context, request web.CategoryCreateRequest, file interface{}) web.CategoryCreateRequest
	FindById(ctx context.Context, categoryId string) web.CategoryResponseWithProduct
	FindAll(ctx context.Context) []web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest
	UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest, file interface{}) web.CategoryUpdateImageRequest
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

func (service *categoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest, file interface{}) web.CategoryCreateRequest {
	_, err := service.CategoryRepository.Create(ctx, domain.Category{
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
		MainImage: &domain.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
		},
	})
	helper.PanicIfError(err)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, file)
	helper.PanicIfError(errUpload)

	return request
}

func (service *categoryServiceImpl) FindById(ctx context.Context, categoryId string) web.CategoryResponseWithProduct {
	res, err := service.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfError(err)

	resProducts, errResProducts := service.ProductRepository.FindByCategoryId(ctx, res.Id.Hex())
	helper.PanicIfError(errResProducts)

	var products []*web.ProductResponseAll
	for _, product := range resProducts {
		imgUrl, errImg := service.CloudinaryRepository.GetImage(ctx, product.MainImage.FileName)
		helper.PanicIfError(errImg)
		products = append(products, &web.ProductResponseAll{
			Id:          product.Id.Hex(),
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			MerchantId:  product.MerchantId,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			MainImage: &web.ImageResponse{
				Id:       product.MainImage.Id.Hex(),
				FileName: product.MainImage.FileName,
				URL:      imgUrl,
			},
		})
	}

	imgUrl, errImg := service.CloudinaryRepository.GetImage(ctx, res.MainImage.FileName)
	helper.PanicIfError(errImg)

	return web.CategoryResponseWithProduct{
		Id:        res.Id.Hex(),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		MainImage: &web.ImageResponse{
			Id:       res.MainImage.Id.Hex(),
			FileName: res.MainImage.FileName,
			URL:      imgUrl,
		},
		Products: products,
	}
}

func (service *categoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	res, err := service.CategoryRepository.FindAll(ctx)
	helper.PanicIfError(err)

	var categories []web.CategoryResponse
	for _, category := range res {
		imgUrl, errImg := service.CloudinaryRepository.GetImage(ctx, category.MainImage.FileName)
		helper.PanicIfError(errImg)
		categories = append(categories, web.CategoryResponse{
			Id:        category.Id.Hex(),
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			Name:      category.Name,
			MainImage: &web.ImageResponse{
				Id:       category.MainImage.Id.Hex(),
				FileName: category.MainImage.FileName,
				URL:      imgUrl,
			},
		})
	}
	return categories
}

func (service *categoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryUpdateRequest {
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	_, errUpdate := service.CategoryRepository.Update(ctx, domain.Category{
		Id:        category.Id,
		UpdatedAt: request.UpdatedAt,
		Name:      request.Name,
	})
	helper.PanicIfError(errUpdate)
	return request
}

func (service *categoryServiceImpl) UpdateMainImage(ctx context.Context, request web.CategoryUpdateImageRequest, file interface{}) web.CategoryUpdateImageRequest {
	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	res, errUpdate := service.CategoryRepository.Update(ctx, domain.Category{
		Id:        category.Id,
		UpdatedAt: request.UpdatedAt,
		MainImage: &domain.Image{
			FileName: request.MainImage.FileName,
		},
	})
	helper.PanicIfError(errUpdate)

	errDelete := service.CloudinaryRepository.DeleteImage(ctx, category.MainImage.FileName)
	helper.PanicIfError(errDelete)

	errUpload := service.CloudinaryRepository.UploadImage(ctx, res.MainImage.FileName, file)
	helper.PanicIfError(errUpload)

	return request
}

func (service *categoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	helper.PanicIfError(err)

	errDeleteCategoryId := service.ProductRepository.PullCategoryIdFromProduct(ctx, category.Id.Hex())
	helper.PanicIfError(errDeleteCategoryId)

	errDelete := service.CategoryRepository.Delete(ctx, category.Id.Hex())
	helper.PanicIfError(errDelete)

	errDeleteImg := service.CloudinaryRepository.DeleteImage(ctx, category.MainImage.FileName)
	helper.PanicIfError(errDeleteImg)
}
