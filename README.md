# asrt

> Assertation for Go.

`asrt` help you to do something you won't do, and laterly, we will support generic for `asrt`.

## Getting Started

### Install

`go get -u "github.com/i0Ek3/asrt"`

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
