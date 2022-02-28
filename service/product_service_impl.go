package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"weplant-backend/helper"
	"weplant-backend/model/schema"
	"weplant-backend/model/web"
	"weplant-backend/repository"
)

type ProductServiceImpl struct {
	ProductRepository    repository.ProductRepository
	CloudinaryRepository repository.CloudinaryRepository
	CategoryRepository   repository.CategoryRepository
	MerchantRepository   repository.MerchantRepository
	CustomerRepository   repository.CustomerRepository
}

func NewProductService(productRepository repository.ProductRepository, cloudinaryRepository repository.CloudinaryRepository, categoryRepository repository.CategoryRepository, merchantRepository repository.MerchantRepository, customerRepository repository.CustomerRepository) ProductService {
	return &ProductServiceImpl{
		ProductRepository:    productRepository,
		CloudinaryRepository: cloudinaryRepository,
		CategoryRepository:   categoryRepository,
		MerchantRepository:   merchantRepository,
		CustomerRepository:   customerRepository,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductCreateRequestResponse {
	merchant, err := service.MerchantRepository.FindById(ctx, request.MerchantId)
	helper.PanicIfErrorNotFound(err)

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	var categoriesCreateRequest []schema.ProductCategory
	for _, category := range request.Categories {
		c, err := service.CategoryRepository.FindById(ctx, category.CategoryId)
		helper.PanicIfErrorNotFound(err)
		categoriesCreateRequest = append(categoriesCreateRequest, schema.ProductCategory{
			CategoryId: c.Id.Hex(),
		})
	}

	var imageCreateRequest []schema.Image
	for _, image := range request.Images {
		url, err := service.CloudinaryRepository.UploadImage(ctx, image.FileName, image.URL)
		helper.PanicIfError(err)
		imageCreateRequest = append(imageCreateRequest, schema.Image{
			Id:       primitive.NewObjectID(),
			FileName: image.FileName,
			URL:      url,
		})
	}

	res, err := service.ProductRepository.Create(ctx, schema.Product{
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		MerchantId:  merchant.Id.Hex(),
		Name:        request.Name,
		Slug:        request.Slug,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		MainImage: &schema.Image{
			Id:       primitive.NewObjectID(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
		Images:     imageCreateRequest,
		Categories: categoriesCreateRequest,
	})
	if err != nil {
		err := service.CloudinaryRepository.DeleteImage(ctx, request.MainImage.FileName)
		helper.PanicIfError(err)
		for _, image := range request.Images {
			err := service.CloudinaryRepository.DeleteImage(ctx, image.FileName)
			helper.PanicIfError(err)
		}
		panic(err.Error())
	}

	var imagesResponse []web.ImageResponse
	for _, image := range imageCreateRequest {
		imagesResponse = append(imagesResponse, web.ImageResponse{
			Id:       image.Id.Hex(),
			FileName: image.FileName,
			URL:      image.URL,
		})
	}

	var categoriesResponse []web.ProductCategoryCreateRequest
	for _, c := range res.Categories {
		category, err := service.CategoryRepository.FindById(ctx, c.CategoryId)
		helper.PanicIfError(err)
		categoriesResponse = append(categoriesResponse, web.ProductCategoryCreateRequest{
			CategoryId: category.Id.Hex(),
		})
	}

	return web.ProductCreateRequestResponse{
		Id:          res.Id.Hex(),
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
		MerchantId:  merchant.Id.Hex(),
		Name:        res.Name,
		Slug:        res.Slug,
		Description: res.Description,
		Price:       res.Price,
		Stock:       res.Stock,
		MainImage: web.ImageResponse{
			Id:       res.MainImage.Id.Hex(),
			FileName: res.MainImage.FileName,
			URL:      res.MainImage.URL,
		},
		Images:     imagesResponse,
		Categories: categoriesResponse,
	}
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) web.ProductDetailResponse {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfErrorNotFound(err)

	merchant, err := service.MerchantRepository.FindById(ctx, product.MerchantId)
	helper.PanicIfErrorNotFound(err)

	var imagesResponse []web.ImageResponse
	for _, img := range product.Images {
		imagesResponse = append(imagesResponse, web.ImageResponse{
			Id:       img.Id.Hex(),
			FileName: img.FileName,
			URL:      img.URL,
		})
	}

	var categoriesResponse []web.CategorySimpleResponse
	for _, v := range product.Categories {
		category, err := service.CategoryRepository.FindById(ctx, v.CategoryId)
		helper.PanicIfError(err)
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

	return web.ProductDetailResponse{
		Id:          product.Id.Hex(),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
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
		Images:     imagesResponse,
		Categories: categoriesResponse,
		Merchant: web.MerchantSimpleResponse{
			Id:    merchant.Id.Hex(),
			Name:  merchant.Name,
			Slug:  merchant.Slug,
			Phone: merchant.Phone,
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
		},
	}
}

func (service *ProductServiceImpl) FindAll(ctx context.Context, page int, perPage int) web.ProductFindAllResponse {
	skip := (page - 1) * perPage
	limit := perPage

	products, err := service.ProductRepository.FindAll(ctx, skip, limit)
	helper.PanicIfError(err)

	itemCount, err := service.ProductRepository.CountDocuments(ctx)
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

	return web.ProductFindAllResponse{
		Products: productsResponse,
		Metadata: web.MetadataPaginationResponse{
			CurrentPage: page,
			PerPage:     perPage,
			TotalData:   itemCount,
		},
	}
}
func (service *ProductServiceImpl) FindAllWithSearch(ctx context.Context, search string, page int, perPage int) web.ProductFindAllResponse {

	skip := (page - 1) * perPage
	limit := perPage

	products, err := service.ProductRepository.FindAllWithSearch(ctx, search, skip, limit)
	helper.PanicIfError(err)

	itemCount, err := service.ProductRepository.CountDocuments(ctx)
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

	return web.ProductFindAllResponse{
		Products: productsResponse,
		Metadata: web.MetadataPaginationResponse{
			CurrentPage: page,
			PerPage:     perPage,
			TotalData:   itemCount,
		},
	}
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductUpdateRequest {
	product, err := service.ProductRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	var categoriesUpdateRequest []schema.ProductCategory
	for _, v := range request.Categories {
		category, err := service.CategoryRepository.FindById(ctx, v.CategoryId)
		helper.PanicIfErrorNotFound(err)
		categoriesUpdateRequest = append(categoriesUpdateRequest, schema.ProductCategory{
			CategoryId: category.Id.Hex(),
		})
	}

	_, err = service.ProductRepository.Update(ctx, schema.Product{
		Id:          product.Id,
		UpdatedAt:   request.UpdatedAt,
		Name:        request.Name,
		Slug:        product.Slug,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Categories:  categoriesUpdateRequest,
	})
	helper.PanicIfError(err)
	return request
}

func (service *ProductServiceImpl) UpdateMainImage(ctx context.Context, request web.ProductUpdateImageRequest) web.ProductUpdateImageRequestResponse {
	product, err := service.ProductRepository.FindById(ctx, request.Id)
	helper.PanicIfErrorNotFound(err)

	url, err := service.CloudinaryRepository.UploadImage(ctx, request.MainImage.FileName, request.MainImage.URL)
	helper.PanicIfError(err)

	_, err = service.ProductRepository.Update(ctx, schema.Product{
		Id:        product.Id,
		UpdatedAt: request.UpdatedAt,
		Slug:      product.Slug,
		MainImage: &schema.Image{
			Id:       product.MainImage.Id,
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	})
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, product.MainImage.FileName)
	helper.PanicIfError(err)

	return web.ProductUpdateImageRequestResponse{
		Id:        product.Id.Hex(),
		UpdatedAt: request.UpdatedAt,
		MainImage: web.ImageResponse{
			Id:       product.MainImage.Id.Hex(),
			FileName: request.MainImage.FileName,
			URL:      url,
		},
	}
}

func (service *ProductServiceImpl) PushImageIntoImages(ctx context.Context, productId string, request []web.ImageCreateRequest) []web.ImageCreateRequest {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfErrorNotFound(err)

	var imagesCreateRequest []schema.Image
	var imagesResponse []web.ImageCreateRequest

	for _, image := range request {
		url, err := service.CloudinaryRepository.UploadImage(ctx, image.FileName, image.URL)
		helper.PanicIfError(err)
		imagesCreateRequest = append(imagesCreateRequest, schema.Image{
			Id:       primitive.NewObjectID(),
			FileName: image.FileName,
			URL:      url,
		})
		imagesResponse = append(imagesResponse, web.ImageCreateRequest{
			FileName: image.FileName,
			URL:      url,
		})
	}

	_, err = service.ProductRepository.PushImageIntoImages(ctx, product.Id.Hex(), imagesCreateRequest)
	helper.PanicIfError(err)

	return imagesResponse
}

func (service *ProductServiceImpl) PullImageFromImages(ctx context.Context, productId string, imageId string) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfErrorNotFound(err)

	res, err := service.ProductRepository.PullImageFromImages(ctx, product.Id.Hex(), imageId)
	helper.PanicIfErrorNotFound(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, res.FileName)
	helper.PanicIfError(err)

}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	helper.PanicIfErrorNotFound(err)

	err = service.CustomerRepository.PullProductFromAllCart(ctx, product.Id.Hex())
	helper.PanicIfError(err)

	err = service.ProductRepository.Delete(ctx, product.Id.Hex())
	helper.PanicIfError(err)

	err = service.CloudinaryRepository.DeleteImage(ctx, product.MainImage.FileName)
	helper.PanicIfError(err)

	for _, image := range product.Images {
		err = service.CloudinaryRepository.DeleteImage(ctx, image.FileName)
		helper.PanicIfError(err)
	}
}
