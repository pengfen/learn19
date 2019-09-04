# 获取元素信息 - 获取属性
from selenium import webdriver
from selenium.webdriver import ActionChains

browser = webdriver.Chrome()
url = 'https://www.zhihu.com/explore'
browser.get(url)
logo = browser.find_element_by_id('zh-top-link-logo') #获取网站logo
print(logo)
print(logo.get_attribute('class'))
browser.close()
