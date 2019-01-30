package player

/*Player3 封杀策略
*策略首先会信任任何人，
*如果有人背叛了他，
*那么这个人则会进入黑名单，
*然后在之后的每一次的博弈中，
*这个策略都会选择背叛。
*/
type Player3 struct{
	BasicPlayer
	blacklist map[int]bool
}

// Init 类初始化
func (p * Player3)Init() bool {
	p.name = "封杀策略"
	p.blacklist = make(map[int]bool)
	return true
}

// Game 实现博弈接口
func (p * Player3)Game(id int) bool {
	_ , ok := p.blacklist[id]
    if !ok {
        return true
	}
	// fmt.Println(p.history,id,p.history[id])
	return false
}

// Process 博弈完成之后要做的处理
func (p * Player3)Process(res bool, id int) bool {
	if !res {
		p.blacklist[id] = res
	}
	return true
}