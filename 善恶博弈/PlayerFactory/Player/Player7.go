package player

// Player7 策略1
// 无论是谁都将无条件的合作
type Player7 struct{
	BasicPlayer
}

// Init 类初始化
func (p * Player7) Init () bool {
	p.name = "报复社会"
	return true
}

// Game 实现博弈接口
func (p * Player7) Game (id int) bool {
	return false
}

// Process 博弈完成之后要做的处理
func (p * Player7) Process (res bool, id int) bool {
	return true
}