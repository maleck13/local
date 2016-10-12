package local_test

import (
	"fmt"
	"testing"

	"github.com/maleck13/local/test"
)

func init() {
	if *test.IntegrationEnabled {
		fmt.Println("integration tests enabled")
	}
}

func TestRegister(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}
