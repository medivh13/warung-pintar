package main

import (
	"context"

	usecases "warung-pintar/checkout-service/src/app/use_cases"

	"warung-pintar/checkout-service/src/infra/config"
	itemService "warung-pintar/checkout-service/src/infra/service/item"

	"warung-pintar/checkout-service/src/interface/rest"

	ms_log "warung-pintar/checkout-service/src/infra/log"

	checkoutUc "warung-pintar/checkout-service/src/app/use_cases/checkout"

	_ "github.com/joho/godotenv/autoload"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	itemsService := itemService.NewItemService()
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			CheckoutUseCase: checkoutUc.NewCheckoutUseCase(itemsService),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
