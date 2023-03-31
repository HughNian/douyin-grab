## golang版抖音直播间弹幕📃礼物🎁等数据抓取

## 🤩🤩
直接传入抖音直播间id就可抓取抖音直播间信息   

如https://live.douyin.com/80017709309 这个直播间网址，最后的数字就为直播间id，不用手动去页面复制wssurl了。 

## ✍✍
需要注意的是，需要本地有chromedriver，在本项目中有一个win版，一个linux版的chromedriver。但最好需要与你本地chrome/chromium-browser版本一直，否则会有问题。另外如果是python运行的，需要在启动Xvfb   

```shell
Xvfb :99 -ac &
export DISPLAY=:99

```

最后要感谢selenium

## use
```go

./bin/douyin-grab.exe -lrid 168465302284

```

```go
./bin/douyin-grab -h

GLOBAL OPTIONS:
   --live_room_id value, --lrid value  live room id
   --help, -h                            show help
   --version, -v                         print the version


./bin/douyin-grab -lrurl xxxx -wssurl xxxx
```  

![](https://raw.githubusercontent.com/HughNian/douyin-grab/main/images/2.png)  

![](https://raw.githubusercontent.com/HughNian/douyin-grab/main/images/3.png)  

![](https://raw.githubusercontent.com/HughNian/douyin-grab/main/images/1.png)  
