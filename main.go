package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)





type Person struct {
    ID          string `json:"id,omitempty"`
    Firstname   string `json:"firstname,omitempty"`
    Lastname    string `json:"lastname,omitempty"`
    Address     *Address `json:"address,omitempty"`
}

type Address struct {
    City string `json:"city,omitempty"`
}

var people []Person

func getPerson(w http.ResponseWriter, req *http.Request){
    json.NewEncoder(w).Encode(people)
}

func getPersonById(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, req *http.Request){

}
func main() {
    fmt.Println("test")
    router := mux.NewRouter()
    people = append(people, 
        Person{ID:"1", 
        Firstname: "sanjok12", 
        Lastname:"dangol",
        Address: &Address{
            City: "ktm",
        }})

    people = append(people, 
        Person{ID:"2", 
        Firstname: "sabi", 
        Lastname:"maharjan",
        Address: &Address{
            City: "ktm",
        }})
    router.HandleFunc("/person", getPerson).Methods("GET")
    router.HandleFunc("/people/{id}", getPersonById).Methods("GET")
    router.HandleFunc("/person", createPerson).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", router))

}