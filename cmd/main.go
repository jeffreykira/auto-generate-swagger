package main

import (
	_ "github.com/jeffreykira/log-management/docs"

	"github.com/jeffreykira/log-management/service/server/api/router"
	// log "github.com/mgutz/logxi/v1"
)

// var logger log.Logger

// @title Sample API
// @version 1.0.0
// @description Documentation of Sample API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/jeffreykira/log-management.git

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5608
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := router.NewRouter()
	r.Run(":5608")
}
