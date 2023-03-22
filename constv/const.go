package constv

import "time"

const DOUYIORIGIN = "https://live.douyin.com"
const DOUYINHOST = "webcast3-ws-web-hl.douyin.com"
const DOUYINPATH = "/webcast/im/push/v2"
const DEFAULTLIVEROOMURL = "https://live.douyin.com/80017709309"
const DEFAULTLIVEWSSURL = "wss://webcast3-ws-web-lf.douyin.com/webcast/im/push/v2/?app_name=douyin_web&version_code=180800&webcast_sdk_version=1.3.0&update_version_code=1.3.0&compress=gzip&internal_ext=internal_src:dim|wss_push_room_id:7213136262230723365|wss_push_did:7208120443965064762|dim_log_id:2023032209282439DEF28ACA41F59074E6|fetch_time:1679448504471|seq:1|wss_info:0-1679448504471-0-0|wrds_kvs:InputPanelComponentSyncData-1679439296929724579_WebcastRoomStatsMessage-1679448500943119124_WebcastRoomRankMessage-1679448470989228757&cursor=r-1_d-1_u-1_h-1_t-1679448504471&host=https://live.douyin.com&aid=6383&live_id=1&did_rule=3&debug=false&maxCacheMessageNumber=20&endpoint=live_pc&support_wrds=1&im_path=/webcast/im/fetch/&user_unique_id=7208120443965064762&device_platform=web&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=zh-CN&browser_platform=Win32&browser_name=Mozilla&browser_version=5.0%20(Windows%20NT%2010.0;%20Win64;%20x64)%20AppleWebKit/537.36%20(KHTML,%20like%20Gecko)%20Chrome/111.0.0.0%20Safari/537.36&browser_online=true&tz_name=Asia/Shanghai&identity=audience&room_id=7213136262230723365&heartbeatDuration=0&signature=WsRYL0+P68+eiCrC"
const DEFAULTHEARTBEATTIME = time.Second * 10
