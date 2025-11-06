package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

const envVarPrefix = "EVENT_SYNC"

var Flags []cli.Flag

func init() {
	Flags = append(requireFlags, optionalFlags...)
}

var requireFlags = []cli.Flag{
	MigrationFlag,
	MasterDbHostFlag,
	MasterDbPortFlag,
	MasterDbUserFlag,
	MasterDbPasswordFlag,
	MasterDbNameFlag,
}

var optionalFlags = []cli.Flag{
	ChainRpcUrlFlag,
	ChainIdFlag,
	StartingHeightFlag,
	ConfirmationsFlag,
	LoopIntervalFlag,
	BlocksStepFlag,
}

func prefixEnvVars(key string) []string {
	return []string{envVarPrefix + "_" + key}
}

var (
	MigrationFlag = &cli.StringFlag{
		Name:    "migration-dir",
		Usage:   "Path to migration file",
		EnvVars: prefixEnvVars("MIGRATION_DIR"),
		Value:   "./migrations",
	}

	ChainRpcUrlFlag = &cli.StringFlag{
		Name:     "chain-rpc-url",
		Usage:    "Chain RPC URL",
		EnvVars:  prefixEnvVars("CHAIN_RPC"),
		Required: true,
	}

	ChainIdFlag = &cli.UintFlag{
		Name:     "chain-id",
		Usage:    "chain-id",
		EnvVars:  prefixEnvVars("CHAIN_ID"),
		Value:    1,
		Required: true,
	}

	StartingHeightFlag = &cli.UintFlag{
		Name:    "starting-height",
		Usage:   "starting height",
		EnvVars: prefixEnvVars("STARTING_HEIGHT"),
	}

	ConfirmationsFlag = &cli.Uint64Flag{
		Name:    "confirmations",
		Usage:   "The confirmation depth of l1",
		EnvVars: prefixEnvVars("CONFIRMATIONS"),
		Value:   64,
	}

	LoopIntervalFlag = &cli.DurationFlag{
		Name:    "loop-interval",
		Usage:   "The interval of synchronization",
		EnvVars: prefixEnvVars("LOOP_INTERVAL"),
		Value:   time.Second * 5,
	}

	BlocksStepFlag = &cli.UintFlag{
		Name:    "blocks-step",
		Usage:   "Scanner blocks step",
		EnvVars: prefixEnvVars("BLOCKS_STEP"),
		Value:   5,
	}

	// MasterDbHostFlag MasterDb Flags
	MasterDbHostFlag = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_HOST"),
		Required: true,
	}
	MasterDbPortFlag = &cli.IntFlag{
		Name:     "master-db-port",
		Usage:    "The port of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PORT"),
		Required: true,
	}
	MasterDbUserFlag = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "The user of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_USER"),
		Required: true,
	}
	MasterDbPasswordFlag = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PASSWORD"),
		Required: true,
	}
	MasterDbNameFlag = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "The db name of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_NAME"),
		Required: true,
	}
)
