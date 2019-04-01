<a href="https://996.icu"><img src="https://img.shields.io/badge/link-996.icu-red.svg" alt="996.icu"></a>

[![LICENSE](https://img.shields.io/badge/license-NPL%20(The%20996%20Prohibited%20License)-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)


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

#### A Guy Who Dresses Up As A Bat Clearly Has Issues
