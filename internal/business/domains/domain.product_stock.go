package domains

import "time"

type ProductStockDomain struct {
	TransactionId string
	CustomerId    int
	Date          time.Time
	Active        bool
	Items         []ProductStockItemDomain
	OrderId       int
}

type ProductStockItemDomain struct {
	TransactionId   string
	ProductCode     string
	Quantity        int
	Amount          float64
	TransactionType int
}

type ProductStockRepository interface {
	Save(productStock ProductStockDomain) error
	Update(productStock ProductStockDomain, transactionId string) error
	FindById(id string) (ProductStockDomain, error)
	FindStockItem(transactionId string, productId string) (ProductStockItemDomain, error)
	UpdateProductStockItem(item ProductStockItemDomain, transactionId string, productId string) error
}

type ProductStockUsecase interface {
	Save(productStock ProductStockDomain) (int, error)
	Update(productStock ProductStockDomain, transactionId string) (int, error)
	FindById(id string) (ProductStockDomain, int, error)
	FindStockItem(transactionId string, productId string) (ProductStockItemDomain, int, error)
	UpdateProductStockItem(item ProductStockItemDomain, transactionId string, productId string) (int, error)
}
