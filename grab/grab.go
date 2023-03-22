package grab

import (
	"douyin-grab/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

type RoomInfo struct {
	App struct {
		InitialState struct {
			RoomStore struct {
				RoomInfo struct {
					RoomId string `json:"roomId"`
					Room   struct {
						Title        string `json:"title"`
						UserCountStr string `json:"user_count_str"`
					} `json:"room"`
				} `json:"roomInfo"`
			} `json:"roomStore"`
		} `json:"initialState"`
	} `json:"app"`
}

func FetchLiveRoomInfo(roomUrl string) (*RoomInfo, string) {
	req, err := http.NewRequest("GET", roomUrl, nil)
	if err != nil {
		logger.Error("fetch live room info err", err)
		return nil, ""
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36") // 设置User-Agent头
	cookie := &http.Cookie{Name: "__ac_nonce", Value: "063abcffa00ed8507d599"}
	req.AddCookie(cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("fetch live room info err", err)
		return nil, ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Error("read res body err", err)
		return nil, ""
	}

	pattern := regexp.MustCompile(`<script id="RENDER_DATA" type="application/json">(.*?)</script>`)
	data := pattern.FindSubmatch(body)
	decodedUrl, err := url.QueryUnescape(string(data[1]))
	if err != nil {
		logger.Error("url decode err", err)
		return nil, ""
	}

	var roomInfo RoomInfo
	err = json.Unmarshal([]byte(decodedUrl), &roomInfo)
	if err != nil {
		logger.Error("json unmarshal err", err)
		return nil, ""
	}
	logger.Info("roomid: %s", roomInfo.App.InitialState.RoomStore.RoomInfo.RoomId)
	logger.Info("title: %s", roomInfo.App.InitialState.RoomStore.RoomInfo.Room.Title)
	logger.Info("user_count: %s", roomInfo.App.InitialState.RoomStore.RoomInfo.Room.UserCountStr)

	var ttwid string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "ttwid" {
			ttwid = cookie.Value
		}
	}
	logger.Info("ttwid: %s", ttwid)

	return &roomInfo, ttwid
}
