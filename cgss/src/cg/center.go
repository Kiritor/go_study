/**
  聊天中央服务器
*/

package cg

import (
   "encoding/json"
   "errors"
   "sync"
   "ipc"
)

var _ ipc.Server = &CenterServer{}    //确认实现了server接口


/**
   消息体
*/
type  Message struct {
    From string "from"
	To string "to"
	Content string "content"
}

/**
   中央服务器
*/
type CenterServer struct {
    servers map[string] ipc.Server
	players []*Player
	rooms []*Room
	mutex sync.RWMutex
}

//建立一个中央服务器
func NewCenterServer() *CenterServer {
   	servers :=make(map[string] ipc.Server)
	players :=make([]*Player,0)
	return &CenterServer{servers:servers,players:players}
}

//添加一个用户
func (server *CenterServer) addPlayer(params string) error {
	player :=NewPlayer()
	err :=json.Unmarshal([]byte(params),&player)
	if err!=nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()
	//偷懒了,没做重复登录检查
	server.players = append(server.players,player)
	return nil
}

//删除一个用户
func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Lock()
	for i,v :=range server.players  {
		if v.Name == params {
			if len(server.players) ==1 {
				server.players = make([]*Player,0)
			}else if i == len(server.players) -1 {
				server.players = server.players[:i-1]
			}else if i==0 {
				server.players = server.players[i:]
			}else {
				server.players = append(server.players[:i-1],server.players[i+1:]...)
			}
			return nil
		}
	}
	return errors.New("Player not found.")
}

//列出在线的用户
func (server *CenterServer) listPlayer() (players string,err error) {
	server.mutex.RLock()
	defer server.mutex.RUnlock()
	if len(server.players) > 0 {
		b,_:=json.Marshal(server.players)
		players = string(b)
	}else {
		err = errors.New("No player online")
	}
	return
}

//广播
func (server *CenterServer) broadcast(params string) error {
	var message Message
	    err:=json.Unmarshal([]byte(params),&message)
	if err!= nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) >0 {
		for _,player:= range server.players {
			player.mq <- &message
		}
	}else {
		err = errors.New("No player online")
	}
	return err
}

func (server *CenterServer) Handle(method,params string) *ipc.Response {
	switch method {
		case "addPlayer":
		     err:=server.addPlayer(params)
		     if err!=nil {
				 return &ipc.Response{Code:err.Error()}
			 }
		case "removePlayer":
		     err:=server.removePlayer(params)
		     if err!=nil {
				 return &ipc.Response{Code:err.Error()}
			 }
		case "listPlayer":
		     players,err :=server.listPlayer()
		     if err !=nil {
				 return &ipc.Response{Code:err.Error()}
			 }
		     return &ipc.Response{Code:"200",Body:players}
		case "broadCast":
		     err:=server.broadcast(params)
		     if err!=nil {
				 return &ipc.Response{Code:err.Error()}
			 }
		     return &ipc.Response{Code:"200"}
		default:
		     return &ipc.Response{"404",method+":"+params}    //模拟http 404错误
	}
	return &ipc.Response{Code:"200"}
}

//实现interface的Name方法
func (server *CenterServer) Name() string {
	return "CenterServer"
}
