package main

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/lengwh/event-sync/common/opio"
	"github.com/lengwh/event-sync/config"
	"github.com/lengwh/event-sync/database"
	"github.com/lengwh/event-sync/flags"
	"github.com/urfave/cli/v2"
)

func NewCli() *cli.App {
	var flags = flags.Flags
	return &cli.App{
		Version:              "v1.0",
		Description:          "An Event Sync Servce",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "runs the database migrations ",
				Action:      runMigrations,
			},
		},
	}
}

func runMigrations(ctx *cli.Context) error {
	log.Info("Running migrations")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("Failed to load config", "err", err)
		return err
	}
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("Failed to connect to database", "err", err)
		return err
	}
	log.Info("Connected to database")

	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return db.ExecuteSQLMigration(cfg.Migrations)
}
