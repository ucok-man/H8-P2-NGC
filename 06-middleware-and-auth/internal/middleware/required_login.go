package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/jwt"
)

func (md *Middleware) RequiredLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			md.JSON.Error.InvalidAuthenticationToken(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			md.JSON.Error.InvalidAuthenticationToken(w, r)
			return
		}
		tokenstr := headerParts[1]

		var claim jwt.JWTClaim
		err := jwt.DecodeToken(tokenstr, &claim, md.Config.JWTSecreet)
		if err != nil {
			md.JSON.Error.InvalidAuthenticationToken(w, r)
			return
		}

		user, err := md.Entity.User.GetByID(claim.UserID)
		if err != nil {
			md.JSON.Error.InvalidAuthenticationToken(w, r)
			return
		}

		// set context
		r = r.WithContext(context.WithValue(r.Context(), md.Config.GetUserCtxKey(), user))

		next.ServeHTTP(w, r)

	})
}
