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
