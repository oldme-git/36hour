// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IReserve interface{}
)

var (
	localReserve IReserve
)

func Reserve() IReserve {
	if localReserve == nil {
		panic("implement not found for interface IReserve, forgot register?")
	}
	return localReserve
}

func RegisterReserve(i IReserve) {
	localReserve = i
}
