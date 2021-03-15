/*
# @Time : 2021/3/14 21:20
# @Author : smallForest
# @SoftWare : GoLand
*/
package logic

type Message struct {
	User    *User  `json:"user"`    // 发动消息的人
	ToUser  *User  `json:"to_user"` // 接受消息的人
	Content string `json:"content"` // 消息内容
}

func CreateMessage(user *User, to_user *User, msg string) *Message {
	return &Message{User: user, ToUser: to_user, Content: msg}
}

