# 简单设备认证服务

马良笔的认证代理服务。

## Build

```sh
GOOS=linux GOARCH=amd64 go build
```

## Deploy

```sh
scp .\things-auth-service root@101.34.24.70:/root/iothub
```
