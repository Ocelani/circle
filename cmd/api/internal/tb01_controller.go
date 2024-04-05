package internal

import (
	"circle/pkg/logger"
	"circle/pkg/tb01"
	"encoding/json"
	"net/http"
)

// TB01Controller represents a controller for TB01.
type TB01Controller struct {
	Service tb01.Service
	log     *logger.APILogger
}

// NewTB01Controller creates a new TB01Controller.
func NewTB01Controller(service tb01.Service, log *logger.APILogger) *TB01Controller {
	return &TB01Controller{
		Service: service,
		log:     log,
	}
}

// postTB01 is a handler for the POST /tb01 route.
func (c *TB01Controller) Post(w http.ResponseWriter, r *http.Request) {
	l := c.log.Debug(r.Method, "/tb01", "request received")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		l.Debug(r.Method, "/tb01", "method not allowed")
		return
	}

	var data tb01.TB01
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		l.Err(r.Method, "/tb01", "failed to decode request body", http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if err := c.Service.Create(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		l.Err(r.Method, "/tb01", "insert on tb01 error", http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		l.Err(r.Method, "/tb01", "encode response data error", http.StatusCreated, err)
		return
	}

	if _, err := w.Write([]byte("\n")); err != nil {
		l.Err(r.Method, "/tb01", "write response error", http.StatusCreated, err)
	}

	l.Info(r.Method, "/tb01", "tb01 data inserted successfully")
}
