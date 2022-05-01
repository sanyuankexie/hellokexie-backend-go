package main

import "encoding/json"

type Request struct {
	UserName string      `json:"userName"`
	Type     string      `json:"type"`
	Data     RequestData `json:"data"`
}

type Response struct {
	UserName string      `json:"userName"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}

func ParseRequest(message string) (*Request, error) {
	request := Request{}
	err := json.Unmarshal([]byte(message), &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

type RequestData struct {
	X        int      `json:"x"`
	Y        int      `json:"y"`
	Avatar   string   `json:"avatar"`
	Visitor  bool     `json:"visitor"`
	Position Position `json:"position"`
}

func CreateResponse(tp string, data any, userName string) string {
	response := Response{
		UserName: userName,
		Type:     tp,
		Data:     data,
	}
	bytes, _ := json.Marshal(&response)
	return string(bytes)
}

func HelloMessage(newUser *User) string {
	return CreateResponse(HelloType, "hello "+newUser.Name, newUser.Name)
}

func MoveMessage(user *User, data any) string {
	return CreateResponse(MoveType, data, user.Name)
}

func TalkMessage(user *User, data any) string {
	return CreateResponse(MoveType, data, user.Name)
}

func EnterMessage(user *User) string {
	response := struct {
		Avatar   string   `json:"avatar"`
		Position Position `json:"position"`
		Visitor  bool     `json:"visitor"`
	}{
		Avatar:   user.Avatar,
		Position: user.Position,
		Visitor:  user.Visitor,
	}

	return CreateResponse(EnterType, response, user.Name)
}

func StandUpMessage(user *User, onlineUserList []*User) string {
	return CreateResponse(StandUpType, onlineUserList, user.Name)
}

func LeaveMessage(user *User) string {
	return CreateResponse(LeaveType, nil, user.Name)
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
