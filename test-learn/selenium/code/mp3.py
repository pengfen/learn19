from selenium import webdriver

browser = webdriver.Chrome()
browser.get("http://ai.baidu.com/tech/speech/tts");

input_close = browser.find_element_by_css_selector(".nlp-textarea-clear")
#input_close.click()

input = browser.find_element_by_css_selector('.nlp-textarea')
input.clear()
input.send_keys("当前时间")

# 播放
player = browser.find_element_by_css_selector(".demo-speech-tts-control-text")
player.click()
