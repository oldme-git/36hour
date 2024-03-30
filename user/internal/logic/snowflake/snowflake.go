package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"user/internal/service"
)

type sSnowflake struct {
	node *sf.Node
}

func init() {
	service.RegisterSnowflake(New())
}

func New() *sSnowflake {
	node, err := sf.NewNode(1)
	if err != nil {
		panic(err)
	}
	return &sSnowflake{node: node}
}

func (s *sSnowflake) Get() sf.ID {
	id := s.node.Generate()
	return id
}
