package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	mypath "github.com/odedlaz/docker-registry-cli/core/os/path"
)

var (
	// DefaultFilename is the default path for the docker config
	DefaultFilename    = fmt.Sprintf("%s/.docker/config.json", os.Getenv("HOME"))
	DefaultOldFilename = fmt.Sprintf("%s/.dockercfg", os.Getenv("HOME"))
)

// EncryptedRegistry blat
type EncryptedRegistry struct {
	Auth string `json:"auth"`
}

// EncryptedSettings blat
type EncryptedSettings struct {
	Auths map[string]EncryptedRegistry `json:"auths"`
}

// Registry represents a single registry
type Registry struct {
	Username string
	Password string
}

// Settings the settings
type Settings struct {
	Path  string
	Auths map[string]Registry
}

// Encrypt encrypts a given settings file
func (s *Settings) Encrypt() *EncryptedSettings {
	encrypted := EncryptedSettings{Auths: make(map[string]EncryptedRegistry)}
	for k, v := range s.Auths {
		text := fmt.Sprintf("%s:%s", v.Username, v.Password)
		encrypted.Auths[k] = EncryptedRegistry{Auth: base64.StdEncoding.EncodeToString([]byte(text))}
	}
	return &encrypted
}

// Save settings
func (s *Settings) Save() error {
	data, err := json.Marshal(s.Encrypt())
	if err != nil {
		return err
	}
	mode, err := mypath.FileMode(s.Path)
	if err != nil {
		mode = 0600
	}
	err = ioutil.WriteFile(s.Path, data, mode)
	return err
}

// AddRegistry adds a new registry to the settings
func (s *Settings) AddRegistry(name string, registry Registry) {
	s.Auths[name] = registry
}

// RemoveRegistry adds a new registry to the settings
func (s *Settings) RemoveRegistry(name string) {
	delete(s.Auths, name)
}

func New(filename string) *Settings {
	s := Settings{Path: filename, Auths: make(map[string]Registry)}
	return &s
}

// New the yaml settings from given path
func Load(filename string) (*Settings, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if _, err := os.Stat(DefaultOldFilename); err == nil {
			return New(filename), fmt.Errorf("config file doesn't exist at: %s, but detected dockercfg at: %s", filename, DefaultFilename)
		}
		return New(filename), fmt.Errorf("config file doesn't exist at: %s", filename)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	encrypted := EncryptedSettings{}
	if err = json.Unmarshal(data, &encrypted); err != nil {
		return nil, err
	}
	settings := New(filename)
	for k, v := range encrypted.Auths {
		decoded, err := base64.StdEncoding.DecodeString(v.Auth)
		if err != nil {
			return nil, err
		}

		parts := strings.Split(string(decoded), ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid registry %s value", k)
		}
		settings.Auths[k] = Registry{Username: parts[0], Password: parts[1]}
	}
	return settings, nil
}
