// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package auth

import (
	"testing"

	"github.com/iotexproject/iotex-antenna-go/v2/jwt"
	"github.com/stretchr/testify/require"
)

func TestClaim(t *testing.T) {
	r := require.New(t)

	c := Claims{JWT: &jwt.JWT{}}

	c.Scope = "Read: wNzY2OTI0OSwia"
	r.True(c.AllowRead())
	c.Subject = "AnalyserAPI: NiIsInR5cCI6Ikp"
	r.True(c.IsAnalyserAPI())
}
