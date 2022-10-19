package main

import (
	"fmt"
	"time"

	"github.com/forget-the-bright/go-dde/dde"
	. "github.com/forget-the-bright/go-dde/types"
)

func runClient() {
	ddecli := dde.DdeClient{}
	ddecli.AppName = "Server"
	ddecli.TopicName = "MyTopic"
	ddecli.Items = []VString{"MyItem0", "MyItem1", "MyItem2", "MyItem3", "MyItem4", "MyItem5", "MyItem6"}
	ddecli.Run()
	for i := 0; i < len(ddecli.Items); i++ {
		data := ddecli.Request(ddecli.HszItem[i])
		fmt.Printf("data: %v\n", data)
		ddecli.Poke(ddecli.HszItem[i], data)
		time.Sleep(1e9)
	}
}
func runServer() {
	ddeser := dde.DdeServer{}
	ddeser.AppName = "Server"
	ddeser.TopicName = "MyTopic"
	ddeser.Items = []VString{"MyItem0", "MyItem1", "MyItem2", "MyItem3", "MyItem4", "MyItem5", "MyItem6"}
	//defer ddeser.DestoryServer()

	ddeser.RunServer()
}
func main() {
	//runClient()
	runServer()
}
