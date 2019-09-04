from selenium import webdriver
from selenium.common.exceptions import NoSuchElementException
import time

# frame相当于独立的网页 如果在父类网frame查找子类的 则必须切换到子类的frame 子类如果查找父类也需要先切换
browser = webdriver.Chrome()
url = 'http://www.runoob.com/try/try.php?filename=jqueryui-api-droppable'
browser.get(url)
browser.switch_to.frame('iframeResult')
source = browser.find_element_by_css_selector('#draggable')
print(source)
try:
    logo = browser.find_element_by_class_name('logo')
except NoSuchElementException:
    print('NO LOGO')
browser.switch_to.parent_frame()
logo = browser.find_element_by_class_name('logo')
print(logo)
print(logo.text)
time.sleep(5)
browser.close()
