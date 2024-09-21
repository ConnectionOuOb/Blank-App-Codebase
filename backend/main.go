package main

import (
	"{}/pkg/api"
	"{}/pkg/config"
	"{}/pkg/object"
)

var CFG object.Config

func main() {
	api.CreateRouter(config.LoadConfig())
}
