package main

import (
	"fmt"
	utils "gosnakego/utils"
)

func main() {
	list := utils.LinkedList[int]{}
	list.
		Prepend(2).
		Prepend(60).
		Prepend(81).
		Prepend(80).
		RemoveFirstValueMatch(60).
		Prepend(43).
		RemoveFirstValueMatch(81).
		Prepend(62)

	if val, found := list.GetNodeByIndex(3); found {
		fmt.Println(val.Value)
	}
}
