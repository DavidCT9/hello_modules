# Hello Modules

This is a repo to test Golang Modules

## Installation
Execute the following commandÑ
```bash
go get -u github.com/davidct9/Go_Modules
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/davidct9/hello_modules"
)

func main() {
	//message := hello_modules.Hello1("Loco")
	//fmt.Println(message)
	fmt.Printf(hello_modules.RandomHello(), "David")
}
```

