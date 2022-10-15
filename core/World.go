package core

import "sync"

/*
	世界管理模块
*/
type WorldManager struct {
	AoiMgr  *AOIManager
	Players map[int32]*Player
	WMLock  sync.RWMutex
}

var WorldMgr *WorldManager

func init() {
	WorldMgr = &WorldManager{
		AoiMgr:  NewAOIMgr(0, 250, 5, 0, 250, 5),
		Players: make(map[int32]*Player),
	}
}

func (w *WorldManager) AddPlayer(player *Player) {
	w.WMLock.Lock()
	w.Players[player.Pid] = player
	w.WMLock.Unlock()

	w.AoiMgr.AddPidByPos(player.X, player.Z, int(player.Pid))
}

func (w *WorldManager) RemovePlayerByPid(pid int32) {
	player := w.Players[pid]
	w.AoiMgr.RemovePidFromGridByPos(player.X, player.Z, int(player.Pid))

	w.WMLock.Lock()
	delete(w.Players, pid)
	w.WMLock.Unlock()
}

func (w *WorldManager) GetPlayerByPid(pid int32) *Player {
	w.WMLock.RLock()
	defer w.WMLock.RUnlock()
	return w.Players[pid]
}

func (w *WorldManager) GetAllPlayers() (players []*Player) {
	w.WMLock.RLock()
	defer w.WMLock.RUnlock()

	for _, v := range w.Players {
		players = append(players, v)
	}
	return
}
