# 采用微服务架构搭建博客
---
## 目前用到的框架
- grpc
- gin
- gorm

---
## 目录结构
```
.
├── README.md
├── api     接口文件
│   ├── blog.pb.go
│   └── blog.proto
├── client  grpc客户端
│   ├── client.go   客户端主程序main
│   └── rpc         接口实现
│       ├── auth.go
│       └── blog.go
├── config               配置文件
│   ├── config.go
│   └── config.toml
├── go.mod
├── go.sum
└── server      服务端代码
    ├── db      数据库定义
    │   └── db.go
    ├── model   使用到的结构体
    │   ├── blog.go
    │   └── user.go
    ├── rpcimpl 服务端接口实现
    │   ├── auth.go
    │   └── blog.go
    ├── server.go   服务端主程序main
    └── web     gin编写的web服务（路由）
        └── main.go

```