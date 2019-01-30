package strategy

// IStrategy 策略接口，作为一个策略必须要实现的
type IStrategy interface{
	// // 是否能够进行博弈
	// CanGame()

	// 所有策略必须实现博弈算法
	Game(id int) bool

	// 博弈之后做的处理
	Process(res bool,id int) bool

	// 所有策略都可以进行初始化
	Init() bool

	// 所有策略都必须可以输出名称
	GetName() string 

	// 所有策略都必须可以获取分数
	GetScore() int

	// 修改分数
	SetScore(s int)
	
	// score的加法运算
	AddScore(s int)

}