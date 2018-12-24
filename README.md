### 编译main
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

### 执行定时任务
```
crontab every-15m.cron
```
### 定时任务相关命令
```
service cron reload
service cron restart
```

