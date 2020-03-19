package main

import (
"fmt"
"strings"
"net/http"
"io/ioutil"
)

func main() {

url := "130.195.10.173:8080/api/changeowner/CAR4"
method := "PUT"

payload := strings.NewReader("{\"owner\":\"XYZ\"}")

client := &http.Client {
}
req, err := http.NewRequest(method, url, payload)

if err != nil {
fmt.Println(err)
}
req.Header.Add("Content-Type", "application/json")

res, err := client.Do(req)
defer res.Body.Close()
body, err := ioutil.ReadAll(res.Body)

fmt.Println(string(body))
}
