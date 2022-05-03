package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
	Avatar   string   `json:"avatar"`
	Visitor  bool     `json:"visitor"`

	connect *websocket.Conn
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
			continue
		}

		switch request.Type {

		case MoveType:
			u.Position = request.Data.Position
			chatroom.broadcast(u, MoveMessage(u, request.Data), false)

		case TalkType:
			chatroom.broadcast(u, TalkMessage(u, request.Data), true)

		case RenameType:
			// prevent same name
			chatroom.rw.RLock()

			var response = true
			for _, user := range chatroom.UserMap {
				if user.Name == request.Data.Name {
					response = false
					break
				}
			}
			chatroom.rw.RUnlock()

			if response == false {
				continue
			}

			u.Name = request.Data.Name
			chatroom.broadcast(u, EnterMessage(u), false)

		case StandUpType:
			chatroom.rw.RLock()
			var onlineUserList []*User
			for _, user := range chatroom.UserMap {
				onlineUserList = append(onlineUserList, user)
			}
			chatroom.rw.RUnlock()

			u.Send(StandUpMessage(u, onlineUserList))

		default:
		}
	}
}
