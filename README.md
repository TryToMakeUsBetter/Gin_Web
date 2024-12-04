# Gin
## gin.Default
返回Gin的type Engine struct{}，里面包含RouterGroup。

router.GET
router.POST
router.PUT
router.DELETE

gin.H{}

# net/http
## http.Server
```go
type Server struct{
    Addr string //监听的TCP地址，格式为:8000
    Handler Handler //http句柄，处理响应HTTP
    TLSConfig *tls.Config //TLS配置
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout time.Duration
    IdleTimeout time.Duration
    MaxHeaderBytes int
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```

### ListenAndServe()
```go
unc (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```

# ORM