package tests

import (
	//"encoding/json"
	"fmt"
	"gin-ticker/models"
	"log"
	"math/rand"
	"net/http"
)

const (
	APIURL = "localhost:5000/api/v1"
)

func GenerateTests(n int) {

	apiTest := make(map[int]string)
	apiTest = map[int]string{
		0: "rates",
		1: "rates",
		2: "rates",
	}

	// requestBody, err := json.Marshal(map[string]string{
	// 	"Authorization": "Bearer cd25",
	// })

	// if err != nil {
	// 	log.Println(err)
	// }

	//fmt.Println(requestBody)

	for i := 0; i < n; i++ {
		var reqStr string
		methodParams := ""
		//reqType := "GET"
		//reqStr := APIURL + randMethod + methodParams // Test Request
		randMethod := apiTest[rand.Intn(2)]
		var ticker models.Ticker
		ticker.SetTicker()
		switch randMethod {
		case "rates":
			// if rand.Intn(2) == 1 {
			methodParams = "&" + "currency=" + "RUB" // add Currency
			// } else {

			// }

		case "convert":

		case "profit":

		}
		reqStr = APIURL + "?" + randMethod + methodParams
		resp, err := http.Get(reqStr)
		//resp, err := http.NewRequest(reqType, reqStr, bytes.NewBuffer(requestBody)) // reqStr
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp.Body)

	}
}
