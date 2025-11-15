package cmd

import (
	"ecom_project/config"
	"ecom_project/rest"
)

func Serve() {

	config := config.GetConfig()
	rest.Start(config)

}
