/*
 * Sample API
 *
 * Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
 *
 * API version: 0.1.9
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */


package main

import (
	"log"
	"net/http"

	openapi "github.com/bjne/test-openapi-golang/go"
)

func main() {
	log.Printf("Server started")

	CustomApiService := openapi.NewCustomApiService()
	DefaultApiController := openapi.NewDefaultApiController(CustomApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
