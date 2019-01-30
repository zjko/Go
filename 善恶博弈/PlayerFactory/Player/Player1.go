package player

// Player1 策略1
// 无论是谁都将无条件的合作
type Player1 struct{
	BasicPlayer
}

// Init 类初始化
func (p * Player1) Init () bool {
	p.name = "老实人"
	return true
}

// Game 实现博弈接口
func (p * Player1) Game (id int) bool {
	return true
}

// Process 博弈完成之后要做的处理
func (p * Player1) Process (res bool, id int) bool {
	return true
}