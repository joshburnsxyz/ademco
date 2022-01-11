# Ademco For Go

A lightweight golang library for dealing with Contact-ID and Ademco messaging formats.

## Installation

_install via go get_

```
$ go get github.com/joshburnsxyz/ademco
```

_and import into your code_

```go
import "github.com/joshburnsxyz/ademco"
```

## Usage Example

```go
import (
  "fmt"
  "github.com/joshburnsxyz/ademco"
)
func parseMsg(adm string) *ademco.Message {
  msg,err := ademco.NewMessage(adm)
  if err != nil {
    panic(err)
  }
  fmt.Printf("New message from %s: %v", msg.ClientCode, str) 
}
```
