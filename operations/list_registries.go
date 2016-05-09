package operations

import (
	"fmt"

	"github.com/odedlaz/docker-registry-cli/core/config"
)

func ListRegistries(settings config.Settings) error {
	for registry := range settings.Auths {
		fmt.Println(registry)
	}
	return nil
}
