# 采用微服务架构搭建博客
## 目前用到
- grpc
- gin
- gorm
- micro


## 项目功能
项目名：Form
- 多人博客展示
- 用户认证
- 用户只能修改自己的博客
- 用户/博客的增删改查

## 目录结构
```
.
├── README.md
├── client
│   ├── client.go
│   ├── microClient.go
│   └── rpc
│       ├── auth.go
│       └── blog.go
├── config
│   ├── config.go
│   └── config.toml
├── go.mod
├── go.sum
├── proto
│   ├── blog.proto
│   ├── grpc
│   │   └── blog.pb.go
│   └── micro
│       ├── blog.pb.go
│       └── blog.pb.micro.go
└── server
    ├── db
    │   └── db.go
    ├── gRPCServer.go
    ├── microServer.go
    ├── microimp
    │   ├── auth.go
    │   └── blog.go
    ├── model
    │   ├── blog.go
    │   └── user.go
    ├── rpcimpl
    │   ├── auth.go
    │   └── blog.go
    └── web
        └── main.go

12 directories, 23 files

```