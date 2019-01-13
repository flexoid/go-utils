package logging_test

import (
	"net/http"

	"github.com/flexoid/go-utils/logging"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func ExampleRequestLogger() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Middleware chain
	http.Handle("/",
		hlog.NewHandler(log.Logger)(
			hlog.RequestIDHandler("req_id", "Request-Id")(
				logging.RequestLogger()(
					handler))))
	http.ListenAndServe(":4000", nil)
}
