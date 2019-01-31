package player

// Player10 策略2
// 针锋相对策略 合作与背叛，取决于上一次对方的选择
type Player10 struct{
	BasicPlayer
	list map[int]bool
	
}
// Init 类初始化
func (p * Player10) Init () bool {
	p.name = "舔狗养成"
	p.list = make(map[int]bool)
	return true
}

// Game 实现博弈接口
func (p * Player10) Game (id int) bool {
	v,ok := p.list[id]
	if !ok {
		p.list[id] = false
		return true
	}
	if v == true {
		return false
	}
	return true
}

// Process 博弈完成之后要做的处理
func (p * Player10) Process (res bool, id int) bool {
	p.list[id] = res
	return true
}