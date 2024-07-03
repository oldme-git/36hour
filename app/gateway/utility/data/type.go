package data

// Id 主键类型
type Id uint32

// IdInput 公共Id api input
type IdInput struct {
	Id Id `json:"id" v:"required|integer|between:1,4294967295"`
}
