package main
import(
	"fmt"
)

type student struct{
	no int
	name string
	class string
	major string
}
func (s student) sayName () {
	fmt.Println(s.name)
}

func main(){
	s := &student{123,"路人甲","计算机","计科专业"}
	fmt.Printf("%v\n",s)
	s.sayName()
	s1 := new(student)
	s1.name = "路人1"
	s1.sayName()
}