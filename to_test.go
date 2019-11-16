// Copyright (c) 2012-today José Nieto, https://xiam.dev
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package to

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "0", String(0))
	assert.Equal(t, "0", String(-0))
	assert.Equal(t, "-10", String(-10))
	assert.Equal(t, "9223372036854775807", String(int64(9223372036854775807)))
	assert.Equal(t, "-9223372036854775807", String(int64(-9223372036854775807)))
	assert.Equal(t, "4294967295", String(uint32(4294967295)))
	assert.Equal(t, "4294967295", String(uint(4294967295)))
	assert.Equal(t, "2147483647", String(int(2147483647)))
	assert.Equal(t, "2147483647", String(int32(2147483647)))
	assert.Equal(t, "18446744073709551615", String(uint64(18446744073709551615)))
}

func TestToBytes(t *testing.T) {
	assert.Equal(t, "0", String(Bytes(0)))
	assert.Equal(t, "0", String(Bytes(-0)))
	assert.Equal(t, "1", String(Bytes(1)))
	assert.Equal(t, "2", String(Bytes(2)))
	assert.Equal(t, "9223372036854775807", String(Bytes(9223372036854775807)))
	assert.Equal(t, "-9223372036854775807", String(Bytes(-9223372036854775807)))
}

func TestFloating(t *testing.T) {
	assert.Equal(t, "1.1", String(float32(1.1)))
	assert.Equal(t, "-1.1", String(float32(-1.1)))
	assert.Equal(t, "1.1234", String(float64(1.1234)))
	assert.Equal(t, "-1.1234", String(float64(-1.1234)))
}

func TestComplex(t *testing.T) {
	assert.Equal(t, "(1+1i)", String(complex(1, 1)))
	assert.Equal(t, "(1-1i)", String(complex(1, -1)))
	assert.Equal(t, "(-1-1i)", String(complex(-1, -1)))
}

func TestBoolean(t *testing.T) {
	assert.Equal(t, "true", String(true))
	assert.Equal(t, "false", String(false))
}

func TestNil(t *testing.T) {
	assert.Equal(t, "", String(""))
	assert.Equal(t, "", String(nil))
}

func TestBytes(t *testing.T) {
	assert.Equal(t, "hello", String([]byte("hello")))
}

func TestIntegers(t *testing.T) {
	assert.Equal(t, int64(1), Int64(1))
	assert.Equal(t, int64(-1), Int64(-1))
	assert.Equal(t, int64(234), Int64("234"))
	assert.Equal(t, int64(44), Int64(int8(44)))
	assert.Equal(t, int64(-234), Int64("-234"))
	assert.Equal(t, int64(-234), Int64("-0234"))
	assert.Equal(t, int64(-234), Int64([]byte("-0234")))
	assert.Equal(t, int(-234), Int([]byte("-0234")))
	assert.Equal(t, int(234), Int("0234"))
}

func TestFloat(t *testing.T) {
	assert.Equal(t, float64(1.123), Float64(1.123))
	assert.Equal(t, float64(-1.123), Float64("-01.12300"))
}

func TestBool(t *testing.T) {
	assert.Equal(t, true, Bool("t"))
	assert.Equal(t, false, Bool("f"))
	assert.Equal(t, true, Bool("TRUE"))
	assert.Equal(t, false, Bool("FALSE"))
	assert.Equal(t, true, Bool(1))
	assert.Equal(t, false, Bool(0))
	assert.Equal(t, true, Bool("1"))
	assert.Equal(t, false, Bool("0"))
}

func TestConvert(t *testing.T) {
	{
		value, err := Convert(1, reflect.Bool)
		assert.NoError(t, err)
		assert.Equal(t, true, value)
	}

	{
		value, err := Convert(int64(-456), reflect.String)
		assert.NoError(t, err)
		assert.Equal(t, "-456", value)
	}

	{
		value, err := Convert("0", reflect.Int64)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), value)
	}

	{
		value, err := Convert("my string", reflect.Slice)
		assert.NoError(t, err)
		assert.Equal(t, []rune("my string"), value)
	}
}

func TestTimeDuration(t *testing.T) {
	assert.Equal(t, time.Duration(123), Duration(123))
	assert.Equal(t, time.Second*12+time.Millisecond*32, Duration("12s32ms"))
	assert.Equal(t, time.Hour*13+time.Minute*37, Duration("13:37"))
	assert.Equal(t, -1*(time.Hour*13+time.Minute*37), Duration("-13:37"))
	assert.Equal(t, -1*(time.Hour*13+time.Minute*37+time.Second*21), Duration("-13:37:21"))
	assert.Equal(t, -1*(time.Hour*13+time.Minute*37+time.Second*21+time.Millisecond*344), Duration("-13:37:21.344"))
	assert.Equal(t, time.Hour*13+time.Minute*37+time.Second*21+time.Millisecond*344, Duration("13:37:21.344"))
	assert.Equal(t, time.Duration(0), Duration("A"))
}

func TestDate(t *testing.T) {
	assert.Equal(t, time.Date(2012, 3, 24, 0, 0, 0, 0, time.Local), Time("2012-03-24"))
	assert.Equal(t, time.Date(2012, 3, 24, 0, 0, 0, 0, time.Local), Time("2012/03/24"))
	assert.Equal(t, time.Date(2012, 3, 24, 23, 13, 37, 0, time.Local), Time("2012-03-24 23:13:37"))
	assert.Equal(t, time.Date(2012, 3, 24, 23, 13, 37, 123, time.Local), Time("2012-03-24 23:13:37.000000123"))
	assert.Equal(t, time.Date(2012, 3, 24, 23, 13, 37, 0, time.Local), Time("24/Mar/2012 23:13:37"))
	assert.Equal(t, time.Date(2012, 3, 24, 23, 13, 37, 123000000, time.UTC), Time("2012-03-24T23:13:37.123Z"))
	assert.Equal(t, time.Date(2012, 3, 24, 23, 13, 37, 123456789, time.UTC), Time("2012-03-24T23:13:37.123456789Z"))
}

func BenchmarkFmtIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", 1)
	}
}

func BenchmarkFmtFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%f", 1.1)
	}
}

func BenchmarkStrconvIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(1)
	}
}

func BenchmarkStrconvFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatFloat(1.1, 'f', 2, 64)
	}
}

func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(1)
	}
}

func BenchmarkFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(1.1)
	}
}

func BenchmarkIntToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(1)
	}
}

func BenchmarkBoolToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(true)
	}
}

func BenchmarkFloatToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes(1.1)
	}
}

func BenchmarkIntToBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool(1)
	}
}

func BenchmarkStringToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Time("2012-03-24")
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Convert("567", reflect.Int64)
		if err != nil {
			b.Fatalf("Test failed.")
			return
		}
	}
}
