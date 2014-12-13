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

## Example 1

The code below can also be found in the examples directory

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

	s.UpdateCharSet(spinner.CharSets[1])
	s.UpdateSpeed(200 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[6])
	s.UpdateSpeed(300 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[3])
	s.UpdateSpeed(400 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[8])
	s.UpdateSpeed(200 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)

	s.UpdateCharSet(spinner.CharSets[26])
	s.UpdateSpeed(100 * time.Millisecond)
	s.Restart()
	time.Sleep(4 * time.Second)
	s.Stop()
}
```

## Example 2

Provide your own spinner or send me an issue or pull request to add to the project.

```Go
package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	time.Sleep(4 * time.Second)
}
```

# Example 3

Generate a sequence of numbers

```Go

```
