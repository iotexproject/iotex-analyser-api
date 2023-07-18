package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/httprate"
	"github.com/iotexproject/go-pkgs/crypto"
)

var (
	whitelistID = []string{
		"io1k768lhjwwf89qd4fqt6gpheqvzqfz7mxept4tp",
	}
)

var CheckWhiteListMiddleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed, ok := r.Context().Value(WhitelistCtxKey).(bool)
		if ok && allowed {
			h.ServeHTTP(w, r)
			return
		}
		claims, ok := r.Context().Value(TokenCtxKey).(*Claims)
		if !ok {
			http.Error(w, "failed to get claims in context", http.StatusUnauthorized)
			return
		}
		trustor, err := crypto.HexStringToPublicKey(claims.Issuer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		for _, id := range whitelistID {
			if strings.EqualFold(id, trustor.Address().String()) {
				h.ServeHTTP(w, r)
				return
			}
		}
		RateLimitMiddleware(h).ServeHTTP(w, r)
	})

}
var RateLimitMiddleware = httprate.Limit(
	5,             // requests
	1*time.Minute, // per duration
	httprate.WithKeyFuncs(func(r *http.Request) (string, error) {
		return r.Header.Get("Authorization"), nil
	}))
