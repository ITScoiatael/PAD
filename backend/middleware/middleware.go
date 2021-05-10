package middleware

import (
	"context"
	"net/http"

	"graphql-server/jwt"
	"graphql-server/prisma/db"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(prisma db.PrismaClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			adminLogin, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			admin, err := prisma.Admin.FindFirst(
				db.Admin.Login.Equals(adminLogin),
			).Exec(r.Context())
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &admin)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *db.AdminModel {
	raw, _ := ctx.Value(userCtxKey).(*db.AdminModel)
	return raw
}
