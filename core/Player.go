package core

import (
	"fmt"
	"github.com/lorenzoyu2000/mmo_game/pb"
	"github.com/lorenzoyu2000/zinx/ziface"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"sync"
)

/*
	玩家管理模块
*/
type Player struct {
	// 玩家ID
	Pid int32
	// 连接
	Conn ziface.IConnection
	// 平面x轴坐标
	X float32
	// 高度
	Y float32
	// 平面y坐标，注意不是Y
	Z float32
	// 玩家角度，0-360度
	V float32
}

// 玩家ID生成器
var PidGenerate int32 = 1

// 玩家ID生成器的保护锁
var PidLock sync.Mutex

func NewPlayer(conn ziface.IConnection) *Player {
	PidLock.Lock()
	id := PidGenerate
	PidGenerate++
	PidLock.Unlock()

	return &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), // 随机在X轴160处，随机偏移
		Y:    0,                            // 高度为0
		Z:    float32(120 + rand.Intn(20)), // 基于y轴120处，随机偏移
		V:    0,                            // 角度为0
	}
}

// 玩家和客户端通信
func (p *Player) SendMsg(msgID uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("data marshal err: ", err)
		return
	}

	if p.Conn == nil {
		fmt.Println("Player", p.Pid, " Connection is nil")
		return
	}

	err = p.Conn.Send(msgID, msg)
	if err != nil {
		fmt.Println("Send Msg err: ", err)
	}
}

// 同步玩家ID
func (p *Player) SyncPid() {
	msg := &pb.SyncPid{
		Pid: p.Pid,
	}

	p.SendMsg(1, msg)
}

// 同步玩家位置
func (p *Player) BroadCastStartPosition() {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)
}

// 世界聊天广播
func (p *Player) Talk(content string) {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	players := WorldMgr.GetAllPlayers()
	for _, v := range players {
		v.SendMsg(200, msg)
	}
}
