package requests

type ProductCreateRequest struct {
	Name          string  `form:"name" validate:"required"`
	Description   string  `form:"description" validate:"required"`
	Price         float64 `form:"price" validate:"required"`
	SubCategoryId int     `form:"sub_category_id" validate:"required"`
	BrandId       int     `form:"brand_id" validate:"required"`
}

type ProductUpdateRequest struct {
	Weight              *float64 `form:"weight"`
	SubCategory         *int     `form:"sub_category_id"`
	BrandId             *int     `form:"brand_id"`
	Volume              *float64 `form:"volume"`
	CountryOfProduction *string  `form:"country_of_production"`
	Sex                 *string  `form:"sex"`
	AttributesIds       []int    `form:"attributes_ids"`
}

type ProductDeleteAttributeRequest struct {
	AttributeIds []int `json:"attribute_ids"`
}

type ProductUpdateFrom1c struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Article     *string  `json:"article"`
	CCode       *string  `json:"c_code"`
	EdIzm       *string  `json:"ed_izm"`
}
