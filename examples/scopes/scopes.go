package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Retrieve a list of scopes for which this user has access.
// GET /scopes

func Retrievealistofscopesforwhichthisuserhasaccess.() {
  apiKey := os.Getenv("SENDGRID_APIKEY")
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/scopes", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

