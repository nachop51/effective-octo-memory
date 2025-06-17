package main

import (
	"flag"
	"log"

	"server/app"
	"server/config"
	"server/domains/accounts"
	"server/domains/users"
	"server/pkg/db"
)

func main() {
	longMigrate := flag.Bool("migrate", false, "Run database migrations")
	shortMigrate := flag.Bool("m", false, "Run database migrations (short)")
	flag.Parse()

	migrate := *longMigrate || *shortMigrate

	appConfig, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dps, err := config.NewDependencies(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	if migrate {
		log.Println("Running database migrations...")
		db.AutoMigrate(dps.DB, &users.User{}, &accounts.Account{})
		log.Println("Database migrations completed.")
		return
	}

	application := app.New(dps)
	application.Setup()

	log.Printf("Server starting on %s:%d", appConfig.Server.Host, appConfig.Server.Port)
	if err := application.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
