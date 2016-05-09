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

var (
	addRegistry           = fmt.Sprintf("%s %s", registryCommandName, addCommandName)
	removeRegistry        = fmt.Sprintf("%s %s", registryCommandName, removeCommandName)
	listRepositories      = fmt.Sprintf("%s %s", repositoryCommandName, listCommandName)
	app                   = kingpin.New("docker-registry", "A command-line un-templating application.")
	configFilename        = app.Flag("config", "path to config file").Default(config.DefaultFilename).String()
	showVersion           = app.Flag("version", "show version and quit").Action(printVersion).Bool()
	registry              = app.Flag(registryCommandName, "registry to act upon").String()
	registryCommand       = app.Command(registryCommandName, "registry commands")
	addRegistryCommand    = registryCommand.Command(addCommandName, "add a new registry")
	removeRegistryCommand = registryCommand.Command(removeCommandName, "delete an existing registry")
	repositoryCommand     = app.Command(repositoryCommandName, "repository commands")
	listRepositoryCommand = repositoryCommand.Command(listCommandName, "list repositories")
)

func printVersion(ctx *kingpin.ParseContext) error {
	fmt.Printf("docker-registry-cli %s\n",
		core.Version)
	os.Exit(0)
	return nil
}

func main() {
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))
	settings, _ := config.Load(*configFilename)
	if parsed != addRegistry {
		if len((*settings).Auths) == 0 {
			app.Errorf("No registered docker registries, maybe add one?")
			os.Exit(1)
		}
		if _, ok := settings.Auths[*registry]; !ok {
			app.Errorf("Registry doesn't exist. maybe add it?")
			os.Exit(1)
		}

	}

	if *registry == "" {
		app.Errorf("required flag --%s not provided, try --help", registryCommandName)
		os.Exit(1)
	}

	switch parsed {
	case addRegistry:
		core.Must(operations.AddRegistry(*settings))
	case removeRegistry:
		core.Must(operations.RemoveRegistry(*settings, *registry))
	case listRepositories:
		core.Must(operations.ListRepositories(*registry, *settings))
	}
}
