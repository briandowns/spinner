# Spinner

[![GoDoc](https://godoc.org/github.com/briandowns/spinner?status.svg)](https://godoc.org/github.com/briandowns/spinner)

For more detail about the library and its features, reference your local godoc once installed.

Contributions welcome!

## Installation

```bash
go get -u github.com/briandowns/spinner
```

## Available Character Sets

* ←↖↑↗→↘↓↙
* ▁▃▄▅▆▇█▇▆▅▄▃
* ▖▘▝▗
* ┤┘┴└├┌┬┐
* ◢◣◤◥
* ◰◳◲◱
* ◴◷◶◵
* ◐◓◑◒
* .oO@*
* |/-\
* ◡◡⊙⊙◠◠
* ⣾⣽⣻⢿⡿⣟⣯⣷
* >))'> >))'>  >))'>   >))'>    >))'>   <'((<  <'((< <'((<
* ⠁⠂⠄⡀⢀⠠⠐⠈
* ⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏
* abcdefghijklmnopqrstuvwxyz
* ▉▊▋▌▍▎▏▎▍▌▋▊▉
* ■□▪▫
* ←↑→↓
* ╫╪
* ⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏
* ⇐⇖⇑⇗⇒⇘⇓⇙
* ⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈
* ⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈
* ⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁
* ⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋

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

	s.UpdateCharSet(spinner.CharSets[8])                 // Update which spinner is used
	s.Start()

	s.UpdateDelay(1*time.Second)                         // Update the spinner speed
	s.Start()
}
```
