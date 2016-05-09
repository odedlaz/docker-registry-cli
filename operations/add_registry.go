package operations

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/odedlaz/docker-registry-cli/core/config"
)

func AddRegistry(registry string, settings config.Settings) error {
	reader := bufio.NewReader(os.Stdin)
	// Prompt and read
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	pass, _ := gopass.GetPasswd()

	r := config.Registry{Username: strings.TrimRight(username, "\n"),
		Password: strings.TrimRight(string(pass), "\n")}
	settings.AddRegistry(strings.TrimRight(registry, "\n"), r)
	return settings.Save()
}
