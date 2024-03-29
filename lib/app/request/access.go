package request

import (
	"net/http"

	"sample/lib/app/response"
	"sample/lib/errs"
)

// AccessMethod проверка вертикальных прав на хендлеры по ролям
func AccessMethod(roles map[string][]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := response.New(r, w)
			us, err := rw.GetUser()
			if err != nil {
				rw.JSON(err)
				return
			}
			for i := range us.Roles {
				if rls, ok := roles[us.Roles[i]]; ok {
					for j := range rls {
						if r.Method == rls[j] {
							next.ServeHTTP(w, r)
							return
						}
					}
				}
			}
			rw.JSON(errs.NewForbidden(nil))
		})
	}
}

// Access проверка вертикальных прав на конкретный хендлер по ролям
func Access(roles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := response.New(r, w)
			us, err := rw.GetUser()
			if err != nil {
				rw.JSON(err)
				return
			}
			for i := range us.Roles {
				for j := range roles {
					if us.Roles[i] == roles[j] {
						next.ServeHTTP(w, r)
						return
					}
				}
			}
			rw.JSON(errs.NewForbidden(nil))
		})
	}
}
