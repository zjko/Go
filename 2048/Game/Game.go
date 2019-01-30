package Game
import(
	"fmt"
	"math/rand"
	"time"
	"os/exec"
	"os"
)
// Game 游戏整体控制类
type Game struct {
	data [4][4]int
	score int
	status int
}

// Point 点位信息
type Point struct{
	x,y int
}

// IGame 游戏类接口
type IGame interface {
	// Input 输入控制器
	Input()
	// Show 显示控制
	Show()


}


func (g * Game) left () {
	for k := 4; k-1 > 0; k-- {
        for y := 1; y < 4; y++ {
            for x := 0; x < 4; x++ {
                if (g.data[y - 1][x] == g.data[y][x] || g.data[y - 1][x] == 0) {
                    g.data[y - 1][x] += g.data[y][x]
                    g.data[y][x] = 0
                }
            }
        }
	}
	g.new()
}

func (g * Game) right () {
	for k := 4; k-1 > 0;k-- {
        for y := 2; y >= 0; y-- {
            for x := 0; x < 4; x++ {
                if (g.data[y + 1][x] == g.data[y][x] || g.data[y + 1][x] == 0) {
                    g.data[y + 1][x] += g.data[y][x];
                    g.data[y][x] = 0
                }
            }
        }
	}
	g.new()
}

func (g * Game) up () {
	for k := 4; k - 1 > 0; k-- {
        for x := 1; x < 4; x++ {
            for y := 0; y < 4; y++ {
                if (g.data[y][x - 1] == g.data[y][x] || g.data[y][x - 1] == 0) {
                    g.data[y][x - 1] += g.data[y][x]
                    g.data[y][x] = 0
                }
            }
        }
	}
	g.new()
}

func (g * Game) down () {
	for k := 4; k - 1 > 0; k-- {
        for x := 2; x >= 0; x-- {
            for y := 0; y < 4; y++ {
                if (g.data[y][x + 1] == g.data[y][x] || g.data[y][x + 1] == 0) {
                    g.data[y][x + 1] += g.data[y][x];
                    g.data[y][x] = 0;
                }
            }
        }
	}
	g.new()
}

func (g * Game) show () {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("score:",g.score)
	for i:=0;i<4;i++ {
		for j:=0;j<4;j++ {
			switch g.data[j][i] {
				case 0:
					fmt.Printf("　　");
					break;
				case 2:
					fmt.Printf("\033[48;32;5m 2　\033[47;32;0m");
					break;
				case 4:
					fmt.Printf("\033[47;31;5m 4　\033[43;37;0m");
					break;
				case 8:
					fmt.Printf("\033[41;37;5m 8　\033[43;37;0m");
					break;
				case 16:
					fmt.Printf("\033[42;37;5m 16 \033[43;37;0m");
					break;
				case 32:
					fmt.Printf("\033[43;37;5m 32 \033[43;37;0m");
					break;
				case 64:
					fmt.Printf("\033[44;37;5m 64 \033[43;37;0m");
					break;
				case 128:
					fmt.Printf("\033[45;37;5m128 \033[43;37;0m");
					break;
				case 256:
					fmt.Printf("\033[46;37;5m256 \033[43;37;0m");
					break;
				case 512:
					fmt.Printf("\033[47;37;5m512 \033[43;37;0m");
					break;
				case 1024:
					fmt.Printf("\033[48;37;5m1024 \033[43;37;0m");
					break;
				case 2048:
					fmt.Printf("\033[49;37;5m2048 \033[43;37;0m");
					break;
				default: 
			}
			
		}
		fmt.Printf("\n")
	}
}


func (g * Game) new () bool{
	var list [16]Point
	// list 用于保存data当中所有的空块信息

	k := 0
	for i:=0;i<4;i++ {
		for j:=0;j<4;j++ {
			if g.data[i][j] == 0 {
				list[k] =  Point{i,j}
				k++
			}
		}
	}
	// fmt.fmt.Printf("%v\n",list)
	if k < 1 {
		// 若找不到一个空闲块则返回false
		g.status = -1
		return false
	}

	// 随机选取list当中的空闲块并随机赋值为2或4
	r := rand.Int() % k
	score:=(rand.Int() % 2 + 1) * 2
	g.data[list[r].x][list[r].y] = score
	g.score += score
	return true
}

func (g * Game) init () {
	rand.Seed(time.Now().Unix())
	g.new()
	g.new()
	g.show()
}

// Input 键盘输入
func (g * Game) Input() bool {
	var key string
	fmt.Scanln(&key)
	switch key {
	case "w": g.up()
	case "s": g.down()
	case "a": g.left()
	case "d": g.right()
	default: 
		g.status = 1
		return false
	}
	return true
}



// Start 游戏开始
func (g * Game) Start () {
	g.init()
	for true {
		g.Input()
		switch g.status {
		case -1: 
			fmt.Println("GAME OVER")
			return
		case 0: 
			g.show()
		case 1:
			fmt.Println("Invalid input")
		default:
			fmt.Println("Unknown fatal error, Status: ",g.status)
			return
		}
		g.status = 0
	}
	


}