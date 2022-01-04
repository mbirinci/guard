package guard_test

import (
	"github.com/mbirinci/guard"
	"testing"
)

func TestValidator_ShouldFailedWhenNoRulesRegister(t *testing.T) {

	g := guard.Guard[string]{}

	_, err := g.Validate("test input")

	if err == nil {
		t.Error("no registered rules spec, err should not be nil")
	}

}
