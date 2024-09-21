package tool

import (
	"os"
	"time"
)

func AutoDelete(filePath string, minute int) {
	go func() {
		time.AfterFunc(5*time.Minute, func() {
			os.Remove(filePath)
		})
	}()
}
