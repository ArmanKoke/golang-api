package utils

import "net/http"

func AddHeaders(response http.ResponseWriter) http.ResponseWriter {
	response.Header().Set("Content-Type", "text/json; charset=utf-8")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	return response
}