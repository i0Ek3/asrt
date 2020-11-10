# asrt

> Assertation for Go.

For now, this package is very simple, but just for now.

## Install

`go get "github.com/i0Ek3/asrt"`

## Usage

```Go
package main

import (
    "github.com/i0Ek3/asrt"
)

func main() {
    //...
    asrt.Asrt(t, got, want)
}
```

## API

- Asrt(t, got, want)
- Equal(t, got, want) bool
- NotEqual(t, got, want) bool


## License

MIT.
