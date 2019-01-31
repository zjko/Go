package player

import(
	"math/rand"
	"time"
)
// Player11 策略2
// 针锋相对策略 合作与背叛，取决于上一次对方的选择
type Player11 struct{
	BasicPlayer
	list map[int]*[2]int
}
// Init 类初始化
func (p * Player11) Init () bool {
	p.name = "欺软怕硬"
	p.list = make(map[int]*[2]int)
	rand.Seed(time.Now().Unix())
	return true
}

// Game 实现博弈接口
func (p * Player11) Game (id int) bool {
	v,ok := p.list[id]
	if !ok {
		p.list[id] = new([2]int)
		// 0 存储的是我方背叛次数 1 对方背叛次数
		p.list[id][0] = 1	
		p.list[id][1] = 0
		return false
	}
	if v[1]==0 || rand.Int()%100 < (int)((float64)(v[0] - v[1])/(float64)(v[0] + v[1]) * 100) {
		p.list[id][0]++
		return false
	}
	return true
}

// Process 博弈完成之后要做的处理
func (p * Player11) Process (res bool, id int) bool {
	if !res {
		p.list[id][1]++
	}
	return true
}