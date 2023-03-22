package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//model for course
//build up api json
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fakeDB amra  clint ke pass  or name info gula share korte  chai na database acess  hide rakhbo ejonno eta

var courses []Course

// middleware.helper -file
func (c *Course) IsEmpty() bool {
	return c.CourseId == " "
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")                  // router r r.HandleFunc("rought", controller).Methods("methods name ")
	r.HandleFunc("/all", getAllCourses).Methods("GET")           // localhost:4000/courses dile ei rought ta show korbe
	r.HandleFunc("/one/{id}", getOneCourses).Methods("GET")      // localhost:4000/courses/{id ta dibo kontar jonno dekhbo}
	r.HandleFunc("/onec", CreateOneCourses).Methods("POST")      // POST is often used by World Wide Web to send user generated data to the web server or when you upload file.
	r.HandleFunc("/oneup/{id}", UpdateOneCourses).Methods("PUT") //PUT method is used to update resource available on the server
	r.HandleFunc("/onede/{id}", deleteCourse).Methods("DELETE")  // DELETE use for delete somethings form the db

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Ishmoth  Ura Nuri", Website: "go.dev"}})
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controlle api json and sending  control
//serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome</h>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get  all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

//request id  for one course

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get  one courses")
	w.Header().Set("content-type", "application/json")

	//id form request
	params := mux.Vars(r)
	//loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course  found ")
	return
}

//add course control
func CreateOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create  one courses")
	w.Header().Set("content-type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data ")
	}
	//what about- {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON ")
		return
	}

	//generate unique id,string
	// append course  into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

//update one course control
func UpdateOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update  one courses")
	w.Header().Set("content-type", "application/json")

	//first - grab id form req
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode("hey value is succesfully updated ")
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete course ")
	w.Header().Set("content-type ", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("sucessfully deleted ")
			return
		}
	}
}
