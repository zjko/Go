package player

// Player4 策略1
// 无论是谁都将无条件的合作
type Player4 struct{
	BasicPlayer
	blacklist map[int]int
}

// Init 类初始化
func (p * Player4) Init () bool {
	p.name = "禁闭策略"
	p.blacklist = make(map[int]int)
	return true
}

// Game 实现博弈接口
func (p * Player4) Game (id int) bool {
	_ , ok := p.blacklist[id]
    if !ok {
        return true
	}
	// fmt.Println(p.history,id,p.history[id])
	return false
}

// Process 博弈完成之后要做的处理
func (p * Player4) Process (res bool, id int) bool {
	// if !res {
	// 	p.blacklist[id] = (int)res
	// }
	return true
}