package operations

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/odedlaz/docker-registry-cli/core/config"
	mypath "github.com/odedlaz/docker-registry-cli/core/os/path"
)

// EncryptedRegistry blat
type OldEncryptedRegistry struct {
	Auth  string `json:"auth"`
	Email string `json:"email"`
}

// Settings the settings
type OldAuth map[string]OldEncryptedRegistry

func Migrate(dockercfgPath, configPath string) error {
	dockercfgData, err := ioutil.ReadFile(dockercfgPath)
	if err != nil {
		return err
	}
	auth := OldAuth{}
	if err = json.Unmarshal(dockercfgData, &auth); err != nil {
		return err
	}

	settings := config.EncryptedSettings{Auths: make(map[string]config.EncryptedRegistry)}
	for k, v := range auth {
		settings.Auths[k] = config.EncryptedRegistry{Auth: v.Auth}
	}

	encryptedSettingsData, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	if isDir, _ := mypath.IsDir(filepath.Dir(configPath)); !isDir {
		os.MkdirAll(filepath.Dir(configPath), 0700)
	}
	err = ioutil.WriteFile(configPath, encryptedSettingsData, 0600)

	return err
}
