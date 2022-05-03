package main

import "encoding/json"

type Request struct {
	Type string `json:"type"`
	Data *Data  `json:"data"`
}

type Response struct {
	Type string `json:"type"`
	Data *Data  `json:"data"`
}

func ParseRequest(message string) (*Request, error) {
	request := Request{}
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

type Data struct {
	Name       string   `json:"name"`
	Avatar     string   `json:"avatar"`
	Visitor    bool     `json:"visitor"`
	Content    string   `json:"content"`
	Position   Position `json:"position"`
	OnlineUser []*User  `json:"onlineUser"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CreateResponse(tp string, data *Data) string {
	response := Response{
		Type: tp,
		Data: data,
	}
	bytes, _ := json.Marshal(&response)
	return string(bytes)
}

func HelloMessage(newUser *User) string {
	return CreateResponse(
		HelloType,
		&Data{
			Name: newUser.Name,
		},
	)
}

func MoveMessage(user *User, data *Data) string {
	return CreateResponse(
		MoveType,
		data,
	)
}

func TalkMessage(user *User, data *Data) string {
	return CreateResponse(
		TalkType,
		data,
	)
}

func EnterMessage(user *User) string {
	return CreateResponse(
		EnterType,
		&Data{
			Name:   user.Name,
			Avatar: user.Avatar,
			Position: Position{
				X: user.Position.X,
				Y: user.Position.Y,
			},
			Visitor: user.Visitor,
		},
	)
}

func StandUpMessage(user *User, onlineUserList []*User) string {
	return CreateResponse(
		StandUpType,
		&Data{
			OnlineUser: onlineUserList,
		},
	)
}

func LeaveMessage(user *User) string {
	return CreateResponse(
		LeaveType,
		&Data{
			Name: user.Name,
		},
	)
}

const (
	HelloType   = "hello"
	EnterType   = "enter"
	MoveType    = "move"
	RenameType  = "rename"
	StandUpType = "stand up"
	LeaveType   = "leave"
	TalkType    = "talk"
)
