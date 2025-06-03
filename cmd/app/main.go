package main

import (
	"fmt"
	"go-absen-be/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	// if viperConfig.GetBool("database.auto_migrate") {
	// 	if err := db.AutoMigrate(entity.RegisteredEntities...); err != nil {
	// 		log.Fatalf("AutoMigrate failed: %v", err)
	// 	}
	// 	fmt.Println("✅ Database migrated!")
	// 	} else {
	// 		fmt.Println("🔒 AutoMigrate skipped (production mode)")
	// 	}
	validate := config.NewValidator(viperConfig)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}