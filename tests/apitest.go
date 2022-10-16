package tests

import (
	"fmt"
	"log"
	"net/http"
)

const (
	APIURL = "localhost:5000/api/v1"
)

type Test struct {
}

var apiTest = make(map[int]string)

func GenerateTests(n int) {
	for i := 0; i < n; i++ {
		//reqStr := APIURL + randMethod + methodParams // Test Request
		resp, err := http.Get("") // reqStr
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp.Body)

	}
}
