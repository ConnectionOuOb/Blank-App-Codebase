package api

import (
	Object "Cll/pkg/object"
	"log"
	"net/http"
)

func CreateRouter(cfg Object.Config) {
	http.HandleFunc(cfg.Api, RouterAPI(cfg))

	err := http.ListenAndServe(":"+cfg.Port, nil)
	log.Fatal(err)
}
