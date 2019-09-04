# 获取文本值

from selenium import webdriver
import time

browser = webdriver.Chrome()
url = 'https://www.zhihu.com/explore'
browser.get(url)
input = browser.find_element_by_class_name('zu-top-add-question')
print(input.text) # input.text 文本值
time.sleep(2)
input = browser.find_element_by_class_name('zu-top-add-question')
print(input.id) # 获取id
print(input.location) # 获取位置
print(input.tag_name) # 获取标签名
print(input.size) # 获取大小
browser.close()
