package main

import (
	"flag"
	"log"
	"os"

	"github.com/congmanh18/XNM-Express/User_Profile/server"
)

func main() {
	// Dùng cách cờ để đọc file config
	// go run main.go -config="./docs-service/deploy/prod_conf.env"
	configPathFlag := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	// Dùng cách set biến môi trường
	// export CONFIG_PATH="./docs-service/deploy/prod_conf.env"
	// go run main.go
	configPath := *configPathFlag
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}
	// code local
	// if configPath == "" {
	// 	configPath = "D://Go-Basic//nextedu//docs-service//deploy//local_conf.env"
	// }

	log.Printf("Starting server with config: %s", configPath)
	server.Run(configPath)
}

// Cách 1:
// go run main.go -config="./config/prod_conf.env"

// Cách 2:
// export CONFIG_PATH="./docs-service/deploy/prod_conf.env"
// go run main.go
