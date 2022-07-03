package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mikarios/golib/logger"
)

// ErrRecover is returned when we recover from an error.
var ErrRecover = errors.New("caught panic stacktrace")

// RecoverPanic is operating as middleware to handle any panic that may occur.
func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		defer func() {
			if err := recover(); err != nil {
				logger.Error(ctx, fmt.Errorf("%w: %v", ErrRecover, err), "middleware recovering from panic error")
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
