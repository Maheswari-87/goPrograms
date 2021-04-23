package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/valyala/fastjson"
)

type Country struct {
	All string `json:"All"`
}

/*type All struct {
	Confirmed float64
	Recovered float64
	Deaths    float64
}*/

type Data struct {
	State     string
	Confirmed float64
	Recovered float64
	Deaths    float64
}

func homePage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "homepage")

	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body))

	if err != nil {
		log.Fatal(err)
	}
	responseObject := map[string]interface{}{}
	//var responseObject Result
	json.Unmarshal(data, &responseObject)
	stringdata := string(data)
	//fmt.Println(stringdata)
	var p fastjson.Parser
	v, err := p.Parse(stringdata)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(v)
	var keys []string
	v.GetObject().Visit(func(key []byte, values *fastjson.Value) {
		keys = append(keys, string(key))
	})
	fmt.Println(keys)

	//fmt.Println(all)
	//m := make(map[string]interface{})
	var Confirmed float64
	var Recovered float64
	var Deaths float64
	for _, i := range keys {
		//var all []string
		//v.GetObject().Visit(func(key []byte, values *fastjson.Value) {
		//	keys = append(all, string(key))
		//})
		//fmt.Println(all)
		//var cont Country
		//json.Unmarshal([]byte(data), &cont)
		//if cont.All == "All" {
		//all, err := json.Get("keys").GetIndex(0).String("All")
		//if err != nil {
		//	panic(err)
		//}
		//log.Println(all)
		//all = responseObject[i].(map[string]interface{})
		state := responseObject[i].(map[string]interface{})
		//if i != "All" {
		//state := responseObject[v].(map[string]interface{})
		//if key == "All" {
		//all := responseObject[key].(map[string]interface{})
		for _, value := range state {
			all := value.(map[string]interface{})
			//country=i
			for key, value := range all {
				if key == "confirmed" && value != nil {
					Confirmed = value.(float64)
				}
				if key == "recovered" && value != nil {
					Recovered = value.(float64)
				}
				if key == "deaths" && value != nil {
					Deaths = value.(float64)
				}
			}
		}
		//}
		//}
		//fmt.Println(v)
		//}

		fmt.Println(i)
		fmt.Println(Confirmed)
		fmt.Println(Recovered)
		fmt.Println(Deaths)

		//s := []float64{Confirmed, Recovered, Deaths}
		//var s []float64 =abc(Confirmed,Recovered,Deaths)
		//p1.Execute(w, s)
		p1, err := template.ParseFiles("html/country.html")
		data := Data{i, Confirmed, Recovered, Deaths}
		if err != nil {
			panic(err)
		}
		p1.Execute(w, data)
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":7034", nil))
}
func main() {
	handleRequests()
}

/*func main() {
	response, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
	if err != nil {
		fmt.Printf("the http request got failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := (ioutil.ReadAll(response.Body))

	if err != nil {
		log.Fatal(err)
	}
	//data1:=(string(data))
	var responseObject Result
	err = json.Unmarshal(data, &responseObject)
	if err != nil {
		panic(err)
	}

	//for key, value := range data {
	// Each value is an interface{} type, that is type asserted as a string
	//	fmt.Println(key, value.(string))
	//}
	//handleRequests()
	fmt.Printf("India:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.All.Confirmed, responseObject.All.Recovered, responseObject.All.Confirmed-responseObject.All.Recovered, responseObject.All.Deaths)
	fmt.Printf("\n\nAndaman:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Andaman.Confirmed, responseObject.Andaman.Recovered, responseObject.Andaman.Confirmed-responseObject.Andaman.Recovered, responseObject.Andaman.Deaths)
	fmt.Printf("\n\nAndhra:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Andhra.Confirmed, responseObject.Andhra.Recovered, responseObject.Andhra.Confirmed-responseObject.Andhra.Recovered, responseObject.Andhra.Deaths)
	fmt.Printf("\n\nArunachal pradesh:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Arunachal.Confirmed, responseObject.Arunachal.Recovered, responseObject.Arunachal.Confirmed-responseObject.Arunachal.Recovered, responseObject.Arunachal.Deaths)
	fmt.Printf("\n\nAssam:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Assam.Confirmed, responseObject.Assam.Recovered, responseObject.Assam.Confirmed-responseObject.Assam.Recovered, responseObject.Assam.Deaths)
	fmt.Printf("\n\nBihar:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Bihar.Confirmed, responseObject.Bihar.Recovered, responseObject.Bihar.Confirmed-responseObject.Bihar.Recovered, responseObject.Bihar.Deaths)
	fmt.Printf("\n\nChandigarh:\nConfirmed cases %v,\nRecovered cases %v,\nActive cases %v,\nDeaths %v.", responseObject.Chandigarh.Confirmed, responseObject.Chandigarh.Recovered, responseObject.Chandigarh.Confirmed-responseObject.Chandigarh.Recovered, responseObject.Chandigarh.Deaths)

	//fmt.Println(len(responseObject.IndianStates))
}*/

//out := Andaman1{Confirmed: responseObject.Andaman.Confirmed, Recovered: responseObject.Andaman.Recovered, Deaths: responseObject.Andaman.Deaths}
//out1 := Andhra1{Confirmed: responseObject.Andhra.Confirmed, Recovered: responseObject.Andhra.Recovered, Deaths: responseObject.Andhra.Deaths}
//out2 := Arunachal1{Confirmed: responseObject.Arunachal.Confirmed, Recovered: responseObject.Arunachal.Recovered, Deaths: responseObject.Arunachal.Deaths}
//p1, err := template.ParseFiles("html/covid.html")
//if err != nil {
//	panic(err)
//}
//p1.Execute(w)
//p1.Execute(w, out)
///p1.Execute(w, out1)
//p1.Execute(w, out2)
