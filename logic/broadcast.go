/*
# @Time : 2021/3/14 21:17
# @Author : smallForest
# @SoftWare : GoLand
*/
package logic

type broadcaster struct {
	users map[string]*User

	leavingChannel chan *User    // 断开连接的用户
	messageChannel chan *Message // 给用户发消息的channel，广播有用，私聊不必
}

// 单例-饿汉式模式
var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	leavingChannel: make(chan *User, 1000),
	messageChannel: make(chan *Message, 1000),
}

// 将用户加入广播器
func (b broadcaster) UserEntering(user *User) int {
	b.users[user.GetToken()] = user
	return len(b.users)
}

// 将用户断开连接 从广播器剔除
func (b broadcaster) UserLeaving(user *User) {
	b.leavingChannel <- user
}

// 向广播器messageChannel 写入消息
// 格式和前端定义 user_token ,to_user_token message
func (b broadcaster) BroadcastMessage(receiveMsg map[string]string) {
	user_token := receiveMsg["user_token"]
	to_user_token := receiveMsg["to_user_token"]
	message := receiveMsg["message"]
	if receiveMsg["heart"] == "ping" || user_token == "" || to_user_token == "" || message == "" || b.users[user_token] == nil || b.users[to_user_token] == nil {
		return
	}

	// 通过token寻找user 创建消息 msg
	b.messageChannel <- CreateMessage(b.users[user_token], b.users[to_user_token], message, "message")
}

// channel数据处理
func (b broadcaster) Start() {
	for {
		select {
		case user := <-b.leavingChannel:
			// 用户断开了或者服务器主动关闭
			delete(b.users, user.token)
			// 避免user中的goroutinue泄露
			user.CloseMessageChannel()
		case msg := <-b.messageChannel:
			// 给to_user发消息
			msg.ToUser.PutMessage(msg)
		}
	}
}
