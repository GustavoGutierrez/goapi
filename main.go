package main

import (
	"fmt"
	"log"

	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	// _ "gorm.io/driver/mysql"

	"github.com/GustavoGutierrez/goapi/config"
	"github.com/GustavoGutierrez/goapi/infrastructure/datastore"
	"github.com/GustavoGutierrez/goapi/infrastructure/router"
	"github.com/GustavoGutierrez/goapi/registry"
)

func main() {
	// config.ReadConfig()

	db := datastore.NewDB()
	// db.LogMode(true)
	db.Debug()
	// defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
