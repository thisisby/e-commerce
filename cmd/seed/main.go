package main

import (
	"ga_marketplace/cmd/seed/seeders"
	"ga_marketplace/internal/config"
	"ga_marketplace/internal/utils"
	"log/slog"
	"os"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Logger initialized")

	if err := config.InitializeAppConfig(); err != nil {
		slog.Error("failed to initialize app config: ", err)
		return
	}
}

func main() {
	conn, err := utils.SetupPostgreConnection()
	if err != nil {
		slog.Error("[Seed]: failed to connect to DB", err)
		return
	}

	defer conn.Close()

	slog.Info("[Seed]: seeding...")

	seeder := seeders.NewSeeder(conn)
	err = seeder.RolesSeeder(seeders.RolesData)
	if err != nil {
		slog.Error("[Seed]: failed to seed roles data", err)
	}

	err = seeder.CitiesSeeder(seeders.CitiesData)
	if err != nil {
		slog.Error("[Seed]: failed to seed cities data", err)
	}

	err = seeder.ProductsSeeder(seeders.ProductsData)
	if err != nil {
		slog.Error("[Seed]: failed to seed products data", err)
	}

	err = seeder.UsersSeeder(seeders.UsersData)
	if err != nil {
		slog.Error("[Seed]: failed to seed users data", err)
	}

	slog.Info("[Seed]: seeding completed")
}
