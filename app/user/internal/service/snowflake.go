// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	sf "github.com/bwmarrin/snowflake"
)

type (
	ISnowflake interface {
		Get() sf.ID
	}
)

var (
	localSnowflake ISnowflake
)

func Snowflake() ISnowflake {
	if localSnowflake == nil {
		panic("implement not found for interface ISnowflake, forgot register?")
	}
	return localSnowflake
}

func RegisterSnowflake(i ISnowflake) {
	localSnowflake = i
}