package dde

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	ddeml "github.com/forget-the-bright/go-dde/ddeml"
	. "github.com/forget-the-bright/go-dde/types"
)

const (
	MAX_BUFFER_SIZE = 2147483647
)

type DdeMeta struct {
	IdInst    DWORD      //全局句柄
	HConv     HCONV      //会话句柄
	Callback  FNCALLBACK //回调函数
	AppName   VString    //服务名称
	TopicName VString    //主题名称
	Items     []VString  //单元项目字符串数组
	HszApp    HSZ        //服务字符串句柄
	HszTopic  HSZ        //主题字符串句柄
	HszItem   []HSZ      //单元项目字符串句柄数组
}
type DdeClient struct {
	DdeMeta
}

func (c *DdeClient) Run() bool {
	var cmd DWORD = DWORD(APPCLASS_STANDARD | APPCMD_CLIENTONLY)
	var res DWORD = 0
	if c.Callback == nil {
		c.Callback = c.MetaCallBackfunc
	}
	iReturn := ddeml.DdeInitialize(&c.IdInst, c.Callback, cmd, res)
	if int(iReturn) != DMLERR_NO_ERROR {
		fmt.Printf("DDE Initialization Failed")
	}
	c.HszApp = ddeml.DdeCreateStringHandle(c.IdInst, c.AppName, 0)
	c.HszTopic = ddeml.DdeCreateStringHandle(c.IdInst, c.TopicName, 0)
	c.HszItem = make([]HSZ, len(c.Items))
	for i := 0; i < len(c.Items); i++ {
		c.HszItem[i] = ddeml.DdeCreateStringHandle(c.IdInst, c.Items[i], 0)
	}
	c.HConv = ddeml.DdeConnect(c.IdInst, c.HszApp, c.HszTopic, nil)
	if c.HConv == HCONV(NULL) {
		fmt.Printf("DDE Connection Failed.\n")
		time.Sleep(2e9) //2e9 相当于2秒
		c.DestoryClient()
		return false
	}
	return true
}

func (c *DdeClient) DestoryClient() {
	ddeml.DdeFreeStringHandle(c.IdInst, c.HszApp)
	ddeml.DdeFreeStringHandle(c.IdInst, c.HszTopic)
	for i := 0; i < len(c.Items); i++ {
		ddeml.DdeFreeStringHandle(c.IdInst, c.HszItem[i])
	}
	fmt.Printf("客户端关闭")
	ddeml.DdeDisconnect(c.HConv)
	ddeml.DdeUninitialize(c.IdInst)
}

type DdeServer struct {
	DdeMeta
	close bool //服务状态
}

func (c *DdeServer) RunServer() bool {
	if c.Callback == nil {
		c.Callback = c.MetaCallBackfunc
	}
	iReturn := ddeml.DdeInitialize(&c.IdInst, c.Callback, DWORD(APPCLASS_STANDARD), 0)
	if int(iReturn) != DMLERR_NO_ERROR {
		fmt.Printf("DDE Initialization Failed")
	}
	c.HszApp = ddeml.DdeCreateStringHandle(c.IdInst, c.AppName, 0)
	c.HszTopic = ddeml.DdeCreateStringHandle(c.IdInst, c.TopicName, 0)
	c.HszItem = make([]HSZ, len(c.Items))
	for i := 0; i < len(c.Items); i++ {
		c.HszItem[i] = ddeml.DdeCreateStringHandle(c.IdInst, c.Items[i], 0)
	}
	sever := ddeml.DdeNameService(c.IdInst, c.HszApp, 0, UINT(DNS_REGISTER))
	defer c.DestoryServer()
	if int(sever) != TRUE {
		fmt.Printf("DdeNameService() failed!")
		return false
	}
	c.close = false
	fmt.Println("服务端开启")
	for {
		//do something
		var msg TagMSG
		if int(ddeml.GetMessage(&msg, 0, 0, 0)) <= 0 || c.close {
			break
		}
		ddeml.TranslateMessage(&msg)
		ddeml.DispatchMessage(&msg)
	}
	return true
}

func (c *DdeServer) DestoryServer() {
	c.close = true
	ddeml.DdeFreeStringHandle(c.IdInst, c.HszApp)
	ddeml.DdeFreeStringHandle(c.IdInst, c.HszTopic)
	for i := 0; i < len(c.Items); i++ {
		ddeml.DdeFreeStringHandle(c.IdInst, c.HszItem[i])
	}
	fmt.Println("服务端关闭")
	ddeml.DdeNameService(c.IdInst, c.HszApp, 0, UINT(DNS_UNREGISTER))
	ddeml.DdeUninitialize(c.IdInst)
}

func (c *DdeServer) IsHSZItem(hsz2 HSZ) HSZ {
	for i := 0; i < len(c.Items); i++ {
		if ddeml.DdeCmpStringHandles(HSZ(hsz2), c.HszItem[i]) <= 0 {
			return c.HszItem[i]
		}
	}
	return 0
}

func (c *DdeMeta) MetaCallBackfunc(wType UINT, wFmt UINT, hConv HCONV, hsz1 HSZ, hsz2 HSZ, hData HDDEDATA, dwData1 ULONG_PTR, dwData2 ULONG_PTR) HDDEDATA {
	return 0
}
func (c *DdeServer) MetaCallBackfunc(wType UINT, wFmt UINT, hConv HCONV, hsz1 HSZ, hsz2 HSZ, hData HDDEDATA, dwData1 ULONG_PTR, dwData2 ULONG_PTR) HDDEDATA {
	selectItem := c.IsHSZItem(hsz2)
	var res HDDEDATA
	switch wType {
	case XTYP_CONNECT:
		//接受客户端链接；
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), c.HszTopic) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), c.HszApp) <= 0 {
			res = HDDEDATA(TRUE)
		} else {
			res = HDDEDATA(FALSE) //接受客户端链接
		}
	case XTYP_ADVSTART:
		//客户端启动咨询循环。
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), c.HszTopic) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			res = HDDEDATA(TRUE)
		} else {
			res = HDDEDATA(FALSE) //接受客户端链接
		}
	case XTYP_ADVREQ:
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), c.HszTopic) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			res = c.SendMessage("XTYP_ADVREQ", selectItem, UINT(wFmt))
		}
		//广播数据；
	case XTYP_REQUEST:
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), c.HszTopic) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			res = c.SendMessage("XTYP_REQUEST哈哈", selectItem, UINT(wFmt))
		}
		//数据请求；
	case XTYP_POKE:
		if ddeml.DdeCmpStringHandles(HSZ(hsz1), c.HszTopic) <= 0 &&
			ddeml.DdeCmpStringHandles(HSZ(hsz2), selectItem) <= 0 {
			message := c.GetMessage(hData)
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
func (c *DdeMeta) Request(hszItem HSZ) string {
	hData := ddeml.DdeClientTransaction(nil, 0, c.HConv, hszItem, UINT(CF_TEXT),
		UINT(XTYP_REQUEST), 5000, nil)
	if hData == 0 {
		return "Request failed"
	} else {
		return DDEGetMessage(hData)
	}
}
func (c *DdeMeta) Poke(hszItem HSZ, szData string) {
	cs := C.CString(szData)
	ddeml.DdeClientTransaction((*BYTE)(unsafe.Pointer(cs)),
		DWORD(len(szData)+1),
		c.HConv, hszItem, UINT(CF_TEXT),
		UINT(XTYP_POKE), 3000, nil)
	C.free(unsafe.Pointer(cs))
}
func (c *DdeMeta) SendMessage(Src string, Item HSZ, Fmt UINT) HDDEDATA {
	cs := C.CString(Src)
	hdata := ddeml.DdeCreateDataHandle(c.IdInst, (*BYTE)(unsafe.Pointer(cs)), DWORD(len(Src)+1), 0, Item, Fmt, 0)
	C.free(unsafe.Pointer(cs))
	return hdata
}
func (c *DdeMeta) GetMessage(hData HDDEDATA) string {
	size := ddeml.DdeGetData(hData, nil, 0, 0)
	var str BYTE
	ddeml.DdeGetData(hData, &str, size, 0)
	buffer := (*[MAX_BUFFER_SIZE]byte)(unsafe.Pointer(&str))[:size-1]
	return string(buffer)
}

func DDERequest(idInst DWORD, hConv HCONV, hszItem HSZ) string {
	hData := ddeml.DdeClientTransaction(nil, 0, hConv, hszItem, UINT(CF_TEXT),
		UINT(XTYP_REQUEST), 5000, nil)
	if hData == 0 {
		return "Request failed"
	} else {
		return DDEGetMessage(hData)
	}
}
func DDEPoke(idInst DWORD, hConv HCONV, hszItem HSZ, szData string) {
	cs := C.CString(szData)
	ddeml.DdeClientTransaction((*BYTE)(unsafe.Pointer(cs)),
		DWORD(len(szData)+1),
		hConv, hszItem, UINT(CF_TEXT),
		UINT(XTYP_POKE), 3000, nil)
	C.free(unsafe.Pointer(cs))
}
func DDESendMessage(Inst DWORD, Src string, Item HSZ, Fmt UINT) HDDEDATA {
	cs := C.CString(Src)
	hdata := ddeml.DdeCreateDataHandle(Inst, (*BYTE)(unsafe.Pointer(cs)), DWORD(len(Src)+1), 0, Item, Fmt, 0)
	C.free(unsafe.Pointer(cs))
	return hdata
}
func DDEGetMessage(hData HDDEDATA) string {
	size := ddeml.DdeGetData(hData, nil, 0, 0)
	var str BYTE
	ddeml.DdeGetData(hData, &str, size, 0)
	buffer := (*[MAX_BUFFER_SIZE]byte)(unsafe.Pointer(&str))[:size-1]
	return string(buffer)
}
