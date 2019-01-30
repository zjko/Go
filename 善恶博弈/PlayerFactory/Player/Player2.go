package player

// import(
// 	"fmt"
// )
// Player2 策略2
// 针锋相对策略 合作与背叛，取决于上一次对方的选择
type Player2 struct{
	BasicPlayer
	history map[int]bool
}
// Init 类初始化
func (p * Player2) Init () bool {
	p.name = "针锋相对"
	p.history = make(map[int]bool)
	return true
}

// Game 实现博弈接口
func (p * Player2) Game (id int) bool {
	v,ok := p.history[id]
    if !ok {
        return true
	}
	// fmt.Println(p.history,id,p.history[id])
	return v
}

// Process 博弈完成之后要做的处理
func (p * Player2) Process (res bool, id int) bool {
	p.history[id] = res
	return true
}