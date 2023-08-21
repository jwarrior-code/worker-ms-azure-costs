package azure_pricing

import (
	"github.com/jmoiron/sqlx"
	"worker-ms-azure-costs/utils/logger"
)

const (
	Mysql = "mysql"
)

type ServicesAzurePricingRepository interface {
	insertBulkAzureProducts(products []AzurePriceItem) error
	updateValorProductoProveedorNube() error
}

func FactoryStorage(db *sqlx.DB) ServicesAzurePricingRepository {
	var s ServicesAzurePricingRepository
	engine := db.DriverName()
	switch engine {
	case Mysql:
		return newAzurePricingMysqlRepository(db)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
