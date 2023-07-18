# must

This Go package provides a couple of utility functions for the common
case of panicking on an encountered error.

Many Go packages (including those in the standard library) provide
specialized functions that do exactly this; e.g. the `html/template`
package provides a [`Must`](https://pkg.go.dev/html/template#Must)
function, that can be used like this:

```go
var t = template.Must(template.New("name").Parse("html"))
```

With the introduction of generics in Go 1.18, packages no longer need
their own specialized panic functions; you can use the function from
this package (or its brethren, `Assert`, and `Must2`).

```go
import "html/template"
import . "github.com/rollcat/etc/must"

var t = Must(template.New("name").Parse("html"))

```

If you ever need `Must3` or higher-arity functions, I'll happily add
them.
