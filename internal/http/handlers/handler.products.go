package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/aws"
	"github.com/labstack/echo/v4"
	"net/http"
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
		Name:        productCreateRequest.Name,
		Description: productCreateRequest.Description,
		Price:       productCreateRequest.Price,
		Image:       mainImageUrl,
		Images:      imageUrls,
	}

	statusCode, err := p.productUsecase.Save(productDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, 200, "success", productDomain)
}

func (p *ProductHandler) FindAllForMe(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	products, statusCode, err := p.productUsecase.FindAllForMe(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Products fetched successfully", responses.ToArrayOfProductResponse(products))
}
