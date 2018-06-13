
import (
	"net/http"

	"go.uber.org/zap"
)

func withLog(logger *zap.Logger) adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			d := newDelegator(w)
			next.ServeHTTP(d, r)
			logger.Info("request",
				zap.String("host", r.Host),
				zap.String("path", r.URL.Path),
				zap.Int("status", d.status),
				zap.String("request_id", requestID(r.Context())),
			)
		})
	}
}
