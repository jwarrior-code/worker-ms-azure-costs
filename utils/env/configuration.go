package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	once   sync.Once
	config = &configuration{}
)

type configuration struct {
	App App `json:"app"`
	DB  DB  `json:"db"`
}

type App struct {
	ServiceName       string `json:"service_name"`
	PathLog           string `json:"path_log"`
	LogReviewInterval int    `json:"log_review_interval"`
}

type DB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Instance string `json:"instance"`
	IsSecure bool   `json:"is_secure"`
	SSLMode  string `json:"ssl_mode"`
}

func NewConfiguration() *configuration {
	fromFile()
	return config
}

// LoadConfiguration lee el archivo configuration.json
// y lo carga en un objeto de la estructura Configuration
func fromFile() {
	once.Do(func() {
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}

		if config.DB.Engine == "" {
			log.Fatal("no se ha cargado la información de configuración")
		}
	})
}
