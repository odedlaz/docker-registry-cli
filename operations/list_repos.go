package operations

import (
	"encoding/json"
	"fmt"

	"github.com/odedlaz/docker-registry-cli/core/api"
	"github.com/odedlaz/docker-registry-cli/core/config"
)

type RepositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func ListRepositories(registry string, settings config.Settings) error {
	body, err, resp := api.Call(registry,
		settings.Auths[registry].Username,
		settings.Auths[registry].Password,
		"_catalog",
		api.Get)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Couldn't list repositories: %v", string(body))
	}

	repos := RepositoriesResponse{}
	err = json.Unmarshal(body, &repos)

	if err != nil {
		return err
	}
	for _, repo := range repos.Repositories {
		fmt.Println(repo)
	}

	return nil
}
