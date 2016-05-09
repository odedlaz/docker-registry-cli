package main

import (
	"os"

	"github.com/odedlaz/docker-registry-cli/core"
	"github.com/odedlaz/docker-registry-cli/core/cli"
	"github.com/odedlaz/docker-registry-cli/core/config"
	"github.com/odedlaz/docker-registry-cli/operations"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	parsed := kingpin.MustParse(cli.App.Parse(os.Args[1:]))
	settings, err := config.Load(*cli.ConfigFilename)

	if parsed != cli.AddRegistryCommandName && parsed != cli.MigrateCommandName {
		if len((*settings).Auths) == 0 {
			cli.App.Errorf(err.Error())
			os.Exit(1)
		}
	}

	switch parsed {
	case cli.AddRegistryCommandName:
		cli.RegistrySet()
		core.Must(operations.AddRegistry(*cli.Registry, *settings))
	case cli.RemoveRegistryCommandName:
		cli.RegistrySetAndExists(settings)
		core.Must(operations.RemoveRegistry(*cli.Registry, *settings))
	case cli.ListRegistryCommandName:
		core.Must(operations.ListRegistries(*settings))
	case cli.ListRepositoriesCommandName:
		cli.RegistrySetAndExists(settings)
		core.Must(operations.ListRepositories(*cli.Registry, *settings))
	case cli.MigrateCommandName:
		core.Must(operations.Migrate(*cli.MigrateCommandConfigFilename, *cli.ConfigFilename))
	default:
		cli.App.Errorf("unknown option: %s", parsed)
	}
}
