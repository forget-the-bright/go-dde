# go-dde

#### 介绍
go 语言来实现调用windowapi,实现动态数据交换 Dynamic Data Exchange

#### 软件架构
go


#### 安装教程
```
go get github.com/forget-the-bright/go-dde
```

#### 使用说明
```这只是个简单的例子```
```
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
```
```server 的回调函数可以自己定义```
```
func  MetaCallBackfunc(wType UINT, wFmt UINT, hConv HCONV, hsz1 HSZ, hsz2 HSZ, hData HDDEDATA, dwData1 ULONG_PTR, dwData2 ULONG_PTR) HDDEDATA {
    //do something
	return 0
}
func runServer() {
	ddeser := dde.DdeServer{}
	ddeser.AppName = "Server"
	ddeser.TopicName = "MyTopic"
	ddeser.Items = []VString{"MyItem0", "MyItem1", "MyItem2", "MyItem3", "MyItem4", "MyItem5", "MyItem6"}
	//defer ddeser.DestoryServer()
    ddeser.Callback = MetaCallBackfunc
	ddeser.RunServer()
}
func main() {
	runServer()
}
```
