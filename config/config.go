package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
)

var Config Configuration

func GetConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println(err, "failed to read configuration")
		os.Exit(1)
	}
	return configuration
}

func getFileName() string {
	strfilename := fmt.Sprintf("config.json")
	return strfilename
}
