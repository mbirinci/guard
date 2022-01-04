package guard

import "errors"

type Strategy uint8

const (
	Lazy Strategy = iota
	Eager
)

type Action[T any] func(input T) bool

type rule[T any] struct {
	action Action[T]
	reason string
}

type Result struct {
	IsValid bool
	Reasons []string
}

type Guard[T any] struct {
	Strategy Strategy
	rules    []rule[T]
}

func (v *Guard[T]) Rule(action Action[T], reason string) {
	v.rules = append(v.rules, rule[T]{
		action: action,
		reason: reason,
	})
}

func (v *Guard[T]) Validate(input T) (result Result, err error) {

	if len(v.rules) == 0 {
		return Result{}, errors.New("could not found any rules to validate")
	}

	result.IsValid = true

	for _, rule := range v.rules {
		if !rule.action(input) {
			result.IsValid = false
			result.Reasons = append(result.Reasons, rule.reason)
			if v.Strategy == Eager {
				return
			}
		}
	}

	return
}
