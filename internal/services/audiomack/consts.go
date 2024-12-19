package services

import (
	"net/http"
	"os"
)

func genApiHeader() http.Header {
	var appId = os.Getenv("AUDIOMACK_APPLICATION_ID")

	return http.Header{
		"Content-Type":   {"application/json"},
		"Application-Id": {appId}}
}
