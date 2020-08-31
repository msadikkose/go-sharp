package main

import (
	"go-sharp/app"
	"go-sharp/dal/db"
	"go-sharp/router"

)

// @title Swagger User Service
// @version 1.0
// @description This is a sample server User Crud.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://profe.com.tr
// @contact.email sadik.kose@profe.com.tr

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {

	app.InitConfig()
	dbHandler:= &db.DbHandler{}
	dbHandler.Migrate()
	engine:=router.GinRouter().InitRouter(router.InjectBusinessManagerFactory)
	engine.Run(app.Config.Port)
}
