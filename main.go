package main

import (
	"fmt"
	"github.com/lorenzoyu2000/mmo_game/core"
	"github.com/lorenzoyu2000/zinx/ziface"
	"github.com/lorenzoyu2000/zinx/znet"
)

func main() {
	s := znet.NewServer()
	s.SetOnConnCreate(OnConnCreate)
	s.Serve()
}

func OnConnCreate(conn ziface.IConnection) {
	// 初始化玩家
	player := core.NewPlayer(conn)
	// 向客户端发送MsgID = 1， 同步玩家ID
	player.SyncPid()
	// 向客户端发送MsgID = 200， 同步玩家位置
	player.BroadCastStartPosition()

	fmt.Println("Player ID: ", player.Pid, " is arrived")
}
