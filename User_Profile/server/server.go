package server

import (
	"github.com/congmanh18/XNM-Express/User_Profile/migration"
	"github.com/congmanh18/XNM-Express/User_Profile/model"
	"github.com/congmanh18/XNM-Express/User_Profile/provider"
	"github.com/congmanh18/lucas-core/config"
	"github.com/congmanh18/lucas-core/transport/http"
	"github.com/congmanh18/lucas-core/transport/http/engine"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

// true để cho phép migrate
const enableMigration = false

func RunMigration(appProvider *provider.AppProvider, enableMigrate bool) {
	if enableMigrate {
		migration.Migration(appProvider.Postgres.Executor)
	}
}

func NewServer(serviceConf model.ServiceConfig, routes []route.GroupRoute) *http.Server {
	e := engine.NewEcho()
	return http.NewHttpServer(
		http.AddName(serviceConf.ServiceName),
		http.AddPort(serviceConf.ServicePort),
		http.AddEngine(e),
		http.AddGroupRoutes(routes),
	)
}

// LoadServiceConfig loads the service configuration.
func LoadServiceConfig(confPath string) model.ServiceConfig {
	var serviceConf model.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)
	return serviceConf
}

func Run(confPath string) {
	serviceConf := LoadServiceConfig(confPath)
	appProvider := InitializeAppProvider(serviceConf)
	docHandler := InitializeUserHandler(appProvider)
	routes := InitializeRoutes(docHandler)
	server := NewServer(serviceConf, routes)

	RunMigration(appProvider, enableMigration)
	server.Run()
}
