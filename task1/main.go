package main

import (
	"encoding/json"
	"fmt"
)

const (
	inputJSONOne = `{"type":1,"result":["res1","res2"]}`
	inputJSONTwo = `{"type":2,"result":{"0":"res1","1":"res2"}}`
)

type Type1 struct {
	Type   int      `json:"type"`
	Result []string `json:"result"`
}
type Type2 struct {
	Type   int `json:"type"`
	Result struct {
		Num0 string `json:"0"`
		Num1 string `json:"1"`
	} `json:"result"`
}

func main() {
	UnmarshalSwitchType(inputJSONOne)
	UnmarshalSwitchType(inputJSONTwo)
}

func UnmarshalSwitchType(strJson string) {

	typeStruct := struct {
		Type int
	}{}

	if err := json.Unmarshal([]byte(strJson), &typeStruct); err != nil {
		fmt.Println("Unmarshal error :", err)
	}

	switch typeStruct.Type {
	case 1:
		type1 := Type1{}
		if err := json.Unmarshal([]byte(strJson), &type1); err != nil {
			panic(err)
		}
		fmt.Println("type 1", type1)
	case 2:
		type2 := Type2{}
		if err := json.Unmarshal([]byte(strJson), &type2); err != nil {
			panic(err)
		}
		fmt.Println("type 2", type2)
	default:
		fmt.Println("different type")
	}
}
