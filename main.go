package main

import (
	"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"github.com/sasquad92/BrainCodeMobi2016/gpio"
	"github.com/sasquad92/BrainCodeMobi2016/rooms"
	"net/http"
	//"time"
)

var vicPos int
var murPos int

func AskRPi() {

	vicPos = gpio.Listen()

	resp, err := http.Get("http://192.168.0.105/boardMur")

	if err != nil {
		panic(err)
	}

	var data rooms.Place // json know structure type

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}
	murPos = data.Murderer
}

func BoardHandler(w http.ResponseWriter, r *http.Request) {

	AskRPi()
	b := rooms.ExportToJSON(vicPos, murPos)

	w.Write(b)
}

func main() {

    //tmp
    gpio.InitPins()
    gpio.Blink26()
    gpio.PinsOff()
    
    /*
	err := gpio.InitPins()

	if err == nil {

		AskRPi()

		mx := mux.NewRouter()

		mx.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
		mx.HandleFunc("/board", BoardHandler)

		if err := http.ListenAndServe(":80", mx); err != nil {
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
*/
}
