package utils

import (
	"log"
	"net/http"
	"strconv"
)

//should handle case when param is empty
func ProcessRawQuery(r *http.Request, fieldName string) string {
	return r.URL.Query()[fieldName][0]
}

func ConvertStringId(id string) uint32 {
	converted := parseUint(id, 10, 32)
	return uint32(converted)
}

//for now id is not incremented but better make it so
func ConvertId(r *http.Request) uint32 {
	id := r.FormValue("id")
	return ConvertStringId(id)
}

func ConvertPrice(r *http.Request) uint32 {
	price := r.FormValue("price")
	converted := parseUint(price, 10, 32)
	return uint32(converted)
}

func ConvertStatus(r *http.Request) uint8 {
	status := r.FormValue("status")
	converted := parseUint(status, 10, 8)
	return uint8(converted)
}

func ConvertMileage(r *http.Request) uint32 {
	mileage := r.FormValue("mileage")
	converted := parseUint(mileage, 10, 32)
	return uint32(converted)
}

func parseUint(value string, base int, bitSize int) uint64 {
	integer, err := strconv.ParseUint(value, base, bitSize)
	if err != nil {
		log.Printf("Error converting to uint for value %v: %s", value, err)
	}
	return integer
}
