package guard

type Strategy uint8

const (
  Lazy Strategy = iota
  Eager
)

type Action[T any] func(input T) bool

type rule[T any] struct {
  action  Action[T]
  message string
}

type PredefinedAction[T any] func(input T) []PredefinedRule

type PredefinedRule func() (bool, string)

type Guard[T any] struct {
  Strategy         Strategy
  rules            []rule[T]
  predefinedAction PredefinedAction[T]
}

func (v *Guard[T]) Must(action Action[T], message string) {
  v.rules = append(v.rules, rule[T]{
    action:  action,
    message: message,
  })
}

func (v *Guard[T]) Rules(action PredefinedAction[T]) {
  v.predefinedAction = action
}

func Equal[P comparable](field P, target P, message string) PredefinedRule {
  return func() (bool, string) {
    if field == target {
      return true, ""
    }
    return false, message
  }
}

func (v *Guard[T]) Validate(input T) (results []string, valid bool) {

  if len(v.rules) == 0 && v.predefinedAction == nil {
    return []string{"could not found any rules to validate"}, false
  }

  valid = true

  for _, rule := range v.rules {
    if !rule.action(input) {
      valid = false
      results = append(results, rule.message)
      if v.Strategy == Eager {
        return
      }
    }
  }

  if v.predefinedAction != nil {

    rules := v.predefinedAction(input)

    for _, rule := range rules {
      if ok, m := rule(); !ok {
        valid = false
        results = append(results, m)
        if v.Strategy == Eager {
          return
        }
      }
    }
  }

  return
}
