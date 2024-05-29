package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func init() {
	var err error
	node, err = sf.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func Get() sf.ID {
	id := node.Generate()
	return id
}
