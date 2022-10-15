package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	// 格子ID
	GID int
	// 格子左边界
	MinX int
	// 格子右边界
	MaxX int
	// 格子上边界
	MinY int
	// 格子下边界
	MaxY int
	// 附近玩家集合
	players map[int]bool
	// 保护map的锁
	playerLock sync.RWMutex
}

func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:     gid,
		MinX:    minX,
		MaxX:    maxX,
		MinY:    minY,
		MaxY:    maxY,
		players: make(map[int]bool),
	}
}

// 添加玩家
func (g *Grid) Add(playerID int) {
	g.playerLock.Lock()
	defer g.playerLock.Unlock()

	g.players[playerID] = true
}

// 移除玩家
func (g *Grid) Remove(playerID int) {
	g.playerLock.Lock()
	defer g.playerLock.Unlock()

	delete(g.players, playerID)
}

// 获取格子内的所有玩家
func (g *Grid) GetAllPlayersFromGrid() (players []int) {
	g.playerLock.RLock()
	defer g.playerLock.RUnlock()

	for id, _ := range g.players {
		players = append(players, id)
	}

	return
}

// 重写String()方法，打印格子信息
func (g *Grid) String() string {
	return fmt.Sprintf("GID: %d, MinX: %d, MaxX: %d, MinY: %d, MaxY: %d, Players: %v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.players)
}
