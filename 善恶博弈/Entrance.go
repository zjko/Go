package main

import(
	"fmt"
	"./PlayerFactory"
	"./PlayerFactory/Strategy"
)

func main(){
	start()
}

func start(){
	// 初始化
	names := [...]string{"Player1","Player2","Player3","Player4","Player5","Player6","Player7","Player8","Player9","Player10","Player11"}
	const ListSize = len(names)
	var list [ListSize] strategy.IStrategy
	factory := new(playerFactory.PlayerFactory)
	for k,v:= range names{
		list[k] = factory.Create(v)
	}

	// 打印所有参赛选手
	fmt.Println("竞技玩家按照序号排序如下：")
	for k,v:= range list{
		fmt.Println(k,v.GetName())
	}

	// 开始竞技
	fmt.Println("PK顺序如下：")
	for t:=0; t<100; t++ {
		for i:=0; i < ListSize - 1; i++ {
			for j:=i+1; j < ListSize ; j++ {
				// fmt.Println(list[i].GetName(),"PK",list[j].GetName())
				PK(list[i],list[j],i,j)
			}
		}
	}
		

	// 最终分数
	fmt.Println("最终每个策略分数：")
	show(list[:])
}

// PK 两个玩家进行对决
func PK(p1, p2 strategy.IStrategy,id1,id2 int){
	
	// 两个策略先分别根据对方编号进行决策，决策结果分别存储在res当中
	res1 := p1.Game(id2)
	res2 := p2.Game(id1)

	// 再分别通知每个决策对方的结果
	p1.Process(res2,id2)
	p2.Process(res1,id1)

	// 计算竞技结果
	if res1 && res2 {
		// 双赢
		p1.AddScore(3)
		p2.AddScore(3)
	} else if !res1 && res2 {
		// p1 背叛 p2被背叛
		p1.AddScore(5)
		p2.AddScore(0)
	} else if res1 && !res2 {
		// p2 背叛 p1被背叛
		p2.AddScore(5)
		p1.AddScore(0)
	} else if !res1 && !res2 {
		// 双输 p1 p2 都被对方背叛
		p2.AddScore(-1)
		p1.AddScore(-1)
	}
	if id1 == 10 || id2 == 10 {
		fmt.Println("PK:\t",p1.GetName(),res1,"\n\t",p2.GetName(),res2)
	}
	
	
}

func show(p []strategy.IStrategy){
	for i:=0;i<len(p);i++ {
		fmt.Println("ID:",i,p[i].GetName(), p[i].GetScore())
	}
}