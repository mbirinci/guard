# guard

A validation package powered by generics

### Basic Usage

```go
package main

import "github.com/mbirinci/guard"

type Sampling struct {
  Probability float64
}

func main() {
  
  g := guard.Guard[Sampling]{
    Strategy: guard.Eager, // specify validation strategy
  }
  
  g.Must(func(s Sampling) bool {
    if s.Probability < 0 || s.Probability > 1 {
      return false
    }
    return true
  }, "Probability field of Sampling must be between 0 and 1")

  s := Sampling{
    Probability: 0.03,
  }
  
  if results, ok := g.Validate(s); !ok {
    // take your action when validation failed
    for _, message := range results {
      fmt.Println(message)
    }
  }

  // everything fine, keep going...
  
}

``` 

### Use Predefined Rules

```go
package main

import "github.com/mbirinci/guard"

type Sampling struct {
  Probability float64
}

func main() {
  
  g := guard.Guard[Sampling]{
    Strategy: guard.Eager, // specify validation strategy
  }
  
  g.Rules(func(s Sampling) guard.Rule[] {
    return guard.Rule[]{
      guard.Equal(s.Probability, 0.3, "Probability field of Sampling must be equal to 0.3"),
      guard.NotEmpty(s.Probability, "Probability field of Sampling must not be empty"),
    }
  })

  s := Sampling{
    Probability: 0.03,
  }

  if results, ok := g.Validate(s); !ok {
    // take your action when validation failed
    for _, message := range results {
      fmt.Println(message)
    }
  }

  // everything fine, keep going...

}

```