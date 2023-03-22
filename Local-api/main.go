//joy code details note

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
	Boss        *Boss   `json:"boss"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

type Boss struct {
	Fullname string `json:"fullname"`
	Lastname string `json:"lastname"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")                 // router r r.HandleFunc("rought", controller).Methods("methods name ")
	r.HandleFunc("/all", getallcourses).Methods("GET")          // localhost:4000/courses dile ei rought ta show korbe
	r.HandleFunc("/one/{id}", getonecourse).Methods("GET")      // localhost:4000/courses/{id ta dibo kontar jonno dekhbo}
	r.HandleFunc("/onec", createonecourse).Methods("POST")      // POST is often used by World Wide Web to send user generated data to the web server or when you upload file.
	r.HandleFunc("/oneup/{id}", updateonecourse).Methods("PUT") //PUT method is used to update resource available on the server
	r.HandleFunc("/onede/{id}", deletecourse).Methods("DELETE") // DELETE use for delete somethings form the db

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}, Boss: &Boss{Fullname: "Hitesh", Lastname: "Choudhary"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Joy Adhikary", Website: "go.dev"}, Boss: &Boss{Fullname: "JOY", Lastname: "Adhikary"}})
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Joy you will be the boss .. just keep working </h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // map type key value
	json.NewEncoder(w).Encode(courses)                 // it convert slice courses into json value
}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getonecourse id")
	w.Header().Set("content-Type ", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("not found")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	var course Course

	if r.Body == nil {
		json.NewEncoder(w).Encode("plz enter some data")
		return
	}

	//check doublicate
	for _, course1 := range courses {
		if course1.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("already exist")
			return
		}
	}

	json.NewDecoder(r.Body).Decode(&course)
	rand.Seed(time.Now().UnixNano())               ///create random number
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert int to string
	courses = append(courses, course)              // push it on courses
	json.NewEncoder(w).Encode(course)              // sending json file format
}

func updateonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
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

func deletecourse(w http.ResponseWriter, r *http.Request) {
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
