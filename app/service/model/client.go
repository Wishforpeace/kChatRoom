package model

import (
	"github.com/gorilla/websocket"
	"kChatRoom/app/client/model/userModel"
)

const (
	ClientTypePeople = "clientPeople"
	ClientTypeRobot  = "clientRobot"
)

type Client struct {
	Conn *websocket.Conn
	Type string
	User *userModel.UserModel
}
