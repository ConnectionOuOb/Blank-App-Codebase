package database

import (
	Object "Cll/pkg/object"

	"github.com/ledisdb/ledisdb/config"
	"github.com/ledisdb/ledisdb/ledis"
)

var LedisDB *ledis.DB

func Connect(cfg Object.Config) bool {
	dbCfg := config.NewConfigDefault()
	//dbCfg.DataDir = cfg.Dir

	l, _ := ledis.Open(dbCfg)
	db, err := l.Select(0)

	if err != nil {
		return false
	}

	LedisDB = db

	return true
}
