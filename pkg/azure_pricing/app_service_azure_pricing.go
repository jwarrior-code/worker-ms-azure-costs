package azure_pricing

import (
	"encoding/json"
	"github.com/fatih/color"
	"io"
	"net/http"
	"worker-ms-azure-costs/utils/logger"
)

const (
	azureApiUrl = "https://prices.azure.com/api/retail/prices?%24filter=armRegionName%20eq%20%27eastus%27"
	batchSize   = 5000
)

type PortsServerAzurePricing interface {
	GetPricing() error
}

type service struct {
	repository ServicesAzurePricingRepository
}

func NewAzurePricingService(repository ServicesAzurePricingRepository) PortsServerAzurePricing {
	return &service{repository: repository}
}

func (s service) GetPricing() error {

	color.Blue("Iniciando worker para obtener datos de Azure Pricing...")

	// Obtener datos de precios de Azure
	azureProducts, err := s.fetchAzurePrices(azureApiUrl)
	if err != nil {
		logger.Error.Println("Error al obtener datos de precios de Azure:", err)
		return err
	}

	// Insertar datos en la base de datos
	err = s.repository.insertBulkAzureProducts(azureProducts)
	if err != nil {
		return err
	}

	color.Green("Información almacenada exitosamente.")

	err = s.repository.updateValorProductoProveedorNube()
	if err != nil {
		logger.Error.Println("Error al obtener actualizar valores:", err)
		return err
	}
	color.Green("Información actualizada exitosamente.")
	return nil
}

func (s service) fetchAzurePrices(url string) ([]AzurePriceItem, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error.Println("no pudo cerrar body http: ", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Items        []AzurePriceItem `json:"items"`
		NextPageLink string           `json:"nextPageLink"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	color.Blue("consultando pagina: ", result.NextPageLink)
	// Handle pagination
	if result.NextPageLink != "" {
		nextItems, err := s.fetchAzurePrices(result.NextPageLink)
		if err != nil {
			return nil, err
		}
		result.Items = append(result.Items, nextItems...)
	}

	return result.Items, nil
}
