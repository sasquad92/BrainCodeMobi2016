package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sasquad92/BrainCodeMobi2016/gpio"
	"github.com/sasquad92/BrainCodeMobi2016/rooms"
	"net/http"
)

func MurdererHandler(w http.ResponseWriter, r *http.Request) {
	mur := gpio.Listen()

	b := rooms.SendMurdererPos(mur)

	w.Write(b)
}

func main() {

	gpio.InitPins()

	mx := mux.NewRouter()

	//mx.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	mx.HandleFunc("/boardMur", MurdererHandler)

	if err := http.ListenAndServe(":8080", mx); err != nil {
		panic(err)
		fmt.Println("ListenAndServe error.")
	}

}
