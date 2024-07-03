package v1

type Layout struct {
	Id    int    `json:"id" sm:"id"`
	Name  string `json:"name" sm:"名称"`
	Seats int    `json:"seats" sm:"座位数"`
}
