package player

import(
	"math/rand"
	"time"
)
// Player6 禁闭策略
// 类似封杀策略，但是在连续背叛对方3次之后，则会将对方从黑名单中消除。
type Player6 struct{
	BasicPlayer
	list map[int]int
}

// Init 类初始化
func (p * Player6) Init () bool {
	p.name = "信誉积分"
	rand.Seed(time.Now().Unix())
	p.list = make(map[int]int)
	return true
}

// Game 实现博弈接口
func (p * Player6)Game(id int) bool {
	v , ok := p.list[id]
    if !ok {
		p.list[id] = 100
		return true
	}
	if v > 100 {
		p.list[id] = 100
	}

	if rand.Int()%100 < v {
		return true
	}
	return false
}

// Process 博弈完成之后要做的处理
func (p * Player6)Process(res bool, id int) bool {
	if !res {
		p.list[id] /= 2
	} else {
		p.list[id] += 30
	}
	return true
}