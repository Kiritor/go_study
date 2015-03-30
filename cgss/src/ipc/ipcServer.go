/**
  @Author:LCore
   简单ipc(进程间通信),封装通信包的编码细节
*/
package ipc

import (
   "encoding/json"
   "fmt"
)

type Request struct {
   Method string "method"   //方法
   Params string "params"   //参数
}

type 　Response struct {
   Code string "code"
   Body string "body"
}

type Server interface {
  Name() string
  Handle(method,params string) *Response
}

type IpcServer struct {
   Server
}

func NewIpcServer(server Server) *IpcServer {
   return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
   session := make(chan string,0)
	go func(c chan string) {
	   for {
	      request :=<-c
		  if request == "CLOSE" {
		     //关闭连接
			 break
		  }
		  var req Request
		  err :=json.Unmarshal([]byte(request),&req)
		  if err!=nil {
		     fmt.Println("request格式错误:",request)
		  }
		  resp :=server.Hanlder(req.Method,req.Params)
		  b, err :=json.Marshal(resp)
		  c<-string(b)  //返回结果
	   }
	   fmt.Println("session closed.")
	}(session)
	fmt.Println("A new Session has been create!")
	return session
}
