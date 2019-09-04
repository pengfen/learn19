from selenium import webdriver
import time

# 元素交互操作 - 搜索框传入关键词进行自动搜索
browser = webdriver.Chrome()
browser.get("https://www.baidu.com")
input = browser.find_element_by_id('kw') #搜索框
input.send_keys('iPhone')
time.sleep(3)
input.clear() #清空搜索框
input.send_keys('welcome')
button = browser.find_element_by_id('su')
button.click()
time.sleep(3)
browser.close()
