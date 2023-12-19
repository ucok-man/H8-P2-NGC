package middleware

import (
	"fmt"
	"net/http"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/entity"
)

func (md *Middleware) RequiredSuperUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj := r.Context().Value(md.Config.GetUserCtxKey())
		user, ok := obj.(*entity.User)
		if !ok {
			md.JSON.Error.InternalServer(w, fmt.Errorf("user must be exist at this stage"))
			return
		}

		if user.Role != entity.RoleSuperAdmin {
			md.JSON.Error.NotPermittedResource(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
