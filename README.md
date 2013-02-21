# gosexy/to

*Convenient* functions for converting between common Go datatypes.

This is an *experimental package* it avoids error handling and reporting, if
any conversion cannot be done, a zero value will be always returned.

Life is too short for properly catching trivial conversion errors.

## Installing

```sh
go get github.com/gosexy/to
```

## Usage

Import the package

```go
import "github.com/gosexy/to"
```

Convert something

```go
i, err := to.Convert("567", reflect.Int64)
// 567, nil

a := to.String(1)
// "1"
b := to.String(1.1)
// "1.1"
c := to.String(true)
// "true"

a := to.Int("1")
// 1
b := to.Int("-1")
// -1
b := to.Int("")
// 0

a := to.Float32("1.4")
// 1.4
b := to.Float32("A")
// 0.0

a := to.Bool("true")
// true
b := to.Bool("f")
// false
```

## License

This is Open Source released under the terms of the MIT License:

> Copyright (c) 2013 José Carlos Nieto, http://xiam.menteslibres.org/
>
> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
