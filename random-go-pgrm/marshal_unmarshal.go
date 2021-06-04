package main

import (
	"encoding/json"
	"fmt"
)

type Hello struct {
	Anishg string `json:"anis,omitempty"`
	Rahman string `json:"rahman,omitempty"`
}

type Bello struct {
	Anislkj string `json:"anis,omitempty"`
	mia     string `json:"mia,omitempty"`
}

func hello_func(object interface{}) error {

	dat, ok := object.(*Bello)
	if ok {
		fmt.Println(dat)
	}

	anisByte, err := json.Marshal(object)
	var anis Bello

	err = json.Unmarshal(anisByte, &anis)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(anis)
	return nil
}

func main() {
	anis := &Hello{
		Anishg: "anis",
		Rahman: "rahman",
	}
	err := hello_func(anis)
	if err != nil {
		fmt.Println(err.Error())
	}
}
