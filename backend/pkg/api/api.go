package api

import (
	Object "Cll/pkg/object"
	"net/http"
)

func RouterAPI(cfg Object.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		switch r.Method {
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)

			return

		case "GET":
			w.WriteHeader(http.StatusOK)

			return

		case "POST":
			w.WriteHeader(http.StatusOK)

			return

		case "PUT":
			w.WriteHeader(http.StatusOK)

			return

		default:
			http.Error(w, "No such method", http.StatusMethodNotAllowed)
		}
	}
}
