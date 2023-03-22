package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Handling get-post")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostRequest()
}

// get request
func PerformGetRequest() {
	const myurl = "http://localhost:8080/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("content length is :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

//post -req
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8080/post"

	requestBody := strings.NewReader(`
	{
		"age":"23",
		"post":"software devloper",
		"company":"adfinix"
	}
`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

//postform-request
func PerformPostRequest() {
	const myurl = "http://localhost:8080/postform"

	data := url.Values{}
	data.Add("firstname", "ishmoth")
	data.Add("lastname", "nuri")
	data.Add("email", "joy@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
