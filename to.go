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

package to

import (
	"strconv"
	"fmt"
)

// String conversion
func String(value interface{}) string {
	result := ""

	if value != nil {
		result = fmt.Sprintf("%v", value)
	}

	return result
}

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

func Int(value interface{}) int64 {
	result := int64(0)
	if value != nil {
		result, _ = strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)
	}
	return result
}

func Float(value interface{}) float64 {
	result := float64(0)
	if value != nil {
		result, _ = strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	}
	return result
}

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
