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

type CellType int

// Cell 一个格子的信息
type Cell struct {
	X       int      `json:"x"`
	Y       int      `json:"y"`
	VectorX int      `json:"vx"`
	VectorY int      `json:"vy"`
	No      int      `json:"n"`
	Label   string   `json:"l"`
	Type    CellType `json:"t"`
}
