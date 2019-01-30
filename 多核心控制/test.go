package main
import (
	"fmt"
	"runtime"
	"math/rand"
	"time"
)

// VECTORSIZE = 10000000
const VECTORSIZE = 1000000
// NCPU is xxx
var NCPU int 

func main () {
	// 打印CPU核心数
	NCPU= runtime.GOMAXPROCS(NCPU)
	// 设置GO处理内核数目为当前内核数
	fmt.Println("CoreNum:",NCPU,"    total data:",VECTORSIZE)	

	var v Vector
	v.init()
	// 初始化数据
	c:= make(chan int,NCPU)
	d:= int(VECTORSIZE/NCPU)

	start_t := time.Now()
	for i:=0; i<NCPU; i++{
		go v.sum(i*d,(i+1)*d,c)
	}
	sum:=0
	for i:=0;i<NCPU;i++{
		sum+=<-c
	}
	time := time.Since(start_t)
	fmt.Println("Test Time:",time,"SUM=",sum)
	// 协程计算 花费的时间

	v.sumAll()
	// 迭代计算 花费的时间
	
}

// Vector is
type Vector [VECTORSIZE]int

func (v *Vector) init (){
	start_t := time.Now()
	SUM := 0
	rand.Seed(time.Now().Unix())
	for i:=0; i<VECTORSIZE; i++{
		v[i] = rand.Int() % 10000
		SUM += v[i]
	}
	time := time.Since(start_t)
	fmt.Println("init time:",time,"SUM = ", SUM, "Ready！")
}

func (v *Vector)sum(start,end int, c chan int) {
	start_t := time.Now()
	sum := 0
	for i:=start; i<end; i++ {
		sum += v[i]
	}
	time := time.Since(start_t)
	c <- sum
	fmt.Println("----GO ID",start, "SUM=",sum,"time:",time)
}

func (v *Vector)sumAll(){
	start_t := time.Now()
	sum:=0
	for i:=0;i<VECTORSIZE;i++ {
		sum+= v[i]
	}
	time := time.Since(start_t)
	fmt.Println("Sum Time:",time,"SUM=",sum)
}