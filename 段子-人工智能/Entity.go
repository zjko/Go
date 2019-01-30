package main
import(
	// "./AICore"
	"fmt"
)

func main(){
	// AICore.Master()
	chs := make([]chan int, 10)
	for i:= 0; i < 10; i++ {
		// fmt.Println("this is", i)
		chs[i] = make(chan int)
		go Add(chs[i], i)

	}

	for _, ch := range(chs) {
		fmt.Println(<- ch)
	}
	
}


func Add(ch chan int, i int) {
	// fmt.Println("continue", ch)
	ch <- i
}
