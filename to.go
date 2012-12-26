/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
	A helper package for converting between datatypes.

	If a certain datatype could not be directly converted to another, the
	zero value of the destination type would be returned instead.
*/
package to

import (
	"fmt"
	"github.com/gosexy/sugar"
	"strconv"
)

/* Converts the given value into a string. */
func String(value interface{}) string {
	result := ""

	if value != nil {
		result = fmt.Sprintf("%v", value)
	}

	return result
}

/* Converts the given value, if possible, into a []interface{}. Returns an empty []interface{} otherwise. */
func List(value interface{}) []interface{} {
	list := []interface{}{}

	if value != nil {
		switch value.(type) {
		case []interface{}:
			for _, v := range value.([]interface{}) {
				list = append(list, v)
			}
		}
	}

	return list
}

/* Converts the given value, if possible, into a map[string]interface{}. Returns nil otherwise. */
/*
func Map(value interface{}) map[string]interface{} {

	if value != nil {
		switch value.(type) {
		case map[string]interface{}:
			return value.(map[string]interface{})
		case sugar.Map:
			mapped := map[string]interface{}{}
			for k, v := range value.(sugar.Map) {
				mapped[k] = v
			}
			return mapped
		}
	}

	return nil
}
*/

/* Converts the given value, if possible, into a sugar.Map. Returns nil otherwise. */
func Map(value interface{}) sugar.Map {

	if value != nil {
		switch value.(type) {
		case map[string]interface{}:
			mapped := sugar.Map{}
			for k, v := range value.(map[string]interface{}) {
				mapped[k] = v
			}
			return mapped
		case sugar.Map:
			return value.(sugar.Map)
		}
	}

	return nil
}

/* Converts the given value, if possible, into an int. Returns 0 otherwise. */
func Int(value interface{}) int {
	result := int(0)
	if value != nil {
		result, _ = strconv.Atoi(String(value))
	}
	return result
}

/* Converts the given value, if possible, into an int64. Returns 0 otherwise. */
func Int64(value interface{}) int64 {
	result := int64(0)
	if value != nil {
		result, _ = strconv.ParseInt(String(value), 10, 64)
	}
	return result
}

/* Converts the given value, if possible, into an uint64. Returns 0 otherwise. */
func Uint64(value interface{}) uint64 {
	result := uint64(0)
	if value != nil {
		result, _ = strconv.ParseUint(String(value), 10, 64)
	}
	return result
}

/* Converts the given value, if possible, into a float64. Returns 0.0 otherwise. */
func Float64(value interface{}) float64 {
	result := float64(0)
	if value != nil {
		result, _ = strconv.ParseFloat(String(value), 64)
	}
	return result
}

/* Converts the given value, if possible, into a bool. Returns false otherwise. */
func Bool(value interface{}) bool {
	result := true

	if value == nil {
		result = false
	} else {
		switch String(value) {
		case "", "0", "false":
			result = false
		}
	}

	return result

}
