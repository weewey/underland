package handlers

import "net/http"

func IncrementSocialIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
