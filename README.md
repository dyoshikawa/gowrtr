gowrtr
==

gowrtr is a library that supports golang code generating.

This library is inspired by [square/javapoet](https://github.com/square/javapoet).

Synopsis
--

Here is a simple example:

```go
package main

import (
	"fmt"

	"github.com/moznion/gowrtr/generator"
)

func main() {
	generator := generator.NewRootGenerator(
		generator.NewCommentGenerator(" THIS CODE WAS AUTO GENERATED"),
		generator.NewPackageGenerator("main"),
		generator.NewNewlineGenerator(),
	).AddStatements(
		generator.NewFuncGenerator(
			nil,
			generator.NewFuncSignatureGenerator("main"),
		).AddStatements(
			generator.NewRawStatementGenerator(`fmt.Println("hello, world!")`),
		),
	).
		EnableGofmt("-s").
		EnableGoimports()

	generated, err := generator.Generate(0)
	if err != nil {
		panic(err)
	}
	fmt.Println(generated)
}
```

then it generates the golang's code like so:

```go
// THIS CODE WAS AUTO GENERATED
package main

import "fmt"

func main() {
        fmt.Println("hello, world!")
}
```

Documents
--

Please refer to the godoc: [![GoDoc](https://godoc.org/github.com/moznion/gowrtr?status.svg)](https://godoc.org/github.com/moznion/gowrtr)

- `RootGenerator` is a entry point to generate the go code.
- `RootGenerator` supports following code formatting on code generating phase. It applies such formatters to generated code.
  - `gofmt`: with `EnableGofmt(gofmtOptions ...string)`
  - `goimports`: with `EnableGoimports()`

License
--

```
The MIT License (MIT)
Copyright © 2019 moznion, http://moznion.net/ <moznion@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

