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
		Prepend(43).
		RemoveFirstValueMatch(80).
		Prepend(62)

	if val, found := list.PopWithValueOnIndex(1); found {
		fmt.Println(val)
		fmt.Println(
			list.
				Prepend(33).
				Prepend(42).
				ReadValueOnIndex(3))
	}
}
