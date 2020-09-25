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

#### The classic "Hello" in Go.
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
#### Receive arguments by command-line prompt.
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
#### Finding duplicate lines
***
### 2. Program Structure
#### 2.1. Names
The names of Go functions, variables, constants, types, statement labels, and packages follow a 
simple rule: a name begins with a letter (that is, anything that Unicode deems a letter) or an
underscore and may have any number of addition al letters, digits, and underscores.

Go has 25 keywords:

break | default | func | interface | select
:---: | :---: | :---: | :---: | :---:
case | defer | go | map | struct
chan | else | goto | package | switch
const | fallthrough | if | range | type
continue | for | import | return | var

Constants:

true | false | iota | nil
--- | --- | --- | ---

Types:

int | int8 | int16 | int32 | int64 | -
:---: | :---: | :---: | :---: | :---: | :---:
uint | uint8 | uint16 | uint32 | uint64 | uintptr
float32 | float64| complex64 | complex128 
bool | byte | rune | string | error

Functions:

make | len | cap | new
:---: | :---: | :---: | :---:
append | copy | close | delete
complex | real | imag | panic
recover |

#### 2.2. Declarations

A declaration names a program entity and specifies some or all of its properties.

A Go program is stored in one or more files whose names end in .go.

```
package main

import "fmt"

const boilingF = 212.0

func main()  {
	var f = boilingF
	var c = (f -32) * 5 / 9
	fmt.Printf("Boling point = %g°F or %g°C\n", f, c)
}
```
```
$ go run boiling.go
```

This prints.

```
Boling point = 212°F or 100°C
```

The function fToC below encapsulates the temperature conversion logic so that it is 
defined only once but may be used from multiple places.

```
package main

import "fmt"

func main()  {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
```
```
$ go run ftoc.go
```

This prints.

```
32°F = 0°C
212°F = 100°C
```

#### 2.3. Variables

A *var* declaration creates a variable of a particular type, attaches a name to it, and sets its initial
value. Each declaration has the general form.

```
var name type = expression
```

The zero-value mechanism ensures that a variable always holds a well-defined value of its type;
in Go there is no such thing as an uninitialized variable.

```
var s string
```

Omitting the type allows declaration of multiple variables of different types:

```
var b, f, s = true, 2.3, "four" // bool, float64, string
```

A set of variables can also be initialize d by calling a function that returns multiple values:

```
var f, err = os.Open(name) // os.Open returns a file and an error
```

An alternate form called a short variable declaration may be used to declare and initialize local variables.

```
freq := rand.Float64() * 3.0
t := 0.0
f, err := os.Open(name)
```

If a variable is declared *var x int*, the expression &x ("address of x") yields a pointer to an
integer variable, that is, a value of type *int, which is pronounced "pointer to int." If this
value is called p, we say "p points to x," or equivalently "p contains the address of x." The variable
to which p points is written *p. The expression *p yields the value of that variable, an
int, but since *p denotes a variable, it may also appear on the left-hand side of an assignment,
in which case the assignment updates the variable.

```
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2 // equivalent to x = 2
fmt.Println(x) // "2"
```

Pointers are key to the *flag* package, which uses a program’s command-line arguments to set
the values of certain variables distributed throughout the program.

```
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```
```
$ go run echo4.go -s / ab df hg
$ go run echo4.go -n ab df hg
```

This prints.

```
ab/df/hg
ab df hg%
```

Another way to create a variable is to use the built-in function new. The expression new(T)
creates an unnamed variable of typ e T, initializes it to the zero value of T, and returns its
address, which is a value of type *T.

```
p := new(int) // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // "0"
*p = 2 // sets the unnamed int to 2
fmt.Println(*p) // "2"
```

The two newInt functions below have identical behaviors.

```
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```

The lifetime of a variable is the interval of time during which it exists as the program executes.
The lifetime of a package-level variable is the entire execution of the program.
A compiler may choose to allocate local variables on the heap or on the stack but, perhaps surprisingly,
this choice is not determined by whether *var or new* was used to declare the variable.

```
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}
```

Garbage collection is a tremendous help in writing correct programs, but it does not relieve
you of the burden of thinking about memory.

#### 2.4. Assignments

The value held by a variable is updated by an assignment statement, which in its simplest form
has a variable on the left of the = sign and an expression on the right.

```
x = 1 // named variable
*p = true // indirect variable
person.name = "bob" // struct field
count[x] = count[x] * scale // array or slice or map element
```

Numeric variables can also be incremented and decremented by ++ and -- statements:

```
v := 1
v++ // same as v = v + 1; v becomes 2
v-- // same as v = v - 1; v becomes 1 again
```

Another form of assignment, known as tuple assignment, allows several variables to be
assigned at once.

When computing the n-t h Fibonacci number iteratively:

```
func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}
```

Tuple assignment can also make a sequence of trivial assignments more compact:

```
i, j, k = 2, 3, 5
```

Appears in an assignment in which two results are expected, each produces an addition al boolean result:

```
v, ok = m[key] // map lookup
v, ok = x.(T) // type assertion
v, ok = <-ch // channel receive
```

As wit h variable declarations, we can assign unwanted values to the blank identifier:

```
_, err = io.Copy(dst, src) // discard byte count
_, ok = x.(T) // check type but discard result
```

This slice:

```
medals := []string{"gold", "silver", "bronze"}
```

If it had been written like this:

```
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```

#### 2.5. Type Declarations

Type declarations most often appear at package level, where the named type is visible throughout
the package, and if the name is exported (it starts with an upper-case letter), it’s accessible
from other packages as well.

```
package main

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

Conversions are also allowed between numeric types, and between string and some slice types,
as we will see in the next chapter. These conversions may change the representation of the
value.

#### 2.6. Packages and Files

Each package serves as a separate name space for its declarations. Within the image package,
for example, the identifier Decode refers to a different function than does the same identifier
in the unicode/utf16 package. To refer to a function from outside its package, we must
*qualify* the identifier to make explicit whether we mean image.Decode or utf16.Decode.

****
### 3. Basic Data Types
****
### 4. Composite Types