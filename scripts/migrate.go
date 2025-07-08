package main

import (
	"auth-go/internal/config"
	"auth-go/internal/models"
	"auth-go/internal/provider"
	"flag"
	"fmt"
	"gorm.io/gorm"
)

func main() {

	action := flag.String("action", "up", "Action to perform: up or down")

	flag.Parse()

	cfg := config.New()
	conn := provider.GetConnection(cfg)
	if conn == nil {
		panic("Failed to connect to the db")
	}

	switch *action {
	case "up":
		migrateUp(conn)
	case "down":
		migrateDown(conn)
	}
}

func migrateUp(db *gorm.DB) {
	fmt.Println("Migrating up...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Printf("Error migrating up: %v\n", err)
	} else {
		fmt.Println("Migration up completed successfully.")
	}
}

func migrateDown(db *gorm.DB) {
	fmt.Println("Migrating down...")
	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		fmt.Printf("Error migrating down: %v\n", err)
	} else {
		fmt.Println("Migration down completed successfully.")
	}
}
