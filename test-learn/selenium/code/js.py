from selenium import webdriver
import time

# 有些动作可能没有提供api 比如进度条 可以通过代码执行JavaScript
browser = webdriver.Chrome()
browser.get('https://www.zhihu.com/explore')
browser.execute_script('window.scrollTo(0, document.body.scrollHeight)')
browser.execute_script('alert("To Bottom")')
time.sleep(5)
browser.close()
