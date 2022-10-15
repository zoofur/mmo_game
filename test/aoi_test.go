package test

import (
	"fmt"
	"github.com/lorenzoyu2000/mmo_game/core"
	"testing"
)

func TestNewAOIMgr(t *testing.T) {
	aoiMgr := core.NewAOIMgr(0, 250, 5, 0, 250, 5)
	fmt.Println(aoiMgr)
}

func TestGetSurround(t *testing.T) {
	aoiMgr := core.NewAOIMgr(0, 250, 5, 0, 250, 5)
	for gid, _ := range aoiMgr.Grids {
		grids := aoiMgr.GetSurroundGrid(gid)
		gids := make([]int, 0, len(grids))
		for _, v := range grids {
			gids = append(gids, v.GID)
		}
		fmt.Println("GID = ", gid, "Surround Grids = ", gids)
	}
}
