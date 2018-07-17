package main

import (
	"encoding/json"
	"fmt"
)

type loopDetail struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Group string `json:"group,omitempty"`
}

type ads struct {
	Loop map[string]*loopDetail `json:"loop,omitempty"`
}

func main() {
	data := new(ads)
	final := new(loopDetail)
	final.Id = 1
	final.Name = "newLoop"
	final.Group = "group"
	data.Loop = make(map[string]*loopDetail)
	data.Loop["final"] = final
	data.Loop["work"] = final
	callData := new(ads)
	data.transform(&callData)
	a := new(int)

	fmt.Println("Hello, playground", data, callData.Loop["final"], a, &a)
	testing(a, &a)
}

func (dData *ads) transform(data interface{}) {
	jdata, err := json.Marshal(dData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jdata))
	err = json.Unmarshal(jdata, data)
	if err != nil {
		fmt.Println("Not worked")
	}
}

func testing(a *int, b **int) {
	var c int
	c = 4
	a = &c
	b = &a
	fmt.Println(*a, " and ", a, *b)
}
