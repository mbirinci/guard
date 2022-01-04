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
  
  g.Rule(func(input Sampling) bool {
    if input.Probability < 0 || input.Probability > 1 {
      return false
    }
    return true
  }, "Probability field of Sampling must be between 0 and 1")

  s := Sampling{
    Probability: 0.03,
  }
  
  result, err := g.Validate(s)
  
  if err != nil {
    panic(err)
  }
  
  if !result.IsValid {
    fmt.Println("validation failed")
  }
  
}

``` 