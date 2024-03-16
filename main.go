package main

import (
	"fmt"
	utils "gosnakego/utils"
)

func main() {
	var list utils.LinkedList[string]
	list.Prepend("60").Prepend("81").Prepend("80").RemoveFirstValueMatch("81")
	fmt.Println(list.GetValueByIndex(0))
}
