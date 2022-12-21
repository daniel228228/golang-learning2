package http_handler

import (
	"context"
	"errors"
	"net/http"

	"http_sample/internal/config"
	"http_sample/internal/logger"
	"http_sample/internal/service"

	"github.com/gorilla/mux"
)

type HttpHandler interface{}

type httpHandler struct {
	config  *config.Config
	log     logger.Logger
	service service.Service
	router  *mux.Router
	server  *http.Server
}

func NewHttpHandler(config *config.Config, log logger.Logger, service service.Service) *httpHandler {
	return &httpHandler{
		config:  config,
		log:     log,
		service: service,
	}
}

func (h *httpHandler) Serve(ctx context.Context) error {
	h.log.Print("serving HTTP Handler")
	defer h.log.Print("finished serving HTTP Handler")

	h.router = mux.NewRouter()
	h.router.StrictSlash(true)
	h.router.HandleFunc("/get", h.Get).Methods("GET")
	h.router.HandleFunc("/write", h.Write).Methods("GET")

	h.server = &http.Server{
		Addr:    ":" + h.config.ServicePort,
		Handler: h.router,
	}

	if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (h *httpHandler) Shutdown(ctx context.Context) error {
	h.log.Print("stopping HTTP Handler")
	defer h.log.Print("stopped HTTP Handler")

	return h.server.Shutdown(ctx)
}

func (h *httpHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.log.Printf("new request: url: %s, method: %s", r.URL.String(), r.Method)

	if resp, err := h.service.Get(0); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("error"))
	} else {
		w.Write(append([]byte("response: "), resp...))
	}
}

func (h *httpHandler) Write(w http.ResponseWriter, r *http.Request) {
	h.log.Printf("new request: url: %s, method: %s", r.URL.String(), r.Method)

	if err := h.service.Write(0, ""); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("error"))
	} else {
		w.Write([]byte("ok"))
	}
}
