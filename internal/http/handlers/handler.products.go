package handlers

import (
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/aws"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productUsecase domains.ProductsUsecase
	s3Client       *aws.S3Client
}

func NewProductHandler(productUsecase domains.ProductsUsecase, s3Client *aws.S3Client) ProductHandler {
	return ProductHandler{
		productUsecase: productUsecase,
		s3Client:       s3Client,
	}
}

func (p *ProductHandler) Save(ctx echo.Context) error {
	var productCreateRequest requests.ProductCreateRequest

	if err := helpers.BindAndValidate(ctx, &productCreateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Image is required")
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	images := form.File["images"]

	mainImageKey := image.Filename

	mainImageUrl, err := p.s3Client.UploadFile(mainImageKey, image)

	var imageUrls []string

	for _, image := range images {
		key := image.Filename
		path, err := p.s3Client.UploadFile(key, image)
		if err != nil {
			continue
		}
		imageUrls = append(imageUrls, path)
	}

	productDomain := &domains.ProductDomain{
		Name:          productCreateRequest.Name,
		Description:   productCreateRequest.Description,
		Price:         productCreateRequest.Price,
		SubcategoryId: productCreateRequest.SubCategoryId,
		Image:         mainImageUrl,
		Images:        imageUrls,
	}

	statusCode, err := p.productUsecase.Save(productDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, 200, "success", productDomain)
}

func (p *ProductHandler) FindAllForMe(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	filter := domains.ProductFilter{
		Name:          ctx.QueryParam("name"),
		MinPrice:      ctx.QueryParam("min_price"),
		MaxPrice:      ctx.QueryParam("max_price"),
		SubcategoryID: ctx.QueryParam("subcategory_id"),
		BrandID:       ctx.QueryParam("brand_id"),
		Page:          helpers.ToInt(ctx.QueryParam("page"), 1),
		PageSize:      helpers.ToInt(ctx.QueryParam("page_size"), 10),
	}

	fmt.Printf("Working")

	log.Info("filter", filter)

	products, statusCode, err := p.productUsecase.FindAllForMe(jwtClaims.UserId, filter)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Products fetched successfully", responses.ToArrayOfProductResponse(products))
}

func (p *ProductHandler) UpdateById(ctx echo.Context) error {
	var productUpdateRequest requests.ProductUpdateRequest
	productId := ctx.Param("id")

	if err := helpers.BindAndValidate(ctx, &productUpdateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	image, _ := ctx.FormFile("image")

	form, _ := ctx.MultipartForm()
	images := form.File["images"]

	mainImageKey := image.Filename
	var mainImageUrl string
	if image != nil {
		imageUrl, err := p.s3Client.UploadFile(mainImageKey, image)
		mainImageUrl = imageUrl

		if err != nil {
			return NewErrorResponse(ctx, http.StatusBadRequest, "Failed to upload image")
		}
	}

	var imageUrls []string
	if images != nil {
		for _, image := range images {
			key := image.Filename
			path, err := p.s3Client.UploadFile(key, image)
			if err != nil {
				continue
			}
			imageUrls = append(imageUrls, path)
		}
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid product id")
	}
	product, statusCode, err := p.productUsecase.FindById(productIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if productUpdateRequest.Weight != nil {
		product.Weight = productUpdateRequest.Weight
	}
	if productUpdateRequest.SubCategory != nil {
		product.SubcategoryId = *productUpdateRequest.SubCategory
	}
	if productUpdateRequest.BrandId != nil {
		product.BrandId = *productUpdateRequest.BrandId
	}
	if image != nil {
		product.Image = mainImageUrl
	}
	if images != nil {
		product.Images = imageUrls
	}

	statusCode, err = p.productUsecase.UpdateById(*product)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Product updated successfully", nil)

}

func (p *ProductHandler) FindBySubCategoryId(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)
	subCategoryId := ctx.Param("subcategory_id")

	subCategoryIdInt, err := strconv.Atoi(subCategoryId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid sub category id")
	}

	products, statusCode, err := p.productUsecase.FindAllForMeBySubcategoryId(jwtClaims.UserId, subCategoryIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Products fetched successfully", responses.ToArrayOfProductResponse(products))
}

func (p *ProductHandler) FindByBrandId(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)
	brandId := ctx.Param("brand_id")

	brandIdInt, err := strconv.Atoi(brandId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid brand id")
	}

	products, statusCode, err := p.productUsecase.FindAllForMeByBrandId(jwtClaims.UserId, brandIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Products fetched successfully", responses.ToArrayOfProductResponse(products))
}

func (p *ProductHandler) SaveFrom1c(ctx echo.Context) error {
	var productCreateRequest domains.ProductDomainV2

	if err := helpers.BindAndValidate(ctx, &productCreateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	statusCode, err := p.productUsecase.SaveFrom1c(&productCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusCreated, "Saved successfully", nil)
}

func (p *ProductHandler) UpdateFrom1c(ctx echo.Context) error {
	var productUpdateRequest requests.ProductUpdateFrom1c
	cCode := ctx.Param("c_code")

	if err := helpers.BindAndValidate(ctx, &productUpdateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	product, statusCode, err := p.productUsecase.FindByCode(cCode)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if productUpdateRequest.Name != nil {
		product.Name = *productUpdateRequest.Name
	}
	if productUpdateRequest.Description != nil {
		product.Description = *productUpdateRequest.Description
	}
	if productUpdateRequest.Price != nil {
		product.Price = *productUpdateRequest.Price
	}
	if productUpdateRequest.Article != nil {
		product.Article = *productUpdateRequest.Article
	}
	if productUpdateRequest.CCode != nil {
		product.CCode = *productUpdateRequest.CCode
	}
	if productUpdateRequest.EdIzm != nil {
		product.EdIzm = *productUpdateRequest.EdIzm
	}

	statusCode, err = p.productUsecase.UpdateFrom1c(cCode, product)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Updated successfully", nil)
}

func (p *ProductHandler) FindByIdForUser(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)
	productId := ctx.Param("id")

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid product id")
	}

	product, statusCode, err := p.productUsecase.FindByIdForUser(productIdInt, jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Product fetched successfully", responses.FromProductDomain(product))
}
