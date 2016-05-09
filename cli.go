package main

import (
	"fmt"
	"os"

	"github.com/odedlaz/docker-registry-cli/core"
	"github.com/odedlaz/docker-registry-cli/core/config"
	"github.com/odedlaz/docker-registry-cli/operations"
	"gopkg.in/alecthomas/kingpin.v2"
)

const registryCommandName = "registry"
const repositoryCommandName = "repository"
const addCommandName = "add"
const removeCommandName = "remove"
const listCommandName = "list"
const migrateCommandName = "migrate"

var (
	addRegistry                 = fmt.Sprintf("%s %s", registryCommandName, addCommandName)
	removeRegistry              = fmt.Sprintf("%s %s", registryCommandName, removeCommandName)
	listRepositories            = fmt.Sprintf("%s %s", repositoryCommandName, listCommandName)
	app                         = kingpin.New("docker-registry", "A command-line un-templating application.")
	configFilename              = app.Flag("config", "path to config file").Default(config.DefaultFilename).String()
	showVersion                 = app.Flag("version", "show version and quit").Action(printVersion).Bool()
	registry                    = app.Flag(registryCommandName, "registry to act upon").String()
	registryCommand             = app.Command(registryCommandName, "registry commands")
	addRegistryCommand          = registryCommand.Command(addCommandName, "add a new registry")
	removeRegistryCommand       = registryCommand.Command(removeCommandName, "delete an existing registry")
	repositoryCommand           = app.Command(repositoryCommandName, "repository commands")
	listRepositoryCommand       = repositoryCommand.Command(listCommandName, "list repositories")
	migrateCommand              = app.Command(migrateCommandName, "migrate dockercfg to config.json")
	MigrateCommandDockercfgPath = migrateCommand.Arg("dockercfg", "path to dockercfg file").Default(config.DefaultOldFilename).String()
)

func printVersion(ctx *kingpin.ParseContext) error {
	fmt.Printf("docker-registry-cli %s\n",
		core.Version)
	os.Exit(0)
	return nil
}

func EnsureRegistrySet(settings *config.Settings) {
	if *registry == "" {
		app.Errorf("required flag --%s not provided, try --help", registryCommandName)
		os.Exit(1)
	}
	if _, ok := settings.Auths[*registry]; !ok {
		app.Errorf("Registry doesn't exist. maybe add it?")
		os.Exit(1)
	}
}

func main() {
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))
	settings, err := config.Load(*configFilename)

	if parsed != addRegistry && parsed != migrateCommandName {
		if len((*settings).Auths) == 0 {
			app.Errorf(err.Error())
			os.Exit(1)
		}
	}

	switch parsed {
	case addRegistry:
		EnsureRegistrySet(settings)
		core.Must(operations.AddRegistry(*settings))
	case removeRegistry:
		EnsureRegistrySet(settings)
		core.Must(operations.RemoveRegistry(*settings, *registry))
	case listRepositories:
		EnsureRegistrySet(settings)
		core.Must(operations.ListRepositories(*registry, *settings))
	case migrateCommandName:
		core.Must(operations.Migrate(*MigrateCommandDockercfgPath, *configFilename))
	}
}
