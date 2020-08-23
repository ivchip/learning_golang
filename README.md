# Learning Go
This repository I save my lessons learned

## Summary Book "The Go Programming Language"

### Preface
"Go is an open source programming language that makes it easy to build simple, reliable,
and efficient software." (From the [Golan (Go) web site](https://golang.org/)).

In Go:
* No implicit numeric conversions.
* No constructors or destructors.
* No operator overloading.
* No default parameters values.
* No inheritance.
* No generics.
* No exceptions.
* No macros.
* No function annotations.
* No tread-local storage.

### 1. Tutorial

The classic "Hello" in Go.
```
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
We will use as the command prompt throughout.
***
``` $ go run hello.go```

This prints.

``` Hello, 世界```
***
``` $ go build hello.go```

This creates an executable binary file called hello that can be run any time.

```
$ ./hello
Hello, 世界
```
***
Receive arguments by command-line prompt.
```
// EchoOne prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main()  {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Argument:", i)
	}
	fmt.Println("Concatenates the values:", s)
}
```
```
$ go run echoOne.go Hi, 🐍
```

These prints received arguments and concatenates strings.

```
Argument: 1
Argument: 2
Concatenates the values: Hi, 🐍
```
***
Syntax otherwise two.
```
// EchoTwo prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("Argument:", i)
	}
	fmt.Println("Concatenates the values:", s)
}
```
```
$ go run echoTwo.go Good bye 🍏
```

This prints.

```
Argument: 0
Argument: 1
Argument: 2
Concatenates the values: Good bye 🍏
```
***
Syntax otherwise three.
```
// EchoThree prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
```
```
$ go run echoThree.go Learn Go now
```

This prints.

```
Learn Go now
```
***
