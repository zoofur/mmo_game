package apis

import (
	"fmt"
	"github.com/lorenzoyu2000/mmo_game/core"
	"github.com/lorenzoyu2000/mmo_game/pb"
	"github.com/lorenzoyu2000/zinx/ziface"
	"github.com/lorenzoyu2000/zinx/znet"
	"google.golang.org/protobuf/proto"
)

/*
	玩家移动
*/
type MoveApi struct {
	znet.BaseRouter
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	pos := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), pos)
	if err != nil {
		fmt.Println("Move unmarshal err: ", err)
		return
	}

	// 获取玩家pid
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("Move GetProperty err: ", err)
		return
	}
	fmt.Printf("user pid: %d, x: %f, y: %f, z: %f, v:%f\n", pid, pos.X, pos.Y, pos.Z, pos.V)
	player := core.WorldMgr.Players[pid.(int32)]
	player.UpdateAOI(pos.X, pos.Y, pos.Z, pos.V)
	player.UpdatePos(pos.X, pos.Y, pos.Z, pos.V)
}
