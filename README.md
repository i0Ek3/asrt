# asrt

`asrt` helps you to do something you won't do, just check it all.

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

- [x] generic support
- [ ] error showing


## Contributing

PRs and Issues are also welcome.

## License

MIT.
