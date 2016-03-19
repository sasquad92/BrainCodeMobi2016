package gpio

import (
	"github.com/stianeikeland/go-rpio"
    "fmt"
)

// global variables are private
// all functions are public

var pin11 rpio.Pin // Death diod buahaha
var pin13 rpio.Pin
var pin15 rpio.Pin
var pin16 rpio.Pin
var pin18 rpio.Pin
var pin19 rpio.Pin
var pin21 rpio.Pin
var pin22 rpio.Pin
var pin23 rpio.Pin
var pin24 rpio.Pin
var pin26 rpio.Pin

// InitPins inits all needed pins and maps the memory
func InitPins() (err error) {

    err = rpio.Open()
    
    if err != nil {
        fmt.Println(err)
        fmt.Println("Error with mapping memory (RPi)")
    }
    
	pin13 = rpio.Pin(13)
	pin15 = rpio.Pin(15)
	pin16 = rpio.Pin(16)
	pin18 = rpio.Pin(18)
	pin19 = rpio.Pin(19)
	pin21 = rpio.Pin(21)
	pin22 = rpio.Pin(22)
	pin23 = rpio.Pin(23)
	pin24 = rpio.Pin(24)
	pin26 = rpio.Pin(26)
    
    return err
}

// GameOver puts high state on pin11
func GameOver() {

		pin11.Output()
		pin11.High()
}

// PinsOff puts low state to all needed pins and unmaps memmory
func PinsOff() {

	pin13.Low()
	pin15.Low()
	pin16.Low()
	pin18.Low()
	pin19.Low()
	pin21.Low()
	pin22.Low()
	pin23.Low()
	pin24.Low()
	pin26.Low()

	rpio.Close()
}

// Listen is listening for high state on any diod
// Returns 0 when all circuits are open
// To change, becouse it is handing is STUPID!
func Listen() (int) {
    var pin int // room id
    
    res13 := pin13.Read()
    res15 := pin15.Read()
    res16 := pin16.Read()
    res18 := pin18.Read()
    res19 := pin19.Read()
    res21 := pin21.Read()
    res22 := pin22.Read()
    res23 := pin23.Read()
    res24 := pin24.Read()
    res26 := pin26.Read()
    
    if res13 > 1 {
        pin = 1
    }
    if res15 > 1 {
        pin = 2
    }
    if res16 > 1 {
        pin =3
    }
    if res18 > 1 {
        pin = 4
    }
    if res19 > 1 {
        pin = 5
    }
    if res21 > 1 {
        pin = 6
    }
    if res22 > 1 {
        pin = 7
    }
    if res23 > 1 {
        pin = 8
    }
    if res24 > 1 {
        pin = 9
    }
    if res26 > 1 {
        pin = 10
    }
    
    return pin
}