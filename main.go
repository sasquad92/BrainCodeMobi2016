package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sasquad92/BrainCodeMobi2016/gpio"
	"github.com/sasquad92/BrainCodeMobi2016/rooms"
	"net/http"
	"time"
)

var vicPos int
var murPos int

/*
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
*/

func AskRPi() {
	vicPos = gpio.Listen()
	// get msg
}

func BoardHandler(w http.ResponseWriter, r *http.Request) {

	AskRPi()
	//vic := gpio.Listen()
	// here id of murderer needed from another device
	mur := 4 // for tests

	b := rooms.ExportToJSON(vicPos, mur)

	w.Write(b)
}

func main() {

	err := gpio.InitPins()

	if err == nil {

		AskRPi()

		mx := mux.NewRouter()

		mx.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
		mx.HandleFunc("/board", BoardHandler)

		if err := http.ListenAndServe(":8080", mx); err != nil {
			panic(err)
			fmt.Println("ListenAndServe error.")
		}

		t := time.NewTicker(3 * time.Second)

		// cyclic call
		for {

			AskRPi()

			if vicPos == murPos {

				gpio.GameOver()
				time.Sleep(10 * time.Second)

				break
			}

			<-t.C
		}

		gpio.PinsOff()

	} else {
		fmt.Println("Error while maping pins on Raspberry Pi.", err)
	}
}
