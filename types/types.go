package types

var (
	DMLERR_NO_ERROR            = 0x0000
	DMLERR_ADVACKTIMEOUT       = 0x4000
	DMLERR_DATAACKTIMEOUT      = 0x4002
	DMLERR_DLL_NOT_INITIALIZED = 0x4003
	DMLERR_EXECACKTIMEOUT      = 0x4006
	DMLERR_NO_CONV_ESTABLISHED = 0x400A // (0x400A)
	DMLERR_POKEACKTIMEOUT      = 0x400B
	DMLERR_POSTMSG_FAILED      = 0x400C
	DMLERR_SERVER_DIED         = 0x400E

	CF_TEXT         = 1
	CF_BITMAP       = 2
	CF_METAFILEPICT = 3
	CF_SYLK         = 4
	CF_DIF          = 5
	CF_TIFF         = 6
	CF_OEMTEXT      = 7
	CF_DIB          = 8
	CF_PALETTE      = 9
	CF_PENDATA      = 10
	CF_RIFF         = 11
	CF_WAVE         = 12
	CF_UNICODETEXT  = 13
	CF_ENHMETAFILE  = 14
	CF_HDROP        = 15
	CF_LOCALE       = 16
	CF_DIBV5        = 17
	CF_MAX          = 18

	DDE_FACK          = 0x8000
	DDE_FBUSY         = 0x4000
	DDE_FDEFERUPD     = 0x4000
	DDE_FACKREQ       = 0x8000
	DDE_FRELEASE      = 0x2000
	DDE_FREQUESTED    = 0x1000
	DDE_FAPPSTATUS    = 0x00FF
	DDE_FNOTPROCESSED = 0x0000

	//DDE_FACKRESERVED = ~(DDE_FACK | DDE_FBUSY | DDE_FAPPSTATUS)
	//DDE_FADVRESERVED = ~(DDE_FACKREQ | DDE_FDEFERUPD)
	//DDE_FDATRESERVED = ~(DDE_FACKREQ | DDE_FRELEASE | DDE_FREQUESTED)
	//DDE_FPOKRESERVED = ~(DDE_FRELEASE)

	XTYPF_NOBLOCK = 0x0002
	XTYPF_NODATA  = 0x0004
	XTYPF_ACKREQ  = 0x0008

	XCLASS_MASK         = 0xFC00
	XCLASS_BOOL         = 0x1000
	XCLASS_DATA         = 0x2000
	XCLASS_FLAGS        = 0x4000
	XCLASS_NOTIFICATION = 0x8000

	XTYP_ERROR           = UINT(0x0000 | XCLASS_NOTIFICATION | XTYPF_NOBLOCK)
	XTYP_ADVDATA         = UINT(0x0010 | XCLASS_FLAGS)
	XTYP_ADVREQ          = UINT(0x0020 | XCLASS_DATA | XTYPF_NOBLOCK)
	XTYP_ADVSTART        = UINT(0x0030 | XCLASS_BOOL)
	XTYP_ADVSTOP         = UINT(0x0040 | XCLASS_NOTIFICATION)
	XTYP_EXECUTE         = UINT(0x0050 | XCLASS_FLAGS)
	XTYP_CONNECT         = UINT(0x0060 | XCLASS_BOOL | XTYPF_NOBLOCK)
	XTYP_CONNECT_CONFIRM = UINT(0x0070 | XCLASS_NOTIFICATION | XTYPF_NOBLOCK)
	XTYP_XACT_COMPLETE   = UINT(0x0080 | XCLASS_NOTIFICATION)
	XTYP_POKE            = UINT(0x0090 | XCLASS_FLAGS)
	XTYP_REGISTER        = UINT(0x00A0 | XCLASS_NOTIFICATION | XTYPF_NOBLOCK)
	XTYP_REQUEST         = UINT(0x00B0 | XCLASS_DATA)
	XTYP_DISCONNECT      = UINT(0x00C0 | XCLASS_NOTIFICATION | XTYPF_NOBLOCK)
	XTYP_UNREGISTER      = UINT(0x00D0 | XCLASS_NOTIFICATION | XTYPF_NOBLOCK)
	XTYP_WILDCONNECT     = UINT(0x00E0 | XCLASS_DATA | XTYPF_NOBLOCK)
	XTYP_MONITOR         = UINT(0x00F0 | XCLASS_NOTIFICATION | XTYPF_NOBLOCK)

	XTYP_MASK  = UINT(0x00F0)
	XTYP_SHIFT = UINT(4)

	TIMEOUT_ASYNC = 0xFFFFFFFF

	APPCLASS_STANDARD = 0x00000000
	APPCMD_CLIENTONLY = 0x00000010

	CP_WINANSI    = 1004
	CP_WINUNICODE = 1200

	DNS_REGISTER   = 0x0001
	DNS_UNREGISTER = 0x0002
	DNS_FILTERON   = 0x0004
	DNS_FILTEROFF  = 0x0008
	NULL           = 0
	FALSE          = 0
	TRUE           = 1
)

type (
	CString                        byte
	UString                        uint16
	OVString                       string
	VString                        string
	Fake_type_Fix_me               uintptr
	enum                           int
	ACCESS_MASK                    DWORD
	AChar                          Char
	ALG_ID                         Unsigned_int
	AString                        *string
	ATOM                           WORD
	BOOL                           int
	BOOLEAN                        byte
	BSTR                           *OLECHAR
	BYTE                           byte
	CALID                          DWORD
	CALTYPE                        DWORD
	CH                             AChar
	CLIPFORMAT                     WORD
	COLOR16                        USHORT
	COLORREF                       DWORD
	Char                           uint8
	DATE                           DOUBLE
	DISPID                         int32
	DOUBLE                         float64
	DWORD                          uint32 //TODO(t):True w32/w64?
	DWORDLONG                      uint64 //TODO(t):True w32/w64?
	DWORD_PTR                      ULONG_PTR
	EXECUTION_STATE                DWORD
	FARPROC                        uintptr //TODO(t):unsafe.Pointer?
	FLOAT                          float32
	FOURCC                         DWORD
	FXPT2DOT30                     LONG
	GEOCLASS                       DWORD
	GEOID                          LONG
	GEOTYPE                        DWORD
	GROUP                          uint
	HACCEL                         HANDLE
	HANDLE                         DWORD //TODO(t):Actually strict -> struct but go does not do coercion
	HBITMAP                        int   //TODO(t):HANDLE but go does not do coercion
	HBRUSH                         HANDLE
	HCOLORSPACE                    HANDLE
	HCONV                          HANDLE
	HCONVLIST                      HANDLE
	HCRYPTHASH                     ULONG_PTR
	HCRYPTKEY                      ULONG_PTR
	HCRYPTPROV                     ULONG_PTR
	HCURSOR                        HANDLE
	HDC                            HANDLE
	HDDEDATA                       int
	HDESK                          HANDLE
	HDEVNOTIFY                     HANDLE
	HDPA                           HANDLE
	HDROP                          HANDLE
	HDRVR                          HANDLE
	HDSA                           HANDLE
	HDWP                           HANDLE
	HENHMETAFILE                   HANDLE
	HFILE                          HANDLE
	HFONT                          HANDLE
	HGDIOBJ                        HANDLE
	HGLOBAL                        HANDLE
	HGLRC                          HANDLE
	HHOOK                          HANDLE
	HICON                          HANDLE
	HIMAGELIST                     HANDLE
	HINSTANCE                      HANDLE
	HKEY                           HANDLE
	HKL                            HANDLE
	HLOCAL                         HANDLE
	HMENU                          HANDLE
	HMETAFILE                      HANDLE
	HMIDI                          HANDLE
	HMIDIIN                        HANDLE
	HMIDIOUT                       HANDLE
	HMIDISTRM                      HANDLE
	HMIXER                         HANDLE
	HMIXEROBJ                      HANDLE
	HMMIO                          HANDLE
	HMODULE                        HINSTANCE
	HMONITOR                       HANDLE
	HOLEMENU                       HANDLE
	HPALETTE                       HANDLE
	HPEN                           HANDLE
	HRAWINPUT                      HANDLE
	HRESULT                        HANDLE
	HRGN                           HANDLE
	HRSRC                          HANDLE
	HSZ                            HANDLE
	HTASK                          HANDLE
	HWAVEIN                        HANDLE
	HWAVEOUT                       HANDLE
	HWINEVENTHOOK                  HANDLE
	HWINSTA                        HANDLE
	HWND                           HANDLE
	IAdviseSink                    struct{}
	IBindCtx                       struct{}
	IBindStatusCallback            struct{}
	IChannelHook                   struct{}
	IClassFactory                  struct{}
	ICreateErrorInfo               struct{}
	ICreateTypeLib                 struct{}
	ICreateTypeLib2                struct{}
	IDataAdviseHolder              struct{}
	IDataObject                    struct{}
	IDispatch                      struct{}
	IDropSource                    struct{}
	IDropTarget                    struct{}
	IEnumFORMATETC                 struct{}
	IEnumOLEVERB                   struct{}
	IErrorInfo                     struct{}
	IFillLockBytes                 struct{}
	IInitializeSpy                 struct{}
	ILockBytes                     struct{}
	IMalloc                        struct{}
	IMallocSpy                     struct{}
	IMarshal                       struct{}
	IMessageFilter                 struct{}
	IMoniker                       struct{}
	INT                            int
	INT_PTR                        INT
	IOAString                      *string
	IOleAdviseHolder               struct{}
	IOleClientSite                 struct{}
	IOleInPlaceActiveObject        struct{}
	IOleInPlaceFrame               struct{}
	IOleObject                     struct{}
	IPersistStorage                struct{}
	IPersistStream                 struct{}
	IRecordInfo                    struct{}
	IRunningObjectTable            struct{}
	IStorage                       struct{}
	IStream                        struct{}
	ISurrogate                     struct{}
	ITypeInfo                      struct{}
	ITypeLib                       struct{}
	IUnknown                       struct{}
	LANGID                         USHORT
	LCID                           DWORD
	LCSCSTYPE                      LONG
	LCSGAMUTMATCH                  LONG
	LCTYPE                         DWORD
	LGRPID                         DWORD
	LONG                           int32 //TODO(t):?SIZE __3264
	LONG64                         int64
	LONGLONG                       int64 //TODO(t):Win64=128
	LONG_PTR                       LONG
	LPARAM                         LONG_PTR
	LRESULT                        LONG_PTR
	MCIDEVICEID                    UINT
	MCIERROR                       DWORD
	MENUTEMPLATE                   *VOID
	MMRESULT                       UINT
	MMVERSION                      UINT
	MVString                       *string //TODO(t):*[]string?
	NLS_FUNCTION                   DWORD
	OAString                       *string
	OLECHAR                        WChar
	OLESTR                         WString
	OLESTREAM                      struct{}
	OMVString                      *string //TODO(t):*[]string ?
	OWString                       *string
	REGSAM                         ACCESS_MASK
	RPC_AUTHZ_HANDLE               *VOID
	RPC_AUTH_IDENTITY_HANDLE       *VOID
	SCODE                          int32
	SC_HANDLE                      HANDLE
	SC_LOCK                        *VOID
	SECURITY_CONTEXT_TRACKING_MODE BOOLEAN
	SECURITY_DESCRIPTOR_CONTROL    USHORT
	SECURITY_INFORMATION           DWORD
	SERVICETYPE                    ULONG
	SERVICE_STATUS_HANDLE          HANDLE
	SHORT                          int16
	SIZE_T                         ULONG_PTR
	SNB                            **OLECHAR
	SOCKET                         UINT_PTR
	STREAM                         IStream
	UCHAR                          byte
	UINT                           uint
	UINT_PTR                       UINT
	ULONG                          DWORD //TODO(t):size
	ULONG64                        uint64
	ULONGLONG                      uint64  //TODO(t):Win64=128?
	ULONG_PTR                      uintptr //TODO(t):true def
	USHORT                         uint16  //TODO(t):size
	U_int                          uint
	U_long                         uint32
	U_short                        uint16
	Unsigned_int                   uint
	Unsigned_long                  uint32
	Unsigned_short                 uint16
	VARIANT_BOOL                   int16
	VARTYPE                        uint16
	VChar                          *uint16  //TODO(t): uint8/uint16
	VOID                           struct{} // //TODO(t):Go does not do coercion; Any side-effects?
	VOID64                         struct{} //TODO(t):__ptr64 //TODO(t):Go does not do coercion; Any side-effects?
	WChar                          uint16
	WORD                           uint16 //TODO(t):True 32/64? signed?
	WPARAM                         LONG_PTR
	WSAEVENT                       HANDLE
	WString                        *string
)

type FNCALLBACK func(
	Type UINT,
	Fmt UINT,
	Conv HCONV,
	Sz1 HSZ,
	Sz2 HSZ,
	Data HDDEDATA,
	Data1 ULONG_PTR,
	Data2 ULONG_PTR) HDDEDATA

type SECURITY_IMPERSONATION_LEVEL enum

type SECURITY_QUALITY_OF_SERVICE struct {
	Length              DWORD
	ImpersonationLevel  SECURITY_IMPERSONATION_LEVEL
	ContextTrackingMode SECURITY_CONTEXT_TRACKING_MODE
	EffectiveOnly       BOOLEAN
}
type CONVCONTEXT struct {
	Size      UINT
	Flags     UINT
	CountryID UINT
	CodePage  int
	LangID    DWORD
	Security  DWORD
	QoS       SECURITY_QUALITY_OF_SERVICE
}
type CONVINFO struct {
	Size        DWORD
	User        DWORD_PTR
	ConvPartner HCONV
	SvcPartner  HSZ
	ServiceReq  HSZ
	Topic       HSZ
	Item        HSZ
	Fmt         UINT
	Type        UINT
	Status      UINT
	Convst      UINT
	LastError   UINT
	ConvList    HCONVLIST
	ConvCtxt    CONVCONTEXT
	Wnd         HWND
	WndPartner  HWND
}
type TagPOINT struct {
	X LONG
	Y LONG
}

type TagMSG struct {
	Hwnd    HWND
	Message UINT
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      TagPOINT
}

type LPMSG *TagMSG
