package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           nil,
	}
}

func (rcv *logResponseWriter) WriteHeader(code int) {
	rcv.statusCode = code
	rcv.ResponseWriter.WriteHeader(code)
}

func (rcv *logResponseWriter) Write(b []byte) (int, error) {
	rcv.body = make([]byte, len(b))
	copy(rcv.body, b)

	return rcv.ResponseWriter.Write(b)
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
			zap.String("body", string(lrw.body)),
		)
	}

	return http.HandlerFunc(f)
}
