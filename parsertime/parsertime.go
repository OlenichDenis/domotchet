package main

import (
	"time"

)

func main() {

	var city = "Хабаровск"
	//var reg = "Хабаровский край"
	//var regSrc = "2020-02-23T14-51-44z03-00-e7f63e"

	//fmt.Println(dateto, "dateto 1")

	for {
		err :=   handler (city)
		if err != nil {
			err =   handler (city)
		}
		time.Sleep(60 * time.Second)
	}

	//r := mux.NewRouter()
	//r.HandleFunc("/", handler)
	//r.HandleFunc("/static/", static)
	//log.Fatal(http.ListenAndServe(":8001", r))

}

