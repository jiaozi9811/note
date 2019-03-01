# net/http

tags : golang
---
#客户端
##http.Get
>func Get(url string) (resp *Response, err error)
```
res, err := http.Get("http://www.google.com/robots.txt")
if err != nil {     log.Fatal(err) }
robots, err := ioutil.ReadAll(res.Body)
res.Body.Close()
if err != nil {    log.Fatal(err) }
fmt.Printf("%s", robots)
```
##type Response
```
type Response struct {
	Header Header   //保管头域的键值对
	Body io.ReadCloser  //代表回复的主体
	ContentLength int64 //记录相关内容的长度
	Close bool  // 记录头域是否指定应在读取完主体后关闭连接。（即Connection头）
	Request *Request    // Request是用来获取此回复的请求
	
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	TransferEncoding []string   //按从最外到最里的顺序列出传输编码
  	Uncompressed bool  //是否压缩
	Trailer Header  // 保存和头域相同格式的trailer键值对，和Header字段相同类型
	TLS *tls.ConnectionState    // 包含接收到该回复的TLS连接的信息
```
##type Request
```
type Request struct {
    Method string   // 指定HTTP方法(GET,POST,PUT等).对客户端,""代表GET
    URL *url.URL    // 在服务端表示被请求的URI,在客户端表示要访问的URL,在客户端,URL的Host字段指定了要连接的服务器
    Header Header   // Header字段用来表示HTTP请求的头域
    Body io.ReadCloser   // Body是请求的主体
    GetBody func() (io.ReadCloser, error)
    ContentLength int64    // 记录相关内容的长度
    Response *Response
    Close bool   // 在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接
    // 在服务端，Host指定URL会在其上寻找资源的主机。
    // Host的格式可以是"host:port"。
    // 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
    // 如过该字段为""，Request.Write方法会使用URL字段的Host。
    Host string
    RemoteAddr string   // 允许HTTP服务器和其他软件记录该请求的来源地址,一般用于日志;客户端会忽略
    RequestURI string  // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
    Cancel <-chan struct{}
    Form url.Values   // 是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据
    PostForm url.Values   // 解析好的POST或PUT的表单数据

    Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
    TransferEncoding []string  // 按从最外到最里的顺序列出传输编码
    // MultipartForm是解析好的多部件表单，包括上传的文件。
    // 本字段只有在调用ParseMultipartForm后才有效
    MultipartForm *multipart.Form
    Trailer Header    // 指定了会在请求主体之后发送的额外的头域
    TLS *tls.ConnectionState  // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
    ctx context.Context
}
```

##type Header
>type Header map[string][]string
##type Cookie
```
type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite SameSite
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
```
#服务端 server
##type Handler
```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
>实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务

##type Server struct
```
type Server struct {
	Addr    string  
	Handler Handler // handler to invoke, http.DefaultServeMux if nil
	ReadTimeout time.Duration// 请求的读取操作在超时前的最大持续时间
	WriteTimeout time.Duration // 回复的写入操作在超时前的最大持续时间
 	MaxHeaderBytes int   // 请求的头域最大长度，如为0则用DefaultMaxHeaderBytes
	mu         sync.Mutex

	TLSConfig *tls.Config// 可选的TLS配置，用于ListenAndServeTLS方法
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

	ConnState func(net.Conn, ConnState)
 	ErrorLog *log.Logger   // ErrorLog指定一个可选的日志记录器

	ReadHeaderTimeout time.Duration
	IdleTimeout time.Duration

	listeners  map[*net.Listener]struct{}
	activeConn map[*conn]struct{}
	doneChan   chan struct{}
	onShutdown []func()

	disableKeepAlives int32     // accessed atomically.
	inShutdown        int32     // accessed atomically (non-zero means we're in Shutdown)
	nextProtoOnce     sync.Once // guards setupHTTP2_* init
	nextProtoErr      error     // result of http2.ConfigureServer if used


}
```
##type serverHandler struct {
##type ServeMux
ServeMux类型是HTTP请求的多路转接器。它会将每一个接收的请求的URL与一个注册模式的列表进行匹配，并调用和URL最匹配的模式的处理器

```
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}
```
##type HandlerFunc func
HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Handler接口
>type HandlerFunc func(ResponseWriter, *Request)

##type conn struct
```
type conn struct {
	server *Server
	cancelCtx context.CancelFunc
	rwc net.Conn
	remoteAddr string
	tlsState *tls.ConnectionState
	werr error
	r *connReader
	bufr *bufio.Reader
	bufw *bufio.Writer
	lastMethod string
	curReq atomic.Value // of *response (which has a Request in it)
	curState struct{ atomic uint64 } // packed (unixtime<<8|uint8(ConnState))
	mu sync.Mutex
	hijackedv bool
}
```
##type ResponseWriter interface
>type ResponseWriter interface {

##type response struct {
##type Flusher interface {
##type ResponseWriter interface {
##type Hijacker interface {
##type CloseNotifier interface {
##type chunkWriter struct {
##type writerOnly struct {
##type readResult struct {
##type connReader struct {
##type expectContinueReader struct {
##type closeWriter interface {




<!--stackedit_data:
eyJoaXN0b3J5IjpbLTY3MTUxNTkxNF19
-->
