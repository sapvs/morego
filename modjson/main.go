package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolFalse, _ := json.Marshal(true)
	fmt.Println(string(bolFalse))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	res1D := &response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.MarshalIndent(res1D, " ", "    ")
	fmt.Println(string(res1B))

	err := error.Error("dasd")
	


}
