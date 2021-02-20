package controller

import (
	"io/ioutil"
	"net/http"
)

// GetFhir : 取得Fhir resource by id
func GetFhir(t, id string) string {
	var req string
	//req := "http://203.64.84.55:8080/fhir/" + t + "/" + id
	req = "https://hapi.fhir.tw/fhir/" + t + "/" + id

	r, _ := http.Get(req)
	data, _ := ioutil.ReadAll(r.Body)
	return string(data)
}

// GetFhirAll :
func GetFhirAll(t string) string {
	var req string
	//req := "http://203.64.84.55:8080/fhir/" + t
	req = "https://hapi.fhir.tw/fhir/" + t

	r, _ := http.Get(req)
	data, _ := ioutil.ReadAll(r.Body)
	return string(data)
}
