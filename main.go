package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct {
  ID string `json:id,omitempty`
  FirstName string `json:firstname,omitempty`
  LastName string `json:lastname,omitempty`
  Address *Address `json:address,omitempty`
}

type Address struct {
  City string `json:city,omitempty`
  State string `json:state,omitempty`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {

  //datos
  people = append(people, Person {ID:"1", FirstName:"Ryan", LastName:"Ray", Address:&Address{City:"Dublin", State:"California"} })
  people = append(people, Person {ID:"2", FirstName:"Francisca", LastName:"Cortés", Address:&Address{City:"Santiago", State:"Providencia"} })
  people = append(people, Person {ID:"3", FirstName:"Simón", LastName:"Arenas", Address:&Address{City:"Santiago", State:"Ñuñoa"} })
  people = append(people, Person {ID:"4", FirstName:"Sole", LastName:"Cortés", Address:&Address{City:"Santiago", State:"Providencia"} })
  people = append(people, Person {ID:"5", FirstName:"Roberto", LastName:"Olmos", Address:&Address{City:"Santiago", State:"Vitacura"} })

  router := mux.NewRouter()
  //endpoint
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", router))
}
