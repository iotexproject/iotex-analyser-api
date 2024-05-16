package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/iotexproject/iotex-antenna-go/v2/jwt"
	graphqlruntime "github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

var (
	whitelistAPI = []string{
		"api.StakingService.VoteByHeight",
		"api.AccountService.Erc20TokenBalanceByHeight",
	}
)

var JWTTokenMiddleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		for _, a := range whitelistAPI {
			if strings.EqualFold(a, api) {
				ctx := context.WithValue(r.Context(), WhitelistCtxKey, true)
				h.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		// Get token from authorization header.
		jwtString := ""
		bearer := r.Header.Get("Authorization")
		if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
			jwtString = bearer[7:]
		}

		if jwtString == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		jwtoken, err := jwt.VerifyJWT(jwtString)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		claim := &Claims{JWT: jwtoken}
		if err := claim.CheckPermisson(); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), TokenCtxKey, claim)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// JWTTokenValid operation midware
func JWTTokenValid() graphqlruntime.MiddlewareFunc {
	return graphqlruntime.MiddlewareFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error) {
		if r.Method == "OPTIONS" || (r.Method == "GET" && r.URL.Path == "/graphql") {
			return ctx, nil
		}
		// Get token from authorization header.
		jwtString := ""
		bearer := r.Header.Get("Authorization")
		if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
			jwtString = bearer[7:]
		}

		if jwtString == "" {
			return ctx, graphqlruntime.NewMiddlewareError("unauthorized", "authorization header is required")
		}

		jwtoken, err := jwt.VerifyJWT(jwtString)
		if err != nil {
			return ctx, graphqlruntime.NewMiddlewareError("unauthorized", "Invalid token")
		}
		claim := &Claims{JWT: jwtoken}
		if err := claim.CheckPermisson(); err != nil {
			return ctx, graphqlruntime.NewMiddlewareError("unauthorized", err.Error())
		}
		return context.WithValue(r.Context(), TokenCtxKey, &Claims{JWT: jwtoken}), nil
	})
}
