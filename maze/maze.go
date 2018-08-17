package maze

// 迷宫,三维迷宫，10*10*10立方体
// TODO 如果迷宫可以变化就更好了

type Maze [10][10][10]int

func NewMaze() *Maze {
	m := &Maze{}
	m.Init()
	return m
}

func (m *Maze) Init() {

}

func (m *Maze) RandomARoute() {
	// 随机出一条路线
	// 从(0, 0, 0)点出发
	// 从(10, 10, 10)点离开
	// 其他边界不可有出口
	//

}
