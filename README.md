# mmo_game

## introduce

mmo_game是一个可基于[zinx](https://github.com/lorenzoyu2000/zinx)进行多人联机的游戏客户端，提供以下功能：

- AOI算法：对游戏地图划分格子，从而显示指定范围内的玩家，避免通知全世界玩家带来的性能损耗。
- 玩家功能：根据坐标提供上线显示、玩家移动，世界聊天。

![AOI算法](https://imgs-1306864474.cos.ap-beijing.myqcloud.com/img/AOI%E7%AE%97%E6%B3%95.jpg)

## quick start

启动服务端：

```go
cd mmo_game
go run main.go
```

启动客户端：

```go
cd mmo_game/client
start client.exe
```



