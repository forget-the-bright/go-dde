package demo

import (
	"fmt"
	"time"

	"github.com/forget-the-bright/go-dde/dde"
	"github.com/forget-the-bright/go-dde/ddeml"
	. "github.com/forget-the-bright/go-dde/types"
)

func isHSZItem(hsz2 HSZ) HSZ {
	for i := 0; i < len(items); i++ {
		if ddeml.DdeCmpStringHandles(HSZ(hsz2), hszItem[i]) <= 0 {
			return hszItem[i]
		}
	}
	return 0
}

func DdeCallback(wType UINT, wFmt UINT, hConv HCONV, hsz1 HSZ, hsz2 HSZ, hData HDDEDATA, dwData1 ULONG_PTR, dwData2 ULONG_PTR) HDDEDATA {
	selectItem := isHSZItem(hsz2)
	var res HDDEDATA
	switch wType {
	case XTYP_CONNECT:
		//接受客户端链接；
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), g_hszTopicName) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), g_hszAppName) <= 0 {
			res = HDDEDATA(TRUE)
		} else {
			res = HDDEDATA(FALSE) //接受客户端链接
		}
	case XTYP_ADVSTART:
		//客户端启动咨询循环。
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), g_hszTopicName) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			res = HDDEDATA(TRUE)
		} else {
			res = HDDEDATA(FALSE) //接受客户端链接
		}
	case XTYP_ADVREQ:
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), g_hszTopicName) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			res = dde.DDESendMessage(idInst, "XTYP_ADVREQ", selectItem, UINT(wFmt))
		}
		//广播数据；
	case XTYP_REQUEST:
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), g_hszTopicName) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			res = dde.DDESendMessage(idInst, "XTYP_REQUEST哈哈", selectItem, UINT(wFmt))
		}
		//数据请求；
	case XTYP_POKE:
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), g_hszTopicName) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			message := dde.DDEGetMessage(hData)
			fmt.Printf("message: %v\n", message)
			res = HDDEDATA(DDE_FACK)
		}
	case XTYP_DISCONNECT:
		fmt.Println("Disconnect notification received from server")
		//接受客户端发送的数据；
	default:
		res = HDDEDATA(NULL)
	}
	return res
}

//var callbacks FNCALLBACK = DdeCallback

var idInst DWORD = 0
var szApp VString = "Server"
var szTopic VString = "MyTopic"
var g_hszAppName HSZ
var g_hszTopicName HSZ
var items []VString = []VString{"MyItem0", "MyItem1", "MyItem2", "MyItem3", "MyItem4", "MyItem5", "MyItem6"}
var hszItem []HSZ = make([]HSZ, len(items))

func runClient() {
	var cmd DWORD = DWORD(APPCLASS_STANDARD | APPCMD_CLIENTONLY)
	var res DWORD = 0
	iReturn := ddeml.DdeInitialize(&idInst, DdeCallback, cmd, res)
	if iReturn != UINT(DMLERR_NO_ERROR) {

	}
	if int(iReturn) != DMLERR_NO_ERROR {
		fmt.Printf("DDE Initialization Failed")
	}
	var hConv HCONV // 会话句柄
	hszApp := ddeml.DdeCreateStringHandle(idInst, szApp, 0)
	hszTopic := ddeml.DdeCreateStringHandle(idInst, szTopic, 0)
	for i := 0; i < len(items); i++ {
		hszItem[i] = ddeml.DdeCreateStringHandle(idInst, items[i], 0)
	}
	hConv = ddeml.DdeConnect(idInst, hszApp, hszTopic, nil)
	ddeml.DdeFreeStringHandle(idInst, hszApp)
	ddeml.DdeFreeStringHandle(idInst, hszTopic)
	if hConv == 0 {
		fmt.Printf("DDE Connection Failed.\n")
		time.Sleep(2e9) //2e9 相当于2秒
		ddeml.DdeUninitialize(idInst)
	}
	for i := 0; i < len(items); i++ {
		data := dde.DDERequest(idInst, hConv, hszItem[i])
		fmt.Printf("data: %v\n", data)
		dde.DDEPoke(idInst, hConv, hszItem[i], data)
		time.Sleep(1e9)
	}
	fmt.Printf("客户端关闭")
	ddeml.DdeDisconnect(hConv)
	ddeml.DdeUninitialize(idInst)
}

func runServer() {
	iReturn := ddeml.DdeInitialize(&idInst, DdeCallback, DWORD(APPCLASS_STANDARD), 0)
	if int(iReturn) != DMLERR_NO_ERROR {
		fmt.Printf("DDE Initialization Failed")
	}
	g_hszAppName = ddeml.DdeCreateStringHandle(idInst, szApp, 0)

	g_hszTopicName = ddeml.DdeCreateStringHandle(idInst, szTopic, 0)
	for i := 0; i < len(items); i++ {
		hszItem[i] = ddeml.DdeCreateStringHandle(idInst, items[i], 0)
	}
	sever := ddeml.DdeNameService(idInst, g_hszAppName, 0, UINT(DNS_REGISTER))
	defer func() {
		ddeml.DdeFreeStringHandle(idInst, g_hszAppName)
		ddeml.DdeFreeStringHandle(idInst, g_hszTopicName)
		fmt.Printf("服务端关闭")
		ddeml.DdeNameService(idInst, g_hszAppName, 0, UINT(DNS_UNREGISTER))
		ddeml.DdeUninitialize(idInst)
	}()
	if int(sever) < 1 {
		fmt.Printf("DdeNameService() failed!")
	}
	fmt.Printf("服务端开启")
	for {
		//do something
		var msg TagMSG
		if int(ddeml.GetMessage(&msg, 0, 0, 0)) <= 0 {
			break
		}
		ddeml.TranslateMessage(&msg)
		ddeml.DispatchMessage(&msg)
	}
}
