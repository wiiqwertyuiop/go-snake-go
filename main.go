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
		RemoveFirstValueMatch(81)

	if val, found := list.GetNodeByIndex(1); found {
		fmt.Println(val.Value)
	}
}
