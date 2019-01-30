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
	case "Player9": 
		product = new(player.Player9)
	default: 
		fmt.Println("undefine:",name)
		return nil	
	}
	product.Init()
	return product
}