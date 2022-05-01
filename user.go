package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	Name     string   `json:"userName"`
	Position Position `json:"position"`
	Avatar   string   `json:"avatar"`
	Visitor  bool     `json:"visitor"`

	connect *websocket.Conn `json:"-"`
}

func (u *User) Send(message string) error {
	return u.connect.WriteMessage(1, []byte(message))
}

func (u *User) listen(chatroom *Chatroom) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("user %s leaves\n, due to %s", u.Name, r)
			chatroom.broadcast(u, LeaveMessage(u), false)
			chatroom.LeaveList <- u.connect
		}
	}()

	for {
		_, message, err := u.connect.ReadMessage()

		if err != nil {
			log.Printf("user %s leaves, read message err: %s\n", u.Name, err.Error())
			chatroom.broadcast(u, LeaveMessage(u), false)
			chatroom.LeaveList <- u.connect
			return
		}

		request, err := ParseRequest(string(message))
		if err != nil {
			log.Printf("parse request err: %s\n", err.Error())
			fmt.Printf("[%s]\n", string(message))
			chatroom.broadcast(u, string(message), true)
			continue
		}

		switch request.Type {

		case MoveType:
			request.Data.Position.X = request.Data.X
			request.Data.Position.Y = request.Data.Y
			u.Position = request.Data.Position
			chatroom.broadcast(u, MoveMessage(u, request.Data.Position), false)

		case TalkType:
			// todo
			chatroom.broadcast(u, string(message), false)

		case RenameType:
			u.Name = request.UserName
			chatroom.broadcast(u, EnterMessage(u), false)

		case StandUpType:
			onlineUserList := []*User{}
			for _, user := range chatroom.UserMap {
				onlineUserList = append(onlineUserList, user)
			}

			u.Send(StandUpMessage(u, onlineUserList))

		default:
		}
	}
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
