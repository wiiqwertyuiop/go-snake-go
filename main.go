package main

import (
	"fmt"
	utils "gosnakego/utils"
)

func main() {
	list := utils.LinkedList[int]{}
	list.
		Prepend(60).
		Prepend(81).
		Prepend(80).
		RemoveFirstValueMatch(81).
		Prepend(62)

	if val, found := list.GetNodeByIndex(2); found {
		fmt.Println(val.Value)
	}
}
