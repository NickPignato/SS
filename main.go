package main

import (
	"encoding/json"
)

func main() {
	hw, err := GetHardware()
	if err != nil {
		panic(err)
	}
	hwJSON, _ := json.Marshal(hw)
	println(string(hwJSON))

}
