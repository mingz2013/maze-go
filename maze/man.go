package maze

type ManPoint struct {
	X int
	Y int
	Z int
}

func NewMainPoint() *ManPoint {
	return &ManPoint{}
}

func (m *ManPoint) Init() {

}

func (m *ManPoint) IsWin() {
	// (10,10,10)

}

func (m *ManPoint) Run() {
	// 随机移动一格，测试是否通过

}
