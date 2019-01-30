package player

// Player3 策略1
// 无论是谁都将无条件的合作
type Player3 struct{
	BasicPlayer
	blacklist map[int]bool
}

// Init 类初始化
func (p * Player3) Init () bool {
	p.name = "封杀策略"
	p.blacklist = make(map[int]bool)
	return true
}

// Game 实现博弈接口
func (p * Player3) Game (id int) bool {
	_ , ok := p.blacklist[id]
    if !ok {
        return true
	}
	// fmt.Println(p.history,id,p.history[id])
	return false
}

// Process 博弈完成之后要做的处理
func (p * Player3) Process (res bool, id int) bool {
	if !res {
		p.blacklist[id] = res
	}
	return true
}