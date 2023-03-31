from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
import time

# 启动Chrome浏览器，并打开WSS调试端口
chrome_options = webdriver.ChromeOptions()
chrome_options.add_argument("headless")             # 无界面
chrome_options.add_argument("--no-sandbox")           # root用户启动需要开启
chrome_options.add_argument("--disable-dev-shm-usage")
chrome_options.add_argument("--remote-debugging-port=9222")
chrome_options.add_argument("--disable-background-timer-throttling")
chrome_options.add_argument("--disable-backgrounding-occluded-windows")
chrome_options.add_argument("--disable-renderer-backgrounding")
chrome_options.add_argument("user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3") #需要agnet,不然一些网站会有限制
#chrome_options.binary_location = "/usr/bin/chromium-browser"  #linux下需要

# service = Service("/path/to/chromedriver") #linux下需要
service = Service("D:\chromedriver.exe") #windows下需要
caps = DesiredCapabilities.CHROME
caps["goog:loggingPrefs"] = {"performance": "ALL"}
driver = webdriver.Chrome(service=service, options=chrome_options, desired_capabilities=caps) #selenium4方式

# 等待页面加载完成
driver.get("https://live.douyin.com/80017709309")
time.sleep(2)

# 获取所有的网络请求日志信息
logs = driver.get_log("performance")

# 打印所有的WSS请求信息
for log in logs:
    message = log["message"]
    if "webSocket" in message:
        print(message)

driver.quit()