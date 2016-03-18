package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
    "github.com/sasquad92/BrainCodeMobi2016/gpio"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	name := vars["name"]
	job := vars["job"]
	age := vars["age"]
	w.Write([]byte(fmt.Sprintf("Name is %s ", name)))
	w.Write([]byte(fmt.Sprintf("Job is %s ", job)))
	w.Write([]byte(fmt.Sprintf("Age is %s ", age)))
}

func main() {
        
	mx := mux.NewRouter()
    
    mx.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	//to handle URL like
	//http://website:8080/person/Boo/CEO/199

	//http://website:8080/person/Boo/CEO/199 <- if age > 199, will cause 404 error
    
    gpio.GameOver()

	if err := http.ListenAndServe(":8080", mx); err != nil {
        panic(err)
    }

}
