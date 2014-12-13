# Spinner

[![GoDoc](https://godoc.org/github.com/briandowns/spinner?status.svg)](https://godoc.org/github.com/briandowns/spinner)

For more detail about the library and its features, reference your local godoc once installed.

Contributions welcome!

## Installation

```bash
go get -u github.com/briandowns/spinner
```

## Example

The code below can also be found in the examples directory

```Go
package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := New(spinner.CharSets[10], 100*time.Millisecond) // Build our new spinner
	s.Start()                                            // Start the spinner
	time.Sleep(5 * time.Second)                          // Run for some time to simulate work
	s.Stop()                                             // Stop the spinner
}
```
