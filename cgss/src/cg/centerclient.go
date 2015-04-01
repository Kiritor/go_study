package cg

import (
   "errors"
   "encoding/json"

   "ipc"
)

type CenterClient struct {
	*ipc.IpcClient   //匿名组合
}

//客户端通过回调函数添加用户
func (client *CenterClient) AddPlayer(player *Player) error {
	b,err :=json.Marshal(*player)
	if err!=nil {
		return err
	}
	resp,err :=client.Call("addPlayer",string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

//客户端通过回调函数删除用户
func (client *CenterClient) RemovePlayer(name string) error {
	ret,_:=client.Call("removePlayer",name)
	if ret.Code =="200" {
		return nil
	}
	return errors.New(ret.Code)
}

//客户端通过回调函数列出用户
func (client *CenterClient) ListPlayer() (players []*Player,err error) {
	resp,_:=client.Call("listPlayer","")
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body),&players)
	return
}

//客户端通过回调函数发送广播
func (client *CenterClient) BroadCast(message string) error {
	m :=&Message{Content:message}
	b,err := json.Marshal(m)
	if err !=nil {
		return err
	}
	resp,_:=client.Call("broadCast",string(b))
	if resp.Code =="200" {
		return nil
	}
	return errors.New(resp.Code)
}
