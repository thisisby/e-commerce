package requests

type ProductCreateRequest struct {
	Name          string  `form:"name" validate:"required"`
	Description   string  `form:"description" validate:"required"`
	Price         float64 `form:"price" validate:"required"`
	Stock         int     `form:"stock" validate:"required"`
	SubCategoryId int     `form:"sub_category_id" validate:"required"`
	BrandId       int     `form:"brand_id" validate:"required"`
}

type ProductUpdateRequest struct {
	Name        *string  `form:"name"`
	Description *string  `form:"description"`
	Price       *float64 `form:"price"`
	Stock       *int     `form:"stock"`
	SubCategory *int     `form:"sub_category_id"`
	BrandId     *int     `form:"brand_id"`
}
