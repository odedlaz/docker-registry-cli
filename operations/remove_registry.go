package operations

import "github.com/odedlaz/docker-registry-cli/core/config"

func RemoveRegistry(settings config.Settings, registry string) error {
	settings.RemoveRegistry(registry)
	return settings.Save()
}
