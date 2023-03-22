package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const url11 string = "https://lco.dev"
const url1 string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	result, _ := url.Parse(url1)
	conn := result.String()
	fmt.Println(conn)

	con := string(result.Port())
	fmt.Println(con)
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Printf("type of this result is %T", result)
	qu := result.Query()
	fmt.Println(qu)
	for key, val := range qu {
		fmt.Println(key, "=", val)
	}

	response, err := http.Get(url11)

	if err != nil {
		panic(err)
	}
	fmt.Printf("type of this request is %T", response)

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	// var responseString strings.Builder
	// responseString.Write(data)
	// fmt.Println(responseString.String())
	con1 := string(data)
	fmt.Println(con1)
}
