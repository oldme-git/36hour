package model

type LayoutSearchCondition struct {
	Page       int
	PageSize   int
	LayoutName string
	Status     int
	// 座位数量区间
	SeatsMin int
	SeatsMax int
}

type PolicyCommonSearchCondition struct {
	Name     string
	Page     int
	PageSize int
}

type PolicyPrepareSearchCondition struct {
	Name     string
	Page     int
	PageSize int
}
