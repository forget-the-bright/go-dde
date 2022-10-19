package ddeml

import (
	. "go-dde/types"
	"syscall"
	"unsafe"
)

func init() {
	user32 := syscall.MustLoadDLL("user32")
	DdeInitialize = func(
		Inst *DWORD,
		Callback FNCALLBACK, //any
		Cmd DWORD,
		Res DWORD) UINT {
		proc := user32.MustFindProc("DdeInitializeW")
		res, _, _ := proc.Call(
			uintptr(unsafe.Pointer(Inst)),
			syscall.NewCallback(Callback),
			uintptr(Cmd),
			0)
		return UINT(res)
	}
	DdeUninitialize = func(Inst DWORD) BOOL {
		proc := user32.MustFindProc("DdeUninitialize")
		res, _, _ := proc.Call(uintptr(Inst))
		return BOOL(res)
	}
	DdeConnectList = func(
		Inst DWORD,
		Service HSZ,
		Topic HSZ,
		ConvList HCONVLIST,
		CC *CONVCONTEXT) HCONVLIST {
		proc := user32.MustFindProc("DdeConnectList")
		res, _, _ := proc.Call(
			uintptr(Inst),
			uintptr(Service),
			uintptr(Topic),
			uintptr(ConvList),
			uintptr(unsafe.Pointer(CC)))
		return HCONVLIST(res)
	}

	DdeQueryNextServer = func(
		ConvList HCONVLIST, ConvPrev HCONV) HCONV {
		proc := user32.MustFindProc("DdeQueryNextServer")
		res, _, _ := proc.Call(uintptr(ConvList), uintptr(ConvPrev))
		return HCONV(res)
	}

	DdeDisconnectList = func(ConvList HCONVLIST) BOOL {
		proc := user32.MustFindProc("DdeDisconnectList")
		res, _, _ := proc.Call(uintptr(ConvList))
		return BOOL(res)
	}

	DdeConnect = func(
		Inst DWORD,
		Service HSZ,
		Topic HSZ,
		CC *CONVCONTEXT) HCONV {
		proc := user32.MustFindProc("DdeConnect")
		res, _, _ := proc.Call(
			uintptr(Inst),
			uintptr(Service),
			uintptr(Topic),
			uintptr(unsafe.Pointer(CC)),
		)
		return HCONV(res)
	}

	DdeDisconnect = func(Conv HCONV) BOOL {
		proc := user32.MustFindProc("DdeDisconnect")
		res, _, _ := proc.Call(uintptr(Conv))
		return BOOL(res)
	}

	DdeReconnect = func(Conv HCONV) HCONV {
		proc := user32.MustFindProc("DdeReconnect")
		res, _, _ := proc.Call(uintptr(Conv))
		return HCONV(res)
	}

	DdeQueryConvInfo = func(
		Conv HCONV,
		Transaction DWORD,
		ConvInfo *CONVINFO) UINT {
		proc := user32.MustFindProc("DdeQueryConvInfo")
		res, _, _ := proc.Call(uintptr(Conv),
			uintptr(Transaction),
			uintptr(unsafe.Pointer(ConvInfo)))
		return UINT(res)
	}

	DdeSetUserHandle = func(
		Conv HCONV, Id DWORD, User DWORD_PTR) BOOL {
		proc := user32.MustFindProc("DdeSetUserHandle")
		res, _, _ := proc.Call(uintptr(Conv), uintptr(Id), uintptr(User))
		return BOOL(res)
	}

	DdeAbandonTransaction = func(
		Inst DWORD, Conv HCONV, Transaction DWORD) BOOL {
		proc := user32.MustFindProc("DdeAbandonTransaction")
		res, _, _ := proc.Call(uintptr(Inst), uintptr(Conv), uintptr(Transaction))
		return BOOL(res)
	}

	DdePostAdvise = func(
		Inst DWORD, Topic HSZ, Item HSZ) BOOL {
		proc := user32.MustFindProc("DdePostAdvise")
		res, _, _ := proc.Call(uintptr(Inst), uintptr(Topic), uintptr(Item))
		return BOOL(res)
	}

	DdeEnableCallback = func(
		Inst DWORD, Conv HCONV, Cmd UINT) BOOL {
		proc := user32.MustFindProc("DdeEnableCallback")
		res, _, _ := proc.Call(uintptr(Inst), uintptr(Conv), uintptr(Cmd))
		return BOOL(res)
	}

	DdeImpersonateClient = func(Conv HCONV) BOOL {
		proc := user32.MustFindProc("DdeImpersonateClient")
		res, _, _ := proc.Call(uintptr(Conv))
		return BOOL(res)
	}

	DdeNameService = func(
		Inst DWORD, S1 HSZ, S2 HSZ, Cmd UINT) HDDEDATA {
		proc := user32.MustFindProc("DdeNameService")
		res, _, _ := proc.Call(uintptr(Inst), uintptr(S1), uintptr(S2), uintptr(Cmd))
		return HDDEDATA(res)
	}

	DdeClientTransaction = func(
		Data *BYTE, //*BYTE
		cData DWORD,
		Conv HCONV,
		Item HSZ,
		Fmt UINT,
		Type UINT,
		Timeout DWORD,
		Result *DWORD) HDDEDATA {
		proc := user32.MustFindProc("DdeClientTransaction")
		res, _, _ := proc.Call(
			uintptr(unsafe.Pointer(Data)),
			uintptr(cData),
			uintptr(Conv),
			uintptr(Item),
			uintptr(Fmt),
			uintptr(Type),
			uintptr(Timeout),
			uintptr(unsafe.Pointer(Result)))
		return HDDEDATA(res)
	}

	DdeCreateDataHandle = func(
		Inst DWORD,
		Src *BYTE,
		C DWORD,
		Off DWORD,
		Item HSZ,
		Fmt UINT,
		Cmd UINT) HDDEDATA {
		proc := user32.MustFindProc("DdeCreateDataHandle")
		res, _, _ := proc.Call(
			uintptr(Inst),
			uintptr(unsafe.Pointer(Src)),
			uintptr(C),
			uintptr(Off),
			uintptr(Item),
			uintptr(Fmt),
			uintptr(Cmd))
		return HDDEDATA(res)
	}

	DdeAddData = func(
		Data HDDEDATA,
		Src *BYTE,
		C DWORD,
		Off DWORD) HDDEDATA {
		proc := user32.MustFindProc("DdeAddData")
		res, _, _ := proc.Call(uintptr(Data),
			uintptr(unsafe.Pointer(Src)),
			uintptr(C),
			uintptr(Off))
		return HDDEDATA(res)
	}

	DdeGetData = func(
		Data HDDEDATA,
		Dst *BYTE,
		Max DWORD,
		Off DWORD) DWORD {
		proc := user32.MustFindProc("DdeGetData")
		res, _, _ := proc.Call(uintptr(Data),
			uintptr(unsafe.Pointer(Dst)),
			uintptr(Max),
			uintptr(Off))
		return DWORD(res)
	}

	DdeAccessData = func(
		Data HDDEDATA, DataSize *DWORD) *BYTE {
		proc := user32.MustFindProc("DdeAccessData")
		res, _, _ := proc.Call(uintptr(Data),
			uintptr(unsafe.Pointer(DataSize)))
		return (*BYTE)(unsafe.Pointer(res))
	}

	DdeUnaccessData = func(Data HDDEDATA) BOOL {
		proc := user32.MustFindProc("DdeUnaccessData")
		res, _, _ := proc.Call(uintptr(Data))
		return BOOL(res)
	}

	DdeFreeDataHandle = func(Data HDDEDATA) BOOL {
		proc := user32.MustFindProc("DdeFreeDataHandle")
		res, _, _ := proc.Call(uintptr(Data))
		return BOOL(res)
	}

	DdeGetLastError = func(Inst DWORD) UINT {
		proc := user32.MustFindProc("DdeGetLastError")
		res, _, _ := proc.Call(uintptr(Inst))
		return UINT(res)
	}

	DdeCreateStringHandle = func(
		Inst DWORD, S VString, CodePage int) HSZ {
		u, _ := syscall.UTF16PtrFromString(string(S))
		proc := user32.MustFindProc("DdeCreateStringHandleW")
		res, _, _ := proc.Call(uintptr(Inst), uintptr(unsafe.Pointer(u)), uintptr(CodePage))
		return HSZ(res)
	}

	DdeQueryString = func(
		Inst DWORD,
		S HSZ,
		OS OVString,
		hMax DWORD,
		CodePage int) DWORD {
		proc := user32.MustFindProc("DdeQueryStringW")
		res, _, _ := proc.Call(
			uintptr(Inst),
			uintptr(S),
			uintptr(unsafe.Pointer(&OS)),
			uintptr(hMax),
			uintptr(CodePage))
		return DWORD(res)
	}

	DdeFreeStringHandle = func(Inst DWORD, S HSZ) BOOL {
		proc := user32.MustFindProc("DdeFreeStringHandle")
		res, _, _ := proc.Call(uintptr(Inst),
			uintptr(S))
		return BOOL(res)
	}

	DdeKeepStringHandle = func(Inst DWORD, S HSZ) BOOL {
		proc := user32.MustFindProc("DdeKeepStringHandle")
		res, _, _ := proc.Call(uintptr(Inst),
			uintptr(S))
		return BOOL(res)
	}

	DdeCmpStringHandles = func(S1 HSZ, S2 HSZ) int {
		proc := user32.MustFindProc("DdeCmpStringHandles")
		res, _, _ := proc.Call(uintptr(S1),
			uintptr(S2))
		return int(res)
	}

	GetMessage = func(lpMsg LPMSG, hWnd HWND, wMsgFilterMin UINT, wMsgFilterMax UINT) BOOL {
		proc := user32.MustFindProc("GetMessageW")
		res, _, _ := proc.Call(uintptr(unsafe.Pointer(lpMsg)),
			uintptr(hWnd), uintptr(wMsgFilterMin), uintptr(wMsgFilterMax))
		return BOOL(res)
	}
	TranslateMessage = func(lpMsg LPMSG) BOOL {
		proc := user32.MustFindProc("TranslateMessage")
		res, _, _ := proc.Call(uintptr(unsafe.Pointer(lpMsg)))
		return BOOL(res)
	}
	DispatchMessage = func(lpMsg LPMSG) BOOL {
		proc := user32.MustFindProc("DispatchMessageW")
		res, _, _ := proc.Call(uintptr(unsafe.Pointer(lpMsg)))
		return BOOL(res)
	}
}
