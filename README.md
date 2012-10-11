```
PACKAGE

package to
    import "github.com/gosexy/to"

    A helper package for converting between datatypes.

    If a certain datatype could not be directly converted to another, the
    zero value of the destination type would be returned instead.

FUNCTIONS

func Bool(value interface{}) bool
    Converts the given value, if possible, into a bool. Returns false
    otherwise.

func Float64(value interface{}) float64
    Converts the given value, if possible, into a float64. Returns 0.0
    otherwise.

func Int(value interface{}) int
    Converts the given value, if possible, into an int. Returns 0 otherwise.

func Int64(value interface{}) int64
    Converts the given value, if possible, into an int64. Returns 0
    otherwise.

func List(value interface{}) []interface{}
    Converts the given value, if possible, into a []interface{}. Returns an
    empty []interface{} otherwise.

func String(value interface{}) string
    Converts the given value into a string.

func Uint64(value interface{}) uint64
    Converts the given value, if possible, into an uint64. Returns 0
    otherwise.

```

> Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/
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
