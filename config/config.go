package config

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/lengwh/event-sync/flags"
	"github.com/urfave/cli/v2"
)

const (
	defaultConfirmations = 64
	defaultLoopInterval  = 5000
	TreasureManagerAddr  = "0x90F274819130d3E33D4c85008035fCEf80847302"
)

type Config struct {
	Migrations    string
	Chain         ChainConfig
	MasterDB      DBConfig
	SlaveDB       DBConfig
	SlaveDbEnable bool
}
type ChainConfig struct {
	ChainRpcUrl    string
	ChainId        uint64
	StartingHeight uint64
	Confirmations  uint64
	BlockSteps     uint64
	Contracts      []common.Address
	LoopInterval   time.Duration
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func LoadConfig(ctx *cli.Context) (Config, error) {
	var cfg Config
	cfg = NewConfig(ctx)
	if cfg.Chain.Confirmations == 0 {
		cfg.Chain.Confirmations = defaultConfirmations
	}

	if cfg.Chain.LoopInterval == 0 {
		cfg.Chain.LoopInterval = defaultLoopInterval
	}
	log.Info("Config", "cfg", cfg)
	return cfg, nil
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationFlag.Name),
		Chain: ChainConfig{
			ChainRpcUrl:    ctx.String(flags.ChainRpcUrlFlag.Name),
			ChainId:        ctx.Uint64(flags.ChainIdFlag.Name),
			StartingHeight: ctx.Uint64(flags.StartingHeightFlag.Name),
			Confirmations:  ctx.Uint64(flags.ConfirmationsFlag.Name),
			BlockSteps:     ctx.Uint64(flags.BlocksStepFlag.Name),
			Contracts:      LoadContracts(),
		},
		MasterDB: DBConfig{
			Host:     ctx.String(flags.MasterDbHostFlag.Name),
			Port:     ctx.Int(flags.MasterDbPortFlag.Name),
			User:     ctx.String(flags.MasterDbUserFlag.Name),
			Password: ctx.String(flags.MasterDbPasswordFlag.Name),
			Name:     ctx.String(flags.MasterDbNameFlag.Name),
		},
	}
}

func LoadContracts() []common.Address {
	var contracts []common.Address
	contracts = append(contracts, common.HexToAddress(TreasureManagerAddr))
	return contracts
}
