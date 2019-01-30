package main
import (
	"fmt"
)

func main(){
	s1 := [...]int{}
	for i:=0; i<10; i++ {
		s1[i] = i
	}
	fmt.Printf("%v",s1)
}