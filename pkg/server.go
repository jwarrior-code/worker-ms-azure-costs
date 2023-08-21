package pkg

import (
	"github.com/jmoiron/sqlx"
	"worker-ms-azure-costs/pkg/azure_pricing"
)

type Server struct {
	SrvAzurePricing azure_pricing.PortsServerAzurePricing
}

func NewServerWorkerAwsCosts(db *sqlx.DB) *Server {

	return &Server{
		SrvAzurePricing: azure_pricing.NewAzurePricingService(
			azure_pricing.FactoryStorage(db)),
	}
}
