from selenium import webdriver
import time

# 注意 python文件名或者包名不要命名为selenium 会导致无法导入

# 访问页面并获取网页html
browser = webdriver.Chrome()
browser.get("https://www.baidu.com")
# print(browser.page_source) # browser.page_source 是获取网页的全部html
browser.close()
time.sleep(3)
browser.close()
