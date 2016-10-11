package test

import "flag"
import "github.com/maleck13/local/config"
import "os"
import "path/filepath"

var (
	//IntegrationEnabled is flag for running the integration tests
	IntegrationEnabled = flag.Bool("integration", false, "enabled integration tests")
)

func init() {
	path := filepath.Join(os.Getenv("GOPATH"), "src/github.com/maleck13/local/config/config-local.json")
	if _, err := config.LoadConfig(path); err != nil {
		panic(err)
	}
}
