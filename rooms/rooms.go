package rooms

import (
    "encoding/json"
    "fmt"
)

/*
type Room struct {
    ID uint `json:"id"`
    Victim bool `json:"victim"`
    Murderer bool `json:"murderer"`
    Exit bool `json:"exit"`
}
*/

// Place holds room ids where Victim and Murderer are
type Place struct {
    Victim int `json:"victim"`
    Murderer int `json:"murderer"`
}

// ExportToJSON returns []byte
func ExportToJSON(vic int, mur int) []byte {
    
    place := Place {
        Victim: vic,
        Murderer: mur,
    }
    
    b, err := json.Marshal(place)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    
    return b
}