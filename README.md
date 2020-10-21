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

### 问题解决
>module github.com/micro/go-plugins/v2@latest found (v2.0.0), but does not contain package github.com/micro/go-plugins/v2/registry/consul
到GitHub下载go-plugins的v2最后一版，替换v2文件夹下的内容即可


### 准备工作
- 完成mysql，consul等配置
- 将配置详情填写到config/config.toml文件对应项中

### 启动方式
- dockerfile
#### 或
- 先执行env.go,修改包依赖
- 然后执行server/start.go文件

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