package main

import (
	"douyin-grab/constv"
	"douyin-grab/grab"
	"douyin-grab/nmid"
	"douyin-grab/pkg/logger"
	"douyin-grab/wsocket"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

const (
	VERSION = `0.0.2`
)

func main() {
	app := cli.NewApp()
	app.Name = `Douyin Grab`
	app.Version = showVersion()
	app.Before = func(ctx *cli.Context) error {
		return nil
	}

	showBanner()
	godotenv.Load("./.env")
	logger.Init("")

	// var live_room_url, wss_url string
	var live_room_id string
	app.Flags = []cli.Flag{
		// cli.StringFlag{
		// 	Name:        "live_room_url, lrurl",
		// 	Usage:       "live room url",
		// 	Destination: &live_room_url,
		// },
		// cli.StringFlag{
		// 	Name:        "wss_url, wssurl",
		// 	Usage:       "live room wws url",
		// 	Destination: &wss_url,
		// },
		cli.StringFlag{
			Name:        "live_room_id, lrid",
			Usage:       "live room id",
			Destination: &live_room_id,
		},
	}

	var err error
	app.Action = func(ctx *cli.Context) error {
		// if len(live_room_url) == 0 {
		// 	live_room_url = constv.DEFAULTLIVEROOMURL //默认直播间url
		// }
		// logger.Info("live room url: %s", live_room_url)

		// if len(wss_url) == 0 {
		// 	wss_url = constv.DEFAULTLIVEWSSURL //默认直播间wss_url
		// }
		// logger.Info("live room wss_url: %s", wss_url)

		live_room_url := constv.DEFAULTLIVEROOMURL //默认直播间url
		wss_url := constv.DEFAULTLIVEWSSURL        //默认直播间wss_url

		if len(live_room_id) == 0 {
			live_room_id = constv.DEFAULTLIVEROOMID //默认直播间id
		}
		logger.Info("live room id: %s", live_room_id)

		if len(live_room_id) > 0 {
			live_room_url = fmt.Sprintf("%s/%s", constv.DOUYIORIGIN, live_room_id)
			logger.Info("live room url: %s", live_room_url)
			wssUrl, err := grab.GetWssUrl(live_room_url)
			logger.Info("get wss url %s", wssUrl)
			if nil == err {
				wss_url = wssUrl
			}
		}

		//获取直播间信息
		_, ttwid := grab.FetchLiveRoomInfo(live_room_url)

		//与直播间进行websocket通信，获取评论数据
		header := http.Header{}
		header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36") // 设置User-Agent头
		header.Set("Origin", constv.DOUYIORIGIN)
		cookie := &http.Cookie{
			Name:  "ttwid",
			Value: ttwid,
		}
		header.Add("Cookie", cookie.String())
		wsclient := wsocket.NewWSClient().SetRequestInfo(wss_url, header)
		wsclient.ConnWSServer(ttwid)
		wsclient.RunWSClient()

		//worker服务
		go nmid.RunWorker()

		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	os.Exit(0)
}

func showBanner() {
	println(`
	 _| _      . _    _  _ _  _|
	(_|(_)|_|\/|| |  (_|| (_|(_|
			 /        _|      `)
}

func showVersion() string {
	bannerData := `
	 _| _      . _    _  _ _  _|
	(_|(_)|_|\/|| |  (_|| (_|(_|
			 /        _|      `
	return bannerData + "\n" + VERSION
}
