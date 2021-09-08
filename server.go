/*
# @Time : 2021/3/14 21:27
# @Author : smallForest
# @SoftWare : GoLand
*/
package main

import (
	"dataSocket/logic"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "HTTP, Hello")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		// conn是客户端连接
		conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
			//OriginPatterns: []string{"localhost:63342", "192.168.0.125:8080"},
		})
		if err != nil {
			log.Println(err)
			return
		}
		// 使用conn生成用户
		user := logic.NewUser(conn)
		log.Println(user)

		// 开启给用户发送消息的goroutinue
		go user.SendMessage(req.Context())

		// 用户连接上就发送token信息
		user.PutMessage(logic.CreateMessage(user, user, user.GetToken(), "token"))

		// 将用户添加到广播器的用户列表
		c := logic.Broadcaster.UserEntering(user)
		log.Println("将用户添加到广播器的用户列表后，当前用户连接", c, "个")

		// 持续接受用户消息 用户不离开阻塞
		err = user.ReceiveMessage(req.Context())

		// 用户离开操作
		logic.Broadcaster.UserLeaving(user)
		if err == nil {
			// 服务器关闭连接
			conn.Close(websocket.StatusNormalClosure, "server close conn")
			log.Println("客户端或服务端主动关闭")
		} else {
			log.Println("客户端错误：", err)
			conn.Close(websocket.StatusInternalError, "客户端错误")
		}

	})
	// 主业务处理协程
	go logic.Broadcaster.Start()

	log.Println("ws://localhost:2021/ws")
	log.Fatal(http.ListenAndServe(":2021", nil))
}
