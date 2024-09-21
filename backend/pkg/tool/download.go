package tool

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, saveTo string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "Error while downloading file"
	}

	out, err := os.Create(saveTo)
	if err != nil {
		return err.Error()
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err.Error()
	}

	return ""
}
