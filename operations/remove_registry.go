package operations

import "github.com/odedlaz/docker-registry-cli/core/config"

func RemoveRegistry(registry string, settings config.Settings) error {
	settings.RemoveRegistry(registry)
	return settings.Save()
}
