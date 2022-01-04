package validation_test

import (
	"testing"
	"validation"
)

func TestValidator_ShouldFailedWhenNoRulesRegister(t *testing.T) {

	pv := validation.Validator[string]{}

	_, err := pv.Validate("test input")

	if err == nil {
		t.Error("no registered rules spec, err should not be nil")
	}

}
