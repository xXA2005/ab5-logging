# ab5 logging package

## installation

run `go get github.com/xXA2005/ab5-logging` to install the package

## usage

### from (255,1,1) to (1,255,1)

```go
package main

import (
	"fmt"

	ab5 "github.com/xXA2005/ab5-logging"
)

func main() {
	ab5.Error(fmt.Errorf("error"))
	ab5.Success("success")
	ab5.Log("log")
	ab5.Warn("warning")
	fmt.Println()
	fmt.Println(ab5.Horizontal(255, 1, 1, 1, 255, 1, "AAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAA\nAAAAAAAA"), ab5.Reset)
	fmt.Println()
	fmt.Println(ab5.Vertical(255, 1, 1, 1, 255, 1, "AAAAAAAAA\nAAAAAAAAA\nAAAAAAAA\nFRRRRRRRRRR\nSSSSSSSSSSSS"), ab5.Reset)

}

```
