package main
/*
* 	2048小游戏 go语言版
*	输入： wsad 分别表示上下左右 回车键确认
*/
import (
	"./Game"
)

func main() {
	g := new(Game.Game)
	g.Start()
}
