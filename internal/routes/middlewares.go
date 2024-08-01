package routes

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/iolave/go-logger"
	"github.com/iolave/terraform-http-backend-service/internal/config"
)

func traceMdw() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			requestId := uuid.New().String()

			trace := config.Trace{
				RequestId: requestId,
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "trace", trace)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func buildRequestLogMsg(status string, r *http.Request) string {
	method := strings.ToLower(r.Method)
	path := strings.ToLower(r.URL.Path)

	return fmt.Sprintf("api %s %s %s", method, path, status)
}

func parseRemoteAddr(remoteAddr string) (ip string, port int, err error) {
	remoteAddrSplitted := strings.Split(remoteAddr, ":")

	port, err = strconv.Atoi(remoteAddrSplitted[len(remoteAddrSplitted)-1])

	if err != nil {
		return "", -1, err
	}

	ip = strings.Join(remoteAddrSplitted[:len(remoteAddrSplitted)-1], ":")

	return ip, port, nil

}

func requestLoggerMdw(log *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			customData := map[string]interface{}{}
			remoteIp, remotePort, err := parseRemoteAddr(r.RemoteAddr)
			if err != nil {
				customData["error"] = err.Error()
				log.Warn(buildRequestLogMsg("warn", r), customData)

			}
			customData["remoteIp"] = remoteIp
			customData["remotePort"] = remotePort
			customData["path"] = r.URL.Path

			ctx := r.Context()
			traceFromCtx := ctx.Value("trace")

			if reflect.TypeOf(traceFromCtx) != reflect.TypeOf(config.Trace{}) {
				customData["error"] = "unable to retrieve request trace"
				log.Warn(buildRequestLogMsg("warn", r), customData)
			} else {
				trace := traceFromCtx.(config.Trace)
				customData["trace"] = trace
			}

			log.Info(buildRequestLogMsg("started", r), customData)
			defer func() {
				log.Info(buildRequestLogMsg("finished", r), customData)
			}()

			next.ServeHTTP(w, r.WithContext(ctx))

		}
		return http.HandlerFunc(fn)
	}
}
