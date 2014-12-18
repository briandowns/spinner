# Spinner

Package spinner is a simple package to add a spinner to an application.

[![GoDoc](https://godoc.org/github.com/briandowns/spinner?status.svg)](https://godoc.org/github.com/briandowns/spinner)

For more detail about the library and its features, reference your local godoc once installed.

Contributions welcome!

## Installation

```bash
go get github.com/briandowns/spinner
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

## Examples

```Go
package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[10], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
	s.Stop()
}
```

## Update the spinner and restart the spinner

```Go
	s.UpdateCharSet(spinner.CharSets[1])  // Update spinner to use a different character set
	s.Restart()                           // Restart the spinner
	time.Sleep(4 * time.Second)
	s.Stop()
```

## Update spin speed and restart the spinner

```Go
	s.UpdateSpeed(200 * time.Millisecond) // Update the speed the spinner spins at
	s.Restart()
	time.Sleep(4 * time.Second)
	s.Stop()
```

## Provide your own spinner

(or send me an issue or pull request to add to the project)

```Go
package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	someSet := []string{"+", "-"}
	s := spinner.New(someSet, 100*time.Millisecond)
	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()
}
```

## Generate a sequence of numbers

```Go
package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	setOfDigits := spinner.GenerateNumberSequence(25)    // Generate a 25 digit string of numbers
	s := spinner.New(setOfDigits, 100*time.Millisecond)
	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()
}
```
