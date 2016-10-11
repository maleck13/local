package local

import (
	"fmt"

	"github.com/maleck13/local/test"
)

func init() {
	if *test.IntegrationEnabled {
		fmt.Println("integration tests enabled")
	}
}
