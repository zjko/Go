package player

import(
	"math/rand"
	"time"
)
// Player9 策略2
// 针锋相对策略 合作与背叛，取决于上一次对方的选择
type Player9 struct{
	BasicPlayer
	
}
// Init 类初始化
func (p * Player9) Init () bool {
	p.name = "摇摆不定"
	rand.Seed(time.Now().Unix())
	return true
}

// Game 实现博弈接口
func (p * Player9) Game (id int) bool {
	r := false
	if rand.Int() %2 == 0 {
		r = true
	}
	return r
}

// Process 博弈完成之后要做的处理
func (p * Player9) Process (res bool, id int) bool {
	return true
}