package main

import (

	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course
type  Course struct{
    CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author  `json:"author"`

}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake db
var courses []Course

//middleware,helper
func  (c *Course)IsEmpty() bool{
   //return c.CourseId == "" && c.CourseName == ""  
   return c.CourseName == ""  
   
}

func main(){

	fmt.Println("Api --GoLang First api is working")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Go Basics",
		CoursePrice: 100,
		Author:      &Author{Fullname: "John Doe", Website: "johndoe.com"},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Go REST-Api",
		CoursePrice: 999,
		Author:      &Author{Fullname: "Hc", Website: "hc123.com"},
	})


	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
    r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	//listen to port
	log.Fatal(http.ListenAndServe(":7000",r))

}

//controller--file

// serve home route
//writer --> w-->write from where server send resppnse==Response
//Readder--> r -->reader from where you get the value==Request
func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang 🥶 first Api</h1>"))
}

//Get all courses
func getAllCourses(w http.ResponseWriter,r *http.Request){
    fmt.Println("Get all Courses")
	w.Header().Set("Content-Type","applicatioan/json")
    json.NewEncoder(w).Encode(courses)
}


//get one course by id
func getOneCourse(w http.ResponseWriter,r *http.Request){
   fmt.Println("get one course")
   w.Header().Set("Content-Type","applicatioan/json")

   //grab id from the req
   params := mux.Vars(r)

   // loop through course & send reposnse
   for _,course := range courses{
	if course.CourseId == params["id"] {
		json.NewEncoder(w).Encode(course)
		return
	}
   }

   json.NewEncoder(w).Encode("No Course Found.")
   return 
}


//crate course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide all data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique ID
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}


// Update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			// remove old course
			courses = append(courses[:index], courses[index+1:]...)

			// decode updated course
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]

			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}

// Delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}