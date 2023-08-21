package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jmoiron/sqlx"
	"worker-ms-azure-costs/pkg"
	"worker-ms-azure-costs/utils/dbx"
	"worker-ms-azure-costs/utils/logger"
	"worker-ms-azure-costs/worker"
)

func main() {
	fmt.Println("test")
	color.Blue("worker-ms-azure-costs v1.0.0")
	db := dbx.GetConnection()
	srv := pkg.NewServerWorkerAwsCosts(db)
	wk := worker.NewWorker(srv)
	wk.Execute()

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			logger.Error.Println("error desconectando de la base de datos", err)
		}
	}(db)
}
