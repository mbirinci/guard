package guard_test

import (
  "github.com/mbirinci/guard"
  "testing"
)

type Sampling struct {
  Probability float64
}

//func TestGuard_Validate_ShouldFailedWhenNoRulesRegister(t *testing.T) {
//
//  g := guard.Guard[string]{}
//
//  _, err := g.Validate("test input")
//
//  if err == nil {
//    t.Error("no registered rules spec, err should not be nil")
//  }
//
//}

func TestGuard_Must_ShouldSuccess(t *testing.T) {

  s := Sampling{
    Probability: 0.03,
  }

  g := guard.Guard[Sampling]{}

  g.Must(func(input Sampling) bool {
    if input.Probability < 0 || input.Probability > 1 {
      return false
    }
    return true
  }, "Probability field of Sampling must be between 0 and 1")

  if _, ok := g.Validate(s); !ok {
    t.Error("guard success spec, result should be valid")
  }

}

//func TestGuard_Validate_ShouldFailed(t *testing.T) {
//
//  //expectations
//  s := Sampling{
//    Probability: 3,
//  }
//  failReason := "Probability field of Sampling must be between 0 and 1"
//
//  g := guard.Guard[Sampling]{}
//
//  g.Must(func(input Sampling) bool {
//    if input.Probability < 0 || input.Probability > 1 {
//      return false
//    }
//    return true
//  }, failReason)
//
//  result, err := g.Validate(s)
//
//  if err != nil {
//    t.Error("guard failed spec, err should be nil")
//  }
//
//  if result.IsValid {
//    t.Error("guard failed spec, guard result should not be valid")
//  }
//
//  if result.Messages[0] != failReason {
//    t.Errorf("guard failed spec, guard fail reason should be %s, but got %s", failReason, result.Messages[0])
//  }
//
//}

func TestGuard_PredefinedRules_Equal_ShouldSuccess(t *testing.T) {

  //expectations
  s := Sampling{
    Probability: 0.3,
  }
  failMessage := "Probability field of Sampling must be equal 0.3"

  g := guard.Guard[Sampling]{}

  g.Rules(func(s Sampling) []guard.PredefinedRule {
    return []guard.PredefinedRule{
      guard.Equal(s.Probability, 0.3, failMessage),
    }
  })

  if _, ok := g.Validate(s); !ok {
    t.Error("guard failed spec, validation should be valid")
  }

}
