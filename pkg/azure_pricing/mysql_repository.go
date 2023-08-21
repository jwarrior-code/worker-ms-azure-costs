package azure_pricing

import (
	"github.com/jmoiron/sqlx"
	"log"
	"worker-ms-azure-costs/utils/logger"
)

// sqlServer estructura de conexión a la BD de mssql
type mysql struct {
	DB *sqlx.DB
}

func newAzurePricingMysqlRepository(db *sqlx.DB) *mysql {
	return &mysql{
		DB: db,
	}
}

func (s mysql) insertBulkAzureProducts(products []AzurePriceItem) error {
	baseQuery := `INSERT INTO azure_prices (skuId, currencyCode, tierMinimumUnits, retailPrice, armRegionName, location, 
                          effectiveStartDate, productName, skuName, serviceName, armSkuName)
			  VALUES `

	for start := 0; start < len(products); start += batchSize {
		end := start + batchSize
		if end > len(products) {
			end = len(products)
		}
		err := s.insertAzureBatch(baseQuery, products[start:end])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s mysql) insertAzureBatch(baseQuery string, prices []AzurePriceItem) error {
	values := []interface{}{}
	insertPlaceholders := ""

	for _, price := range prices {
		insertPlaceholders += "(?,?,?,?,?,?,?,?,?,?,?),"
		values = append(values, price.SkuId, price.CurrencyCode, price.TierMinimumUnits, price.RetailPrice, price.ArmRegionName, price.Location, price.EffectiveStartDate, price.ProductName, price.SkuName, price.ServiceName, price.ArmSkuName)
	}

	insertPlaceholders = insertPlaceholders[:len(insertPlaceholders)-1] // Eliminar la última coma

	query := baseQuery + insertPlaceholders + ` ON DUPLICATE KEY UPDATE  
					skuId = VALUES(skuId), 
					currencyCode = VALUES(currencyCode), 
					tierMinimumUnits = VALUES(tierMinimumUnits), 
					retailPrice = VALUES(retailPrice), 
					armRegionName = VALUES(armRegionName),
					location = VALUES(location),
					effectiveStartDate = VALUES(effectiveStartDate),
					productName = VALUES(productName),
					skuName = VALUES(skuName),
					serviceName = VALUES(serviceName),
					armSkuName = VALUES(armSkuName),
					productName = VALUES(productName) ;`

	// Convertir los marcadores de posición
	query = s.DB.Rebind(query)

	_, err := s.DB.Exec(query, values...)
	if err != nil {
		logger.Error.Println("Error al insertar productos en bulk:", err)
		return err
	}
	return nil
}

func (s mysql) updateValorProductoProveedorNube() error {
	query := `UPDATE producto_proveedor_nube AS ppn
		LEFT JOIN azure_prices AS ap ON ppn.SKU = ap.skuId
		SET ppn.valor = ap.retailPrice
		WHERE ppn.proveedores_nube_idProveedor = 2
		AND (ap.skuId IS NOT NULL);
	`

	// Ejecuta el query
	result, err := s.DB.Exec(query)
	if err != nil {
		logger.Error.Println("Error al ejecutar la actualización de valor azure:", err)
		return err
	}

	// Puedes también obtener la cantidad de filas afectadas si lo necesitas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Println("Error al obtener filas afectadas:", err)
		return err
	}

	log.Printf("Se actualizó %d fila(s)", rowsAffected)
	return nil
}
