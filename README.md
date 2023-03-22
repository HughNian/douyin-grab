## Golang版抖音相关数据抓取服务

## use
```go

go run main.go / ./bin/douyin-grab

```

目前需要传入抖音直播间url和直播间wssurl，可以写入常量constv中，也可以运行时传参
```go
./bin/douyin-grab -h

GLOBAL OPTIONS:
   --live_room_url value, --lrurl value  live room url
   --wss_url value, --wssurl value       live room wws url
   --help, -h                            show help
   --version, -v                         print the version


./bin/douyin-grab -lrurl xxxx -wssurl xxxx
```