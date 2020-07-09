package main

import (
	"os"

	_ "github.com/jeffreykira/log-management/docs"

	"github.com/jeffreykira/log-management/core/api/router"
	log "github.com/mgutz/logxi/v1"
)

var logger log.Logger

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
	// create a logger with a unique identifier which
	// can be enabled from environment variables
	logger = log.New("main")
	logger.Debug("start")

	// specify a writer, use NewConcurrentWriter if it is not concurrent safe
	modelLogger := log.NewLogger(log.NewConcurrentWriter(os.Stdout), "model")
	modelLogger.Info("test", "show", "info level")

	fruit := "apple"
	languages := []string{"go", "javascript"}
	if log.IsDebug() {
		// use key-value pairs after message
		logger.Debug("OK", "fruit", fruit, "languages", languages)
	}

	r := router.NewRouter()
	r.Run(":5608")
}
