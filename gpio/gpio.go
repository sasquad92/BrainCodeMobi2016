package gpio

import (
    "github.com/stianeikeland/go-rpio"
    //"time"
)

// Open memory range for GPIO access
func Open () bool {
    err := rpio.Open()
    var open bool
    
    if err == nil {
        open = true
    } else {
        open = false
    }
    return open
}

func GameOver() {
    err := rpio.Open()
    
    defer rpio.Close()
    
    if err == nil {
        pin := rpio.Pin(26)
        
        pin.Output()
        pin.High()
    }
}