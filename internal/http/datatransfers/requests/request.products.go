package requests

type ProductCreateRequest struct {
	Name        string  `form:"name" validate:"required"`
	Description string  `form:"description" validate:"required"`
	Price       float64 `form:"price" validate:"required"`
}
