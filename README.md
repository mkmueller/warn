

# warn
`import "github.com/mkmueller/warn"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package warn implements methods to utilize multiple warnings




## <a name="pkg-index">Index</a>
* [type Warn](#Warn)
  * [func New(text string) Warn](#New)

#### <a name="pkg-examples">Examples</a>
* [New](#example_New)

#### <a name="pkg-files">Package files</a>
[warn.go](/src/github.com/mkmueller/warn/warn.go) 






## <a name="Warn">type</a> [Warn](/src/target/warn.go?s=143:278#L8)
``` go
type Warn interface {
    // Append another warning message
    Append(string)

    // Return a slice of warning messages
    Warnings() []string
}
```






### <a name="New">func</a> [New](/src/target/warn.go?s=354:380#L21)
``` go
func New(text string) Warn
```
New creates a new warning message









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
