package main

import (
	"fmt"
	"github.com/lorenzoyu2000/mmo_game/apis"
	"github.com/lorenzoyu2000/mmo_game/core"
	"github.com/lorenzoyu2000/zinx/ziface"
	"github.com/lorenzoyu2000/zinx/znet"
)

func main() {
	s := znet.NewServer()
	// 注册钩子函数
	s.SetOnConnCreate(onConnCreate)
	s.SetOnConnDestroy(onConnDestroy)
	// 注册处理函数
	s.AddRouter(2, &apis.WorldChat{})
	s.AddRouter(3, &apis.MoveApi{})
	s.Serve()
}

func onConnCreate(conn ziface.IConnection) {
	// 初始化玩家
	player := core.NewPlayer(conn)
	// 向客户端发送MsgID = 1， 同步玩家ID
	player.SyncPid()
	// 向客户端发送MsgID = 200， 同步玩家位置
	player.BroadCastStartPosition()
	// 将当前新上线玩家添加到世界中
	core.WorldMgr.AddPlayer(player)
	// 向conn中添加属性pid，来提供给WorldChat获取
	conn.SetProperty("pid", player.Pid)
	// 同步玩家位置信息
	player.SyncSurroundPlayer()

	fmt.Println("Player ID: ", player.Pid, " is arrived")
}

func onConnDestroy(conn ziface.IConnection) {
	pid, err := conn.GetProperty("pid")
	if err != nil {
		fmt.Println("conn GetProperty err: ", err)
		return
	}

	player := core.WorldMgr.Players[pid.(int32)]
	player.Offline()
	fmt.Println("Player ID: ", player.Pid, " is offline")
}
