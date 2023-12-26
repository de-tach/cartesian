# go-cartesian-product

[![Build Status](https://travis-ci.org/schwarmco/go-cartesian-product.svg?branch=master)](https://travis-ci.org/schwarmco/go-cartesian-product)
[![GoDoc](https://godoc.org/github.com/schwarmco/go-cartesian-product?status.svg)](https://godoc.org/github.com/schwarmco/go-cartesian-product)

a package for building [cartesian products](https://en.wikipedia.org/wiki/Cartesian_product) in golang

keep in mind, that because [how golang handles maps](https://blog.golang.org/go-maps-in-action#TOC_7.) your results will not be "in order"

## Installation

In order to start, `go get` this repository:

```
go get github.com/de-tach/cartesian
```

## Usage

```go
import (
    "fmt"
    "github.com/de-tach/cartesian"
)

func main() {

    a := []interface{}{1,2,3}
    b := []interface{}{"a","b","c"}

    c := cartesian.Iter(a, b)

    // receive products through channel
    for product := range c {
        fmt.Println(product)
    }

    // Unordered Output:
    // [1 c]
    // [2 c]
    // [3 c]
    // [1 a]
    // [1 b]
    // [2 a]
    // [2 b]
    // [3 a]
    // [3 b]
}
```

## Working with Types

The original code used the idea of using `interface{}` types which makes the code
less ideal when you know you only want to handle the same type of object. So for
example, if we know that if we have a set $X, Y \subset Z$ then we know that the
type space of this cartesian product is $X \cross Y$. For this problem the new
generics are quite ideal and should be easy to use without breaking the original
source API.
