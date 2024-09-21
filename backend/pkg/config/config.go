package config

import (
	Object "Cll/pkg/object"
)

func CreateConfig() Object.Config {
	return Object.Config{
		Api:  "/api",
		Dir:  "./data",
		Port: "8080",
	}
}
