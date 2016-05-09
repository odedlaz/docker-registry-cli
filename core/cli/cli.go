package cli

import (
	"fmt"
	"os"

	"github.com/odedlaz/docker-registry-cli/core"
	"github.com/odedlaz/docker-registry-cli/core/config"
	"gopkg.in/alecthomas/kingpin.v2"
)

// global command names

var (
	AddCommandName         = "add"
	RemoveCommandName      = "remove"
	ListCommandName        = "list"
	RegistryName           = "registry"
	VersionFlagName        = "version"
	ConfigFilenameFlagName = "config"
)

// global flags
var (
	App            = kingpin.New("docker-registry", "A cli for docker registries")
	ConfigFilename = App.Flag(ConfigFilenameFlagName, "path to config file").Default(config.DefaultFilename).String()
	Registry       = App.Flag(RegistryName, "registry to act upon").String()
	showVersion    = App.Flag(VersionFlagName, "show version and quit").Action(printVersion).Bool()
)

// registry commands
var (
	registryCommand = App.Command(RegistryName, "registry commands")

	AddRegistryCommandName = fmt.Sprintf("%s %s", RegistryName, AddCommandName)
	addRegistryCommand     = registryCommand.Command(AddCommandName, "add a new local registry entry")

	RemoveRegistryCommandName = fmt.Sprintf("%s %s", RegistryName, RemoveCommandName)
	removeRegistryCommand     = registryCommand.Command(RemoveCommandName, "delete local registry entry")

	ListRegistryCommandName = fmt.Sprintf("%s %s", RegistryName, ListCommandName)
	listRegistryCommand     = registryCommand.Command(ListCommandName, "list local registry entries")
)

// repository commands
var (
	RepositoryCommandName = "repository"
	repositoryCommand     = App.Command(RepositoryCommandName, "repository commands")

	ListRepositoriesCommandName = fmt.Sprintf("%s %s", RepositoryCommandName, ListCommandName)
	listRepositoryCommand       = repositoryCommand.Command(ListCommandName, "list repositories on remote registry")
)

// other commands
var (
	MigrateCommandName           = "migrate"
	migrateCommand               = App.Command(MigrateCommandName, "migrate dockercfg to config.json")
	MigrateCommandConfigFilename = migrateCommand.Arg("dockercfg", "path to dockercfg file").Default(config.DefaultOldFilename).String()
)

func printVersion(ctx *kingpin.ParseContext) error {
	fmt.Printf("docker-registry-cli %s\n",
		core.Version)
	os.Exit(0)
	return nil
}

// RegistrySet checks if the registry flag is set. if not, raises an error and quits
func RegistrySet() {
	if *Registry == "" {
		App.Errorf("required flag --%s not provided, try --help", RegistryName)
		os.Exit(1)
	}
}

// RegistrySetAndExists checks if the registry flag is set and that it's registered. if not, raises an error and quits
func RegistrySetAndExists(settings *config.Settings) {
	RegistrySet()
	if _, ok := settings.Auths[*Registry]; !ok {
		App.Errorf("Registry doesn't exist. maybe add it?")
		os.Exit(1)
	}
}
