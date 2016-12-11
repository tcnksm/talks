package main

import (
	"net/http"
	"os"
)

const EnvStatus = "API_STATUS"

const defaultMsg = "All systems operational"

func statusHandler(w http.ResponseWriter, r *http.Request) {
	statusMsg := os.Getenv(EnvStatus)
	if len(statusMsg) == 0 {
		statusMsg = defaultMsg
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(statusMsg + "\n"))
}
