package main

import (
	"fmt"
	"go-homework2/config"
)

func main() {
	fmt.Println("Application Start ....")
	AppConfig := config.GetConfig()

	fmt.Println(AppConfig.Host)
	fmt.Println(AppConfig.Port)
	fmt.Println(AppConfig.DB.User)
	fmt.Println(AppConfig.DB.Host)
	fmt.Println(AppConfig.DB.Port)

}
