## 简介

IM 即时聊天系统demo的后端，实现了单聊、群聊、表情发送、图片发送、语音发送。

## 使用的技术栈
* go-zero
* ent
* redis
* postgresql

## 配置文件
* etc\go-zero-chat.yaml
```yaml
Name: go-zero-chat
Host: 0.0.0.0
Port: 8081

# 数据库配置
DatabaseConf:
  Type: postgres
  Host: 127.0.0.1
  Port: 5432
  DBName: 数据库的名称
  Username: xxxx
  Password: xxxx
  MaxOpenConn: 100
  SSLMode: disable
  CacheTime: 5
  PGConfig: TimeZone=Asia/Shanghai
# Redis配置
RedisConf:
  Host: 127.0.0.1:6379
  Type: node
# JWT配置
Auth:
  AccessSecret: 300dc233-6227-4224-a5c6-180e0ss397e1
  AccessExpire: 604800
  RefreshTime: 86400
# 多点登录拦截
MultiSignOnBlocking:
  Type: true
# 日志配置
Log:
  ServiceName: go-zero-chat-logger
  Mode: file
  Level: info
  Compress: false
  KeepDays: 7
  StackCoolDownMillis: 100
# 文件上传路径
Local:
  Path: uploads

```

## 项目运行
* 下载依赖
```bash
go mod tidy
```

* 运行
```bash
go run go-zero-chat.go
```

* .api 文件生成api 服务
```bash
# go-zero文档https://go-zero.dev/
# goctl安装教程https://go-zero.dev/docs/tasks/installation/goctl
goctl api go -api ./api/all.api -dir ./ -style go_zero -home ./template/.goctl
```

* ent schema 生成
```bash
# ent文档https://entgo.io/
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery ./schema
```