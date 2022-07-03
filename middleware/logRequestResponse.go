package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mikarios/golib/logger"
)

var ErrHijackNotSupported = errors.New("hijack not supported")

type request struct {
	URI     string
	Method  string
	Headers http.Header
	BODY    string
}

type responseData struct {
	Body    string
	Status  int
	Headers http.Header
}

type loggingResponseWriter struct {
	http.ResponseWriter
	data *responseData
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.data.Body += string(b)

	return size, err
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.data.Status = statusCode
}

func (w *loggingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, ErrHijackNotSupported
	}

	return h.Hijack()
}

// LogRequestResponse can be used as a middleware in order to log the request as it comes towards the server,
// as well as the answer. ExcludedURIs can be used in order to not log specific urls such as login.
func LogRequestResponse(excludedURIS ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lw := &loggingResponseWriter{
				ResponseWriter: w,
				data:           &responseData{},
			}

			uri := r.RequestURI

			for _, exclude := range excludedURIS {
				if strings.Contains(uri, exclude) {
					next.ServeHTTP(lw, r)

					return
				}
			}

			start := time.Now()
			body, err := io.ReadAll(r.Body)
			if err != nil {
				panic(fmt.Errorf("could not read request body to log: %w", err))
			}

			_ = r.Body.Close()

			r.Body = io.NopCloser(bytes.NewBuffer(body))

			req := request{
				Headers: r.Header,
				URI:     uri,
				BODY:    string(body),
				Method:  r.Method,
			}
			reqBytes, _ := json.Marshal(req)
			reqStr := string(reqBytes)

			go logger.Debug(r.Context(), "New request:", reqStr)

			defer func() {
				lw.data.Headers = lw.Header()
				j, _ := json.Marshal(lw.data)
				go logger.Debug(
					r.Context(),
					"Request finished",
					"Request:", reqStr,
					"Response:", string(j),
					"Execution took:", time.Since(start),
				)
			}()

			next.ServeHTTP(lw, r)
		})
	}
}
