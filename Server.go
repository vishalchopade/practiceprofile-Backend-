package main

import (
	"log"
	"net/http"

	"survayData/api"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main(){
	// Making a new echo Server
	e := echo.New()

	// e.Use(middleware.Recover())

	confighelper.InitViper()

	// for Development environment isProduction flag set to 'false' in config.json file
	// for Staging and Production environment isProduction flag set to 'true' in config.json file
	// Parameters: filepath, environment, no.of files to be generated, filesize in MB, how many days file should live in system, safemode
	logginghelper.Init(confighelper.GetConfig("logFilePath"), viper.GetBool("isProduction"), viper.GetInt("log_maxBackupCount"), viper.GetInt("log_maxBackupFileSize"), viper.GetInt("log_maxAgeForBackupFiles"), true)	

	// Hiding ECHO banner from terminal
	e.HideBanner = true

	//Bind API
	api.Init(e)

	serverPort := confighelper.GetConfig("serverPort")

	logginghelper.LogInfo("server started on localhost:", serverPort)
	
	// Profiling port
	go func() {
		http.ListenAndServe(":"+confighelper.GetConfig("serverProfilingPort"), nil)
	}()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	err := e.Start(":" + serverPort)
	if err != nil {
		log.Fatal(err)
	}
}