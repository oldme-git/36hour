package svc_disc

import "google.golang.org/grpc"

// UserClient 获取用户服务的客户端
func UserClient() *grpc.ClientConn {
	return Client("user")
}

// LibManagerClient 获取图书馆管理服务的客户端
func LibManagerClient() *grpc.ClientConn {
	return Client("lib-manager")
}

// SeatClient 获取座位服务的客户端
func SeatClient() *grpc.ClientConn {
	return Client("seat")
}
