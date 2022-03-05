package initialize

import (
	"github.com/raghav-rathi/LookItUp/tree/main/newcomponent_scanner/Backend/internal/http"

	"github.com/raghav-rathi/LookItUp/tree/main/newcomponent_scanner/Backend/pkg/db/mysql"

	"github.com/raghav-rathi/LookItUp/tree/main/newcomponent_scanner/Backend/internal/initialize/http"
)

func InitializeApp() {

	// setting configurations
	// config.InitConfig()

	// connecting to mysql database
	mysql.CreateClient()

	// initializing echo server
	http.InitHTTPServer()
}
