# gob64

Marshal/Unmarshal Base64 Encoded Gob To String.

## Usage

Marshal a struct:
```go
package main

import (
	"fmt"

	"github.com/moizalicious/gob64"
)

type data struct {
	Foo string
	Bar string
}

func main() {
	d := data{
		Foo: "ABC",
		Bar: "XYZ",
	}

	s, err := gob64.Marshal(d)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
```

The above will print the following to the console:
```
Iv+BAwEBBGRhdGEB/4IAAQIBA0ZvbwEMAAEDQmFyAQwAAAAN/4IBA0FCQwEDWFlaAA==
```

Unmarshal an encoded string:
```go
package main

import (
	"fmt"

	"github.com/moizalicious/gob64"
)

type data struct {
	Foo string
	Bar string
}

func main() {
	s := "Iv+BAwEBBGRhdGEB/4IAAQIBA0ZvbwEMAAEDQmFyAQwAAAAN/4IBA0FCQwEDWFlaAA=="

	var d data
	if err := gob64.Unmarshal(s, &d); err != nil {
		panic(err)
	}

	fmt.Println(d)
}
```

The above will print the following to the console:
```
{ABC XYZ}
```
