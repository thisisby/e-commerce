package domains

import "time"

type ProductStockDomain struct {
	Id                int
	CCode             string
	Date              time.Time
	TransactionType   int
	TransactionId     string
	Quantity          int
	TotalSum          float64
	TransactionStatus int
}

type ProductStockRepository interface {
	Save(productStock ProductStockDomain) error
	Update(productStock ProductStockDomain) error
	FindById(id string) (ProductStockDomain, error)
}

type ProductStockUsecase interface {
	Save(productStock ProductStockDomain) (int, error)
	Update(productStock ProductStockDomain) (int, error)
	FindById(id string) (ProductStockDomain, int, error)
}
