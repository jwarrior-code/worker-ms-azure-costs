package worker

import (
	"worker-ms-azure-costs/pkg"
	"worker-ms-azure-costs/utils/logger"
)

type Worker struct {
	srv *pkg.Server
}

func NewWorker(srv *pkg.Server) IWorker {
	return &Worker{srv: srv}
}

func (w Worker) Execute() {

	err := w.srv.SrvAzurePricing.GetPricing()
	if err != nil {
		logger.Error.Println("no se puede actualizar costos de azure", err)
		return
	}

}
