package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/mikarios/golib/logger"
)

// TransactionID is operating as middleware to set a transaction ID in request header and context. If a header with
// a transaction ID does not exist, generates a new one.
func TransactionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		txKey := string(logger.Settings.TransactionKey)

		txID := r.Header.Get(txKey)
		if txID == "" {
			txID = uuid.New().String()
			r.Header.Add(txKey, txID)
		}

		ctx := context.WithValue(r.Context(), logger.Settings.TransactionKey, txID)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
