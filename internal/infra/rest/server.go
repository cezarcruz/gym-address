package rest

import (
	"github.com/cezarcruz/gym-address/internal/infra/rest/routes"
	"github.com/cezarcruz/gym-address/internal/infra/storage"
	"github.com/labstack/echo/v4"
)

func Start() {

	storage.Db = storage.MysqlStorageConfig()

	e := echo.New()

	routes.Load(e)

	e.Logger.Fatal(e.Start(":8082"))
}
