package player

// BasicPlayer 策略基类
type BasicPlayer struct{
	score int
	name string
}

// Init 类初始化
func (p * BasicPlayer)Init() bool {
	p.name = "未定义名称"
	return true
}

// Game 实现博弈接口
func (p * BasicPlayer)Game(id int) bool {
	return true
}

// Process 博弈完成之后要做的处理
func (p * BasicPlayer)Process(res bool,id int) bool {
	return true
}

// GetName 返回名称
func (p * BasicPlayer)GetName() string {
	return p.name
}

// GetScore 返回分数
func (p * BasicPlayer)GetScore() int {
	return p.score
}

// SetScore 设置分数
func (p * BasicPlayer)SetScore(s int) {
	p.score = s
}

// AddScore score的加法运算
func (p * BasicPlayer)AddScore(s int) {
	p.score += s
}