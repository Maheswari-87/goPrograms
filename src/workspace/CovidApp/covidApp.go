package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	States IndianStates `json:"Andhra Pradesh"`
}
type IndianStates struct {
	Confirmed int `json:"confirmed"`
	Recovered int `json:"recovered"`
	Deaths    int `json:"Deaths"`
}

/*type Result struct {
	States IndianStates `json:"Andhra Pradesh"`
}
type IndianStates struct {
	Confirmed int `json:"confirmed"`
	Recovered int `json:"recovered"`
	Deaths    int `json:"Deaths"`
}*/

func main() {
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//data1:=(string(data))
	var responseObject Result
	err = json.Unmarshal(data, &responseObject)
	if err != nil {
		panic(err)
	}
	//fmt.Println(responseObject.States.Confirmed, responseObject.States.Recovered, responseObject.States.Deaths)
	fmt.Printf("Confirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.States.Confirmed, responseObject.States.Recovered, responseObject.States.Confirmed-responseObject.States.Recovered, responseObject.States.Deaths)
	//fmt.Println(len(responseObject.IndianStates))
}
