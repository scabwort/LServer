package main

import (
	"fmt"
	. "types"
	"utils"
)

func main() {
	fmt.Println("station start")

	FromNode()

	v := Recurlyservers{}
	utils.GetFileXml("server.xml", &v)
	fmt.Printf("%v\n", v)
	//fmt.Printf("app path:%s\n", utils.GetAppPath())

}
