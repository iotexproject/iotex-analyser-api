// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/iotexproject/iotex-antenna-go/v2/jwt"
)

// const
const (
	AnalyserAPI = "AnalyserAPI"
)

// Context keys
var (
	TokenCtxKey     = &contextKey{"Token"}
	WhitelistCtxKey = &contextKey{"Whitelist"}
	ErrorCtxKey     = &contextKey{"Error"}
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenAllow   = errors.New("token not allowed")
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "auth context value " + k.name
}

type Claims struct {
	*jwt.JWT
}

func (c *Claims) AllowRead() bool {
	return strings.Contains(c.Scope, jwt.READ)
}

func (c *Claims) IsAnalyserAPI() bool {
	return strings.Contains(c.Subject, AnalyserAPI)
}

func (c *Claims) CheckPermisson() error {
	if !c.AllowRead() || !c.IsAnalyserAPI() {
		return ErrTokenAllow
	}
	//check if the token is expired
	if c.ExpiresAt < time.Now().Unix() {
		return ErrTokenExpired
	}
	return nil
}
