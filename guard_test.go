package guard_test

import (
	"github.com/mbirinci/guard"
	"testing"
)

type Sampling struct {
	Probability float64
}

func TestGuard_Validate_ShouldFailedWhenNoRulesRegister(t *testing.T) {

	g := guard.Guard[string]{}

	_, err := g.Validate("test input")

	if err == nil {
		t.Error("no registered rules spec, err should not be nil")
	}

}

func TestGuard_Validate_ShouldSuccess(t *testing.T) {

	s := Sampling{
		Probability: 0.03,
	}

	g := guard.Guard[Sampling]{}

	g.Rule(func(input Sampling) bool {
		if input.Probability < 0 || input.Probability > 1 {
			return false
		}
		return true
	}, "Probability field of Sampling must be between 0 and 1")

	result, err := g.Validate(s)

	if err != nil {
		t.Error("validation success spec, err should be nil")
	}

	if !result.IsValid || len(result.Reasons) > 0 {
		t.Error("validation success spec, validation result should be valid")
	}

}

func TestGuard_Validate_ShouldFailed(t *testing.T) {

	//expectations
	s := Sampling{
		Probability: 3,
	}
	failReason := "Probability field of Sampling must be between 0 and 1"

	g := guard.Guard[Sampling]{}

	g.Rule(func(input Sampling) bool {
		if input.Probability < 0 || input.Probability > 1 {
			return false
		}
		return true
	}, failReason)

	result, err := g.Validate(s)

	if err != nil {
		t.Error("validation failed spec, err should be nil")
	}

	if result.IsValid {
		t.Error("validation failed spec, validation result should not be valid")
	}

	if result.Reasons[0] != failReason {
		t.Errorf("validation failed spec, validation fail reason should be %s, but got %s", failReason, result.Reasons[0])
	}

}