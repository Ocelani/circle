package main

import (
	"encoding/json"
	"net/http"
)

// postTB01 is a handler for the POST /tb01 route.
func postTB01(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	l := NewAPILogger(r.Method, "/tb01").Debug("request received")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		l.Debug("method not allowed")
		return
	}

	var tb01 TB01
	if err := json.NewDecoder(r.Body).Decode(&tb01); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		l.Err("failed to decode request body", http.StatusBadRequest, err)
		return
	}

	if err := tb01.Create(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		l.Err("failed to create tb01", http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	l.Info("tb01 data inserted successfully")
}
