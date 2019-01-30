package player

// Player2 策略2
// 针锋相对策略 合作与背叛，取决于上一次对方的选择
type Player2 struct{
	BasicPlayer
}
// Init 类初始化
func (p * Player2) Init () bool {
	p.name = "针锋相对"
	return true
}

// Game 实现博弈接口
func (p * Player2) Game (id int) bool {
	return true
}

// Process 博弈完成之后要做的处理
func (p * Player2) Process (res bool, id int) bool {
	return true
}