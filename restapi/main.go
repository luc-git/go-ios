package main

import (
	"os"

	"github.com/luc-git/go-ios/restapi/api"
	_ "github.com/luc-git/go-ios/restapi/docs"
	log "github.com/sirupsen/logrus"
)

// @title           Go-iOS API
// @version         0.01
// @description     Exposes go-ios features as REST API calls.
// @termsOfService  https://github.com/luc-git/go-ios

// @contact.name   Daniel Paulus
// @contact.url    https://github.com/luc-git/go-ios

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	log.WithFields(log.Fields{"args": os.Args, "version": api.GetVersion()}).Infof("starting go-iOS-API")
	api.Main()
}
