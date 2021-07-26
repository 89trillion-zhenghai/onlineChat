# 项目介绍

本项目实现了在线聊天功能，使用websocket作为连接协议，protobuf类型的二进制数据作为通信数据格式。当用户连接至本服务器时，服务器用一个map作为保存连接信息的容器。并且为此连接开辟两个goroutine负责接收该连接的消息和发送消息到这个连接。客户端的连接、断开、发消息以及获取所在在线用户列表均由其实现。

## 目录结构

```tree
├── README.md
├── __pycache__
│   └── locust.cpython-39.pyc
├── app
│   ├── http
│   │   └── httpServer.go
│   ├── main
│   ├── main.go
│   └── ws
│       ├── wsClient.go
│       └── wsServer.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── chatCtrl.go
│   ├── handler
│   │   └── chatHandler.go
│   ├── model
│   │   ├── Client.go
│   │   └── Server.go
│   └── route
│       └── chatRoute.go
├── locust.py
├── log
│   └── server.log
├── logUtil
│   └── logger.go
├── message
│   ├── message.pb.go
│   ├── message.proto
│   └── msgUtil.go
└── report.html
```

## 代码逻辑分层

| 层      | 文件夹            | 功能介绍                                                     | 调用关系      |
| ------- | ----------------- | ------------------------------------------------------------ | ------------- |
| ws      | /app/ws           | 管理websocket连接，断开，发送消息，服务器的初始化            | 被ctrl调用    |
| Route   | /internal/route   | 路由转发                                                     | 调用ctrl层    |
| ctrl    | /Internal/ctrl    | 参数校验、升级连接成ws连接                                   | 调用handler层 |
| handler | /Internal/handler | 对新的连接初始化，保存到manager管理容器，启动发送、接收消息的goroutine | 被ctrl层调用  |
| model   | /internal/model   | 数据模型、                                                   | 被其他层调用  |
| log     | /log              | 存放日志文件server.log                                       | 被logUtil写入 |
| logUtil | /logUtil          | 初始化日志文件对象                                           | 被其他层调用  |
| message | /message          | 存放protubuf文件以及封装了几个常用类型的消息                 | 被其他层调用  |

## 存储设计

#### 通信数据格式

| 变量名     | 变量类型 | 变量含义 |
| ---------- | -------- | -------- |
| msgType    | string   | 消息类型 |
| msgContent | string   | 消息内容 |
| sendName   | string   | 用户名   |
| userList   | []string | 用户列表 |

## 接口设计

| 接口地址               | headers             | 响应参数   | 请求方法  |
| ---------------------- | ------------------- | ---------- | --------- |
| ws://localhost:8080/ws | {"name":"smallbai"} | 二进制数据 | websocket |

## 第三方库

### gin

```
go语言的web框架
https://github.com/gin-gonic/gin
```

### Gorilla WebSocket

```
go语言对websocket协议的实现
http://github.com/gorilla/websocket
```

### protobuf

```
包含go语言处理proto数据的函数
http://github.com/golang/protobuf/proto
```

## 如何编译执行

进入app目录编译

```
go build
```

运行可执行文件

```
./main
```

## 流程图

![流程图](/Users/yangzhenghai/workspace/onlineChat/流程图.jpg)