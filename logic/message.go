/*
# @Time : 2021/3/14 21:20
# @Author : smallForest
# @SoftWare : GoLand
*/
package logic

type Message struct {
	User    *User  `json:"-"`       // 发动消息的人
	ToUser  *User  `json:"-"`       // 接受消息的人
	Content string `json:"content"` // 消息内容
	Type    string `json:"type"`    //消息类型 token | message
}

func CreateMessage(user *User, to_user *User, msg string, msgType string) *Message {
	return &Message{User: user, ToUser: to_user, Content: msg, Type: msgType}
}
