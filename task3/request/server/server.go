package server

import (
	"task3/pb"
)

type User struct {
	userClient pb.UserClient
}

func NewDevice(client pb.UserClient) *User {
	return &User{userClient: client}
}
