package main

import (
	"os"
	"time"

	"github.com/felixa1996/go-plate/infrastructure/database"

	"github.com/felixa1996/go-plate/infrastructure"
	"github.com/felixa1996/go-plate/infrastructure/log"
	"github.com/felixa1996/go-plate/infrastructure/router"
	"github.com/felixa1996/go-plate/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres)
	//DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort("80").
		WebServer(router.InstanceGorillaMux).
		Start()
}
