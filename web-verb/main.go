// nuri file in go-lang repo
///joy note

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	getpostreq()
	getrequest()
}
func getrequest() {
	const myurl = "http://localhost:3000/"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	defer response.Body.Close()

}
func getpostreq() {
	const myurl = "http://localhost:3000/postform"
	data := url.Values{}
	data.Add("firstname", "joy adhikary")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))

}
