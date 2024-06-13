package layout

// 1-座位，2-桌子，3-书架，4-墙，5-门，6-窗户
const (
	CellSeat = iota + 1
	CellTable
	CellBookshelf
	CellWall
	CellDoor
	CellWindow
)

// 1-空闲，2-使用中，3-锁定
const (
	CellStatusFree = iota + 1
	CellStatusUsing
	CellStatusLock
)

// CellType 格子类型
type CellType int

// CellStatus 格子状态
type CellStatus int

// Cell 一个格子的信息
type Cell struct {
	X       int        `json:"x"`
	Y       int        `json:"y"`
	VectorX int        `json:"vx"`
	VectorY int        `json:"vy"`
	No      int        `json:"n"`
	Label   string     `json:"l"`
	Type    CellType   `json:"t"`
	Status  CellStatus `json:"s"`
}
