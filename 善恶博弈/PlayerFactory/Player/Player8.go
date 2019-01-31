package player


// Player8 禁闭策略
// 类似封杀策略，但是在连续背叛对方3次之后，则会将对方从黑名单中消除。
type Player8 struct{
	BasicPlayer
	list map[int]int
}

// Init 类初始化
func (p * Player8) Init () bool {
	p.name = "背后捅刀"
	p.list = make(map[int]int)
	return true
}

// Game 实现博弈接口
func (p * Player8)Game(id int) bool {
	v , ok := p.list[id]
    if !ok {
		p.list[id] = 1
		return true
	}
	if v > 1 {
		p.list[id] = 0
		return false
	}
	p.list[id]++
	return true
}

// Process 博弈完成之后要做的处理
func (p * Player8)Process(res bool, id int) bool {
	
	return true
}