package core

import (
	"fmt"
)

/*
	AOI算法管理模块
*/
type AOIManager struct {
	// 区域左边界坐标
	MinX int
	// 区域右边界坐标
	MaxX int
	// x轴格子数量
	CntsX int
	// 区域上边界坐标
	MaxY int
	// 区域下边界坐标
	MinY int
	// y轴格子数量
	CntsY int
	// 当前区域中有哪些格子
	Grids map[int]*Grid
}

func NewAOIMgr(minX, maxX, cntsX, maxY, minY, cntsY int) *AOIManager {
	aoi := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MaxY:  maxY,
		MinY:  minY,
		CntsY: cntsY,
		Grids: make(map[int]*Grid),
	}

	// 初始化AOI区域的所有格子和编号
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			gid := y*cntsX + x

			aoi.Grids[gid] = NewGrid(
				gid,
				aoi.MinX+x*aoi.gridWidth(),
				aoi.MinX+(x+1)*aoi.gridWidth(),
				aoi.MinY+y*aoi.gridLength(),
				aoi.MinY+(y+1)*aoi.gridLength(),
			)
		}
	}

	return aoi
}

// 获取x轴格子宽度
func (a *AOIManager) gridWidth() int {
	return (a.MaxX - a.MinX) / a.CntsX
}

// 获取y轴格子长度
func (a *AOIManager) gridLength() int {
	return (a.MaxY - a.MinY) / a.CntsY
}

// 打印格子信息
func (a *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager:\nMinX: %d, MaxX: %d, CntsX: %d, MinY: %d, MaxX: %d, CntsY: %d\nGrids:\n",
		a.MinX, a.MaxX, a.CntsX, a.MinX, a.MaxY, a.CntsY)
	for _, grid := range a.Grids {
		s += fmt.Sprintln(grid)
	}

	return s
}

// 获取九宫格
func (a *AOIManager) GetSurroundGrid(gridID int) (grids []*Grid) {
	// 不存在当前格子，就直接退出
	if _, ok := a.Grids[gridID]; !ok {
		return
	}

	// 获取gridID所在x轴的格子
	grids = append(grids, a.Grids[gridID])
	idX := gridID % a.CntsX
	if idX > 0 {
		grids = append(grids, a.Grids[gridID-1])
	}
	if idX < a.CntsX-1 {
		grids = append(grids, a.Grids[gridID+1])
	}

	// 根据gridID所在x轴的格子来获取y轴格子
	indexX := make([]int, 0, len(grids))
	for _, v := range grids {
		indexX = append(indexX, v.GID)
	}
	for _, v := range indexX {
		idY := v / a.CntsY
		if idY > 0 {
			grids = append(grids, a.Grids[v-a.CntsX])
		}
		if idY < a.CntsY-1 {
			grids = append(grids, a.Grids[v+a.CntsY])
		}
	}

	return grids
}

// 通过坐标获取格子ID
func (a *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - a.MinX) / a.gridWidth()
	idy := (int(y) - a.MinY) / a.gridLength()
	return idy*a.CntsX + idx
}

// 通过坐标获取周边九宫格玩家信息
func (a *AOIManager) GetPidsByPos(x, y float32) (players []int) {
	gridID := a.GetGidByPos(x, y)
	grids := a.GetSurroundGrid(gridID)

	for _, g := range grids {
		for p, _ := range g.players {
			players = append(players, p)
		}
	}
	fmt.Printf("x : %v, y: %v, surround players: %v\n", x, y, players)

	return
}

// 添加一个PlayerID到一个格子中
func (a *AOIManager) AddPidToGrid(pid, gid int) {
	a.Grids[gid].Add(pid)
}

// 移除一个格子中的PlayerID
func (a *AOIManager) RemovePidFromGrid(pid, gid int) {
	a.Grids[gid].Remove(pid)
}

// 通过GID获取全部的PlayerID
func (a *AOIManager) GetPidByGid(gid int) (players []int) {
	players = a.Grids[gid].GetAllPlayersFromGrid()

	return
}

// 通过坐标将Player添加到一个格子中
func (a *AOIManager) AddPidByPos(x, y float32, pid int) {
	gid := a.GetGidByPos(x, y)
	a.Grids[gid].Add(pid)
}

// 通过坐标把一个Player从格子中删除
func (a *AOIManager) RemovePidFromGridByPos(x, y float32, pid int) {
	gid := a.GetGidByPos(x, y)
	a.Grids[gid].Remove(pid)
}
