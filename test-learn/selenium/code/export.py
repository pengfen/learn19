# from selenium import webdriver

# 异常处理
# browser = webdriver.Chrome()
# browser.get('https://www.baidu.com')
# browser.find_element_by_id('hello')

from selenium import webdriver
from selenium.common.exceptions import TimeoutException, NoSuchElementException

browser = webdriver.Chrome()
try:
    browser.get('https://www.baidu.com')
except TimeoutException:
    print('Time Out')
try:
    browser.find_element_by_id('hello')
except NoSuchElementException:
    print('No Element')
finally:
    browser.close()
