package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	// type to json
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapBytes, _ := json.Marshal(mapD)
	fmt.Println(string(mapBytes))

	// json to type
	amap := make(map[string]interface{})
	if err := json.Unmarshal(mapBytes, &amap); err != nil {
		fmt.Println(err)
	}
	fmt.Println(amap)

	// type struct to json
	u := user{Name: "sofikul", Age: 20}
	userJSON, _ := json.Marshal(u)
	fmt.Println(string(userJSON))

	// json to struct typ
	//u1 := user{}
	userStr := `{"name":"sofikul","age":20}`
	if err := json.Unmarshal([]byte(userStr), &amap); err != nil {
		fmt.Println("An error occurred to unmarshal the json ", err)
	}

	fmt.Printf("%+v", amap)
}
