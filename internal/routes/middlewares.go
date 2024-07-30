package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/iolave/go-logger"
)

func buildRequestLogMsg(status string, r *http.Request) string {
	method := strings.ToLower(r.Method)
	path := strings.ToLower(r.URL.Path)

	return fmt.Sprintf("api %s %s %s", method, path, status)
}

func requestLoggerMdw(log *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			customData := map[string]interface{}{
				"remoteAddr": r.RemoteAddr,
			}

			log.Info(buildRequestLogMsg("started", r), customData)
			defer func() {
				log.Info(buildRequestLogMsg("finished", r), customData)
			}()

			next.ServeHTTP(w, r)

		}
		return http.HandlerFunc(fn)
	}
}
