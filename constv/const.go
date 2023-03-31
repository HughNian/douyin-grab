package constv

import "time"

const DOUYIORIGIN = "https://live.douyin.com"
const DOUYINHOST = "webcast3-ws-web-hl.douyin.com"
const DOUYINPATH = "/webcast/im/push/v2"
const DEFAULTLIVEROOMURL = "https://live.douyin.com/80017709309"
const DEFAULTLIVEWSSURL = "wss://webcast3-ws-web-lq.douyin.com/webcast/im/push/v2/?app_name=douyin_web&version_code=180800&webcast_sdk_version=1.3.0&update_version_code=1.3.0&compress=gzip&internal_ext=internal_src:dim|wss_push_room_id:7213879049037351741|wss_push_did:7212546907507426876|dim_log_id:2023032408370306ED00C9212A7373D8F3|fetch_time:1679618223364|seq:1|wss_info:0-1679618223364-0-0|wrds_kvs:WebcastRoomStatsMessage-1679618222864617493_InputPanelComponentSyncData-1679612204850357394_WebcastRoomRankMessage-1679618186891224988&cursor=u-1_h-1_t-1679618223364_r-1_d-1&host=https://live.douyin.com&aid=6383&live_id=1&did_rule=3&debug=false&maxCacheMessageNumber=20&endpoint=live_pc&support_wrds=1&im_path=/webcast/im/fetch/&user_unique_id=7212546907507426876&device_platform=web&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=zh-CN&browser_platform=Win32&browser_name=Mozilla&browser_version=5.0%20(Windows%20NT%2010.0;%20Win64;%20x64)%20AppleWebKit/537.36%20(KHTML,%20like%20Gecko)%20Chrome/111.0.0.0%20Safari/537.36%20Edg/111.0.1661.41&browser_online=true&tz_name=Asia/Shanghai&identity=audience&room_id=7213879049037351741&heartbeatDuration=0&signature=RMvK5n+6uFRU6ou8"
const DEFAULTHEARTBEATTIME = time.Second * 10
const DEFAULTLIVEROOMID = "80017709309"
