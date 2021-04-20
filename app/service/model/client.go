package model

import (
	"github.com/gorilla/websocket"
	"kChatRoom/app/client/model/userModel"
)

type Client struct {
	Conn *websocket.Conn
	User *userModel.UserModel
}
