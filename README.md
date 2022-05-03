# 这个后端非常的简单

### 它只有一个 WebSocket 

#### HelloType、RenameType、EnterType

当有新用户连接时，服务器会主动向它发送一个 hello 信息

```json
{
  "type": "hello",
  "data": {
    "name": "XXXXXXXXX"
  }
}
```

客户端收到后，会尝试改名，发送 rename 信息

```json
{
  "type": "rename",
  "data": {
    "avatar": "https://xxxx.com/avatar.png",
    "visitor": false,
    "name": "YYYYYYYYY"
  }
}
```

服务器收到 Rename 信息后，会把信息 enter 广播给其他用户。这个 enter 信息包括了用户的头像、位置、是否为访客。

```json
{
  "type": "enter",
  "data": {
    "name": "YYYYYYYYY",
    "avatar": "https://xxxx.com/avatar.png",
    "position": {
      "x": 0,
      "y": 0
    },
    "visitor": false
  }
}
```

#### StandUpType

如果用户需要获取在线列表，需要他站起来，它会发送一个 stand up 信息给服务器

```json
{
  "type": "stand up"
}
```

服务器收到后会发送给用户当前在线列表

```json
{
  "type": "stand up",
  "data": {
    "onlineUser": [
      {
        "name": "YYYYYYYYY",
        "avatar": "https://xxxx.com/avatar.png",
        "position": {
          "x": 0,
          "y": 0
        },
        "visitor": false
      },
      {
        "name": "ZZZZZZZZZ",
        "avatar": "https://xxxx.com/avatar.png",
        "position": {
          "x": 0,
          "y": 0
        },
        "visitor": false
      }
    ]
  }
}
```

#### MoveType

当用户在界面上拖动小球，会主动向服务器发送 move 消息。服务器会将这个小球原封不动的广播给其他用户。

```json
{
  "type": "move",
  "data": {
    "name": "YYYYYYYYY",
    "position": {
      "x": 0,
      "y": 0
    }
  }
}
```

#### TalkType

当用户在界面上双击小球，会主动向服务器发送 talk 消息。服务器会将这个小球原封不动的广播给其他用户。

```json
{
  "type": "talk",
  "data": {
    "name": "YYYYYYYYY",
    "content": "Hello, World!"
  }
}
```

#### LeaveType

当用户离开页面，或者因不明原因断开时，服务器会广播改用户 Leave 的消息

```json
{
  "type": "leave",
  "data": {
    "name": "YYYYYYYYY"
  }
}
```

