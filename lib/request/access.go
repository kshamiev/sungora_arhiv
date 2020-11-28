package request

import (
	"net/http"

	"sungora/lib/errs"
	"sungora/lib/response"
	"sungora/src/typ"
)

// AccessMethod проверка вертикальных прав на хендлеры по ролям
func AccessMethod(roles map[typ.Role][]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := response.New(r, w)
			us, err := rw.GetUser()
			if err != nil {
				rw.JSONError(err)
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
			rw.JSONError(errs.NewForbidden(nil))
		})
	}
}

// Access проверка вертикальных прав на конкретный хендлер по ролям
func Access(roles ...typ.Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := response.New(r, w)
			us, err := rw.GetUser()
			if err != nil {
				rw.JSONError(err)
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
			rw.JSONError(errs.NewForbidden(nil))
		})
	}
}
