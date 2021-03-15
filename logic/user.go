/*
# @Time : 2021/3/14 21:18
# @Author : smallForest
# @SoftWare : GoLand
*/
package logic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type User struct {
	token          string        // 用户唯一标识
	MessageChannel chan *Message // 用户接受消息的channel，发给本用户的消息直接放到这个channel即可
	conn           *websocket.Conn
}

/**
创建新用户
*/
func NewUser(conn_user *websocket.Conn) *User {
	return &User{token: uuid.New().String(), conn: conn_user, MessageChannel: make(chan *Message,10)}
}

// 获取token
func (u User) GetToken() string {
	return u.token
}

// 向用户的上下文中写入消息
func (u User) SendMessage(ctx context.Context) {
	// 定义死循环
	for msg := range u.MessageChannel {
		_ = wsjson.Write(ctx, u.conn, msg.Content)
	}
}
// 向message channel放信息
func (u User) PutMessage(msg *Message)  {
	u.MessageChannel<-msg
}

//获取用户发过来的信息
func (u User) ReceiveMessage(ctx context.Context) error {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {
		// 从用户连接获取信息
		err = wsjson.Read(ctx, u.conn, &receiveMsg)
		if err != nil {
			// 判断连接是否已经关闭，正常关闭不认为error
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			}
			return err
		}
		// 把消息放入到广播器，由广播器进行处理
		Broadcaster.BroadcastMessage(receiveMsg)
	}
}

func (u User) CloseMessageChannel() {
	close(u.MessageChannel)
}
