package merchandiser_utils

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

func Base64ByUrl(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	sEnc := base64.StdEncoding.EncodeToString(data)

	return sEnc
}
