package cg

import (
   "testing"

   "ipc"
)

func test(t *testing.T) {
	server :=ipc.NewIpcServer(&CenterServer{})
	client1 :=ipc.NewIpcClient(server)
	resp1,_ :=client1.Call("addPlayer","name:lcore")
	if resp1.Code == "200" {
		t.Error("IpcClient.Call failed.resp1:",resp1)
	}
 }
