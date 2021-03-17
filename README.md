# 基于websocket的两端通讯
在实际开发业务中某些场景要用，当然客服聊天系统，聊天室系统都可以基于此开发。
目前的功能可满足我目前的实际需求。
# logic声明struct
- 广播站结构体
- 用户结构体
- message结构体
# 使用方式
1. 连接websocket之后会立马接收到token，该socket链接唯一标识
```javascript
{"content":"1896d382-8cd5-466b-a3e5-5d955600c6d6","type":"token"}
```
type=message 或者 token  
type=token时候content是token标识  
type=message是发送的消息  
2. 发送信息格式
```javascript
setInterval(function () {
    //user_token ,to_user_token message
    let obj = {
        user_token:'ca7c99cb-2120-4f3f-9a2b-a2d045848e89',
        to_user_token:'efb744d3-5ff5-462a-808b-ae3d1bdf4f96',
        message:"xxx",
    }
    websocket.send(JSON.stringify(obj))
}, 3000)
```
