package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
)

func IsTaskOwner() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(controllers.UserKey).(domain.User)
			task := r.Context().Value(controllers.TaskKey).(domain.Task)

			if task.UserId != user.Id {
				log.Printf("IsTaskOwner: user %d attempted to access task %d owned by user %d", user.Id, task.Id, task.UserId)
				controllers.Forbidden(w, errors.New("access denied"))
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
