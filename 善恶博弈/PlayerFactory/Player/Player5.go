package player

// Player5 禁闭策略
// 类似封杀策略，但是在连续背叛对方3次之后，则会将对方从黑名单中消除。
type Player5 struct{
	BasicPlayer
	blacklist map[int]int
}

// Init 类初始化
func (p * Player5) Init () bool {
	p.name = "双倍偿还"
	p.blacklist = make(map[int]int)
	return true
}

// Game 实现博弈接口
func (p * Player5)Game(id int) bool {
	_ , ok := p.blacklist[id]
    if !ok {
		// 不在黑名单中
        return true
	}
	// fmt.Println(p.history,id,p.history[id])
	return false
}

// Process 博弈完成之后要做的处理
func (p * Player5)Process(res bool, id int) bool {
	v , ok := p.blacklist[id]
	if ok {
		// 在黑名单中
		if res {
			v++
			if v > 2 {
				delete(p.blacklist, id)
				return true
			}
			p.blacklist[id]++
			// 被我背叛
		} else {
			p.blacklist[id] = 0
		}
	} else {
		// 不在黑名单中
		if !res {
			// 加入黑名单
			p.blacklist[id] = 0
		}
	}
	

	return true
}