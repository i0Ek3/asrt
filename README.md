# asrt

> Assertation for Go.

For now, this package is very simple, but just for now.

## Getting Started

### Install

`go get "github.com/i0Ek3/asrt"`

### Usage

```Go
package main

import (
    "github.com/i0Ek3/asrt"
)

func main() {
    //...
    asrt.Asrt(t, got, want)
    asrt.Equal(got, want)
    asrt.NotEqual(got, want)
    asrt.NotNil(value)
    asrt.Neg(prevalue, neg)
}
```

## TODO

- error showing


## Contributing

PRs and Issues are also welcome.

## License

MIT.
