# 项目说明
该项目为 go 备忘录项目，可以调用一些接口

- 技术栈：go

# 代码结构
```text
.
├── app
│   ├── gateway
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── internal
│   │   │   └── http
│   │   │       ├── task.go
│   │   │       └── user.go
│   │   ├── middleware
│   │   │   ├── cors.go
│   │   │   ├── init.go
│   │   │   └── jwt.go
│   │   ├── routes
│   │   │   └── router.go
│   │   ├── rpc
│   │   │   ├── init.go
│   │   │   ├── task.go
│   │   │   └── user.go
│   │   └── wrapper
│   │       ├── common.go
│   │       └── user.go
│   ├── task
│   │   ├── cmd
│   │   │   └── main.go
│   │   └── internal
│   │       ├── repository
│   │       │   └── db
│   │       │       ├── dao
│   │       │       │   ├── db_init.go
│   │       │       │   ├── migration.go
│   │       │       │   ├── task.go
│   │       │       │   └── task_test.go
│   │       │       └── model
│   │       │           └── task.go
│   │       └── service
│   │           └── task.go
│   └── user
│       ├── cmd
│       │   └── main.go
│       └── internal
│           ├── repository
│           │   └── db
│           │       ├── dao
│           │       │   ├── db_init.go
│           │       │   ├── migration.go
│           │       │   ├── user.go
│           │       │   └── user_test.go
│           │       └── model
│           │           └── user.go
│           └── service
│               └── user.go
├── config
│   ├── config.go
│   └── config.yml
├── consts
│   └── common.go
├── go.mod
├── go.sum
├── idl
│   └── pb
│       ├── task
│       │   ├── task_grpc.pb.go
│       │   └── task.pb.go
│       ├── task.proto
│       ├── user
│       │   ├── user_grpc.pb.go
│       │   └── user.pb.go
│       └── user.proto
├── Makefile
└── pkg
    ├── ctl
    │   ├── ctl.go
    │   └── user_info.go
    ├── discovery
    │   ├── instance.go
    │   ├── register.go
    │   └── resolver.go
    ├── e
    │   ├── code.go
    │   └── msg.go
    ├── res
    │   └── reponse.go
    └── util
        ├── jwt
        │   └── jwt.go
        ├── logger
        │   └── logger.go
        └── shutdown
            └── gracefully_shutdown.go
```

# 实现功能
1. 使用 jwt、gin、gorm、grpc、etcd 技术栈
2. 实现 etcd 服务注册与发现
3. 实现 gin 层面中间件拦截

# 运行须知
1. 先把 *dockerpack* 目录下的两个 yaml 文件执行，部署到 docker
2. 参考 *Makefile* 可以编译出文件

# 备忘命令
1. docker 导出 yaml 动作
```shell
#获取docker autocompose的镜像
docker pull ghcr.io/red5d/docker-autocompose:latest
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock ghcr.io/red5d/docker-autocompose 19add0a044fa > docker-compose-etcd.yaml
```
2. 增加 mock 动作