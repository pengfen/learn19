from selenium import webdriver
from selenium.webdriver import ActionChains #引入动作链
import time

# 交互动作 驱动浏览器进行动作 模拟拖动作 将动作附加到动作链中串行执行
browser = webdriver.Chrome()
url = 'http://www.runoob.com/try/try.php?filename=jqueryui-api-droppable'
browser.get(url)
browser.switch_to.frame('iframeResult') #切换到iframeResult框架
source = browser.find_element_by_css_selector('#draggable') # 找到被拖对象
target = browser.find_element_by_css_selector('#droppable') # 找到目标
actions = ActionChains(browser) # 声明actions对象
actions.drag_and_drop(source, target)
actions.perform() #执行动作
time.sleep(5)
