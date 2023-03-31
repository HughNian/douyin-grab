package grab

import (
	"douyin-grab/pkg/logger"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	slog "github.com/tebeka/selenium/log"
)

type DYMessage struct {
	Message struct {
		Method string `json:"method"`
		Params struct {
			Url string `json:"url"`
		} `json:"params"`
	} `json:"message"`
}

func GetWssUrl(liveRoomUrl string) (string, error) {
	var (
		chromeDriverPath = os.Getenv("CHROME_DRIVER_PATH")
		chromePath       = os.Getenv("CHROME_PATH")
		port, _          = strconv.Atoi(os.Getenv("CHROME_DRIVER_SERVICE_PORT"))
	)

	options := []selenium.ServiceOption{}
	//selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, options...)
	if err != nil {
		logger.Error("selenium new chrome driver service error %s", err.Error())
		return "", err
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
		"goog:loggingPrefs": map[string]string{ //can get log with this, important!!
			"performance": "ALL",
		},
	}
	chromeCaps := chrome.Capabilities{
		Path: chromePath,
		Args: []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-background-timer-throttling",
			"--disable-backgrounding-occluded-windows",
			"--disable-renderer-backgrounding",
			"user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		},
	}
	caps.AddChrome(chromeCaps)
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		logger.Error("selenium new remote error %s", err.Error())
		return "", err
	}
	defer webDriver.Quit()

	//start request
	_ = webDriver.Get(liveRoomUrl)

	//wait web page loading
	time.Sleep(2 * time.Second)

	//get request log
	logs, err := webDriver.Log(slog.Performance)
	if nil == err {
		if nil == err {
			for _, message := range logs {
				// fmt.Println(message.Message)
				var msg DYMessage
				err = json.Unmarshal([]byte(message.Message), &msg)
				if nil == err {
					if msg.Message.Method == "Network.webSocketCreated" {
						wssurl := msg.Message.Params.Url

						return wssurl, nil
					}
				} else {
					return "", err
				}

				// file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// defer file.Close()
				// _, err = file.WriteString(message.Message)
				// if err != nil {
				// 	log.Fatal(err)
				// }
			}
		}
	}

	// 截屏
	//img, _ := webDriver.Screenshot()
	//_ = ioutil.WriteFile("douyin.png", img, 0666)

	return "", nil
}
