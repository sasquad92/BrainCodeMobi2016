package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
    //"github.com/sasquad92/BrainCodeMobi2016/gpio"
    "github.com/sasquad92/BrainCodeMobi2016/rooms"
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
    
    
    //if err := gpio.InitPins(); err != nil {
        //gpio.GameOver()
        // sth
        // sth
        // sth
    
        //gpio.PinsOff()
    //}
    
    // JSON test
    b := rooms.ExportToJSON(1, 2)
    
    fmt.Println("JSON")
    fmt.Printf("%s", b)

	if err := http.ListenAndServe(":8080", mx); err != nil {
        panic(err)
        fmt.Println("ListenAndServe error.")
    }

}
