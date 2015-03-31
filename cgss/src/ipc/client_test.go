package ipc

import (
   "testing"
)

type EchoServer struct {

}

func (server *EchoServer) Handle(request ,params string) *Response {
	return  &Response{"200","<b>success</b>"}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server :=NewIpcServer(&EchoServer{})
	client1 :=NewIpcClient(server)
	client2 :=NewIpcClient(server)

	resp1,_ :=client1.Call("post","name:lcore")
	resp2,_ :=client2.Call("post","name:liangtao")
	if resp1.Code !="200" ||resp2.Code!="200" {
		t.Error("IpcClient.Call failed.resp1:",resp1,"  resp2:",resp2)
	}
	client1.Close()
	client2.Close()
}
