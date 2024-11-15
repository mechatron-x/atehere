package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           bytes.NewBuffer([]byte{}),
	}
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *logResponseWriter) Write(b []byte) (int, error) {
	_, err := lrw.body.Write(b)
	if err != nil {
		return 0, err
	}

	return lrw.ResponseWriter.Write(b)
}

func (lrw *logResponseWriter) Header() http.Header {
	return lrw.ResponseWriter.Header()
}

func Logger(next http.Handler, log *zap.Logger) http.Handler {
	traceId := uuid.NewString()

	f := func(w http.ResponseWriter, r *http.Request) {
		buf, _ := io.ReadAll(r.Body)

		rdr := io.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr

		lr := io.NopCloser(bytes.NewBuffer(buf))
		lBuf, _ := io.ReadAll(lr)

		log.Info("HTTP Request",
			zap.String("trace_id", traceId),
			zap.String("protocol", r.Proto),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_address", r.RemoteAddr),
			zap.String("header", headerToString(r.Header.Clone())),
			zap.String("body", string(lBuf)),
		)

		lrw := newLogResponseWriter(w)
		next.ServeHTTP(lrw, r)

		log.Info("HTTP Response",
			zap.String("trace_id", traceId),
			zap.String("protocol", r.Proto),
			zap.Int("status", lrw.statusCode),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_address", r.RemoteAddr),
			zap.String("header", headerToString(lrw.Header().Clone())),
			zap.String("body", lrw.body.String()),
		)
	}

	return http.HandlerFunc(f)
}

func headerToString(header http.Header) string {
	h := strings.Builder{}
	if header == nil {
		return h.String()
	}

	for key, values := range header {
		h.WriteString(fmt.Sprintf("%s:[%s];", key, strings.Join(values, ", ")))
	}

	return h.String()
}
