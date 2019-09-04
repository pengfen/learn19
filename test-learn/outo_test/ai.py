from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import random
import time

# 启动浏览器
driver = webdriver.Chrome(r"chromedriver.exe")
#driver = webdriver.Chrome()
driver.get("http://ai.baidu.com/tech/speech/tts")

# 获取当前时间
current_time = time.strftime('%H:%M',time.localtime(time.time()))

# 输入当前时间
input_text = driver.find_element_by_css_selector('.nlp-textarea')
input_text.clear()
input_text.send_keys(current_time)

# 点击播放按钮
playBtn = driver.find_element_by_css_selector('.demo-speech-tts-play')
playBtn.click()
time.sleep(3)

driver.close()