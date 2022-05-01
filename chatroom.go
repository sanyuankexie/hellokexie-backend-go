package main

import (
	"sync"

	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Chatroom struct {
	JoinList  chan *websocket.Conn
	LeaveList chan *websocket.Conn

	UserMap map[*websocket.Conn]*User
	rw      sync.RWMutex
}

var chatroom *Chatroom

func init() {
	chatroom = &Chatroom{
		JoinList:  make(chan *websocket.Conn, 100),
		LeaveList: make(chan *websocket.Conn, 100),
		UserMap:   make(map[*websocket.Conn]*User),
	}

	go chatroom.welcome()
	go chatroom.goodbye()
}

func (c *Chatroom) addUser(newUser *User) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.UserMap[newUser.connect] = newUser
}

func (c *Chatroom) removeUser(newUser *User) {
	c.rw.Lock()
	defer c.rw.Unlock()

	delete(c.UserMap, newUser.connect)
}

func (c *Chatroom) welcome() {
	for ws := range c.JoinList {
		newUser := User{
			Name:    uuid.New().String(),
			connect: ws,
		}
		err := newUser.Send(HelloMessage(&newUser))
		if err != nil {
			log.Printf("send hello message err: %s\n", err.Error())
			c.LeaveList <- ws
			return
		}

		go newUser.listen(c)
		c.addUser(&newUser)
	}
}

func (c *Chatroom) goodbye() {
	for ws := range c.LeaveList {
		ws.Close()
		c.removeUser(c.UserMap[ws])
	}
}

func (c *Chatroom) broadcast(sender *User, message string, self bool) error {
	for _, user := range c.UserMap {
		if self || sender.Name != user.Name {
			err := user.Send(message)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
