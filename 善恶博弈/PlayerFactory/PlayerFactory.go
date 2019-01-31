package playerFactory

import (
	"./Player"
	
	"./Strategy"
	"fmt"
)
// PlayerFactory 简单策略工厂
type PlayerFactory struct {

}

// Create 创建一个对象
func (pf * PlayerFactory) Create(name string) strategy.IStrategy {
	var product strategy.IStrategy
	switch name {
	case "Player1": 
		product = new(player.Player1)
	case "Player2": 
		product = new(player.Player2)
	case "Player3": 
		product = new(player.Player3)
	case "Player4": 
		product = new(player.Player4)
	case "Player5": 
		product = new(player.Player5)
	case "Player6": 
		product = new(player.Player6)
	case "Player7": 
		product = new(player.Player7)
	case "Player8": 
		product = new(player.Player8)
	case "Player9": 
		product = new(player.Player9)
	case "Player10": 
		product = new(player.Player10)
	case "Player11": 
		product = new(player.Player11)
	default: 
		fmt.Println("undefine:",name)
		return nil	
	}
	product.Init()
	return product
}